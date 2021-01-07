package Bitcoin_conne

import (
	"Bitcoin_core_web/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func PrepareJSON(method string, prejson models.Prejson) (string, error) {

	request := models.RPCrequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: RPCVERSION,
	}

	switch method {
	case "getblockcount":
		request.Params = nil
		break
	case "getblock":
		request.Params = []interface{}{prejson.Args1}
		break
	case "getbestblockhash":
		request.Params = nil
		break
	case "getblockhash":
		 height,_:=strconv.Atoi(prejson.Args1)

		request.Params = []interface{}{height}
		break
	case "getnewaddress":
		request.Params = nil
		break
	case "getblockchaininfo":
		request.Params = nil
		break
	case "getdifficulty":
		request.Params = nil
		break
	case "uptime":
		request.Params = nil
		break
	case "validateaddress":
		request.Params = []interface{}{prejson.Args1}
		break
	case "getbalances":
		request.Params = nil
		break
	case "getmemoryinfo":
		request.Params = nil
		break
	case "gettxoutsetinfo":
		request.Params = nil
		break
	case "getmempoolinfo":
		request.Params = nil
		break
	case "getchaintxstats":
		request.Params = []interface{}{prejson.Args1,prejson.Args2}
		break
	case "getblockheader":
		request.Params = []interface{}{prejson.Args1}
		break
	case "getwalletinfo":
		request.Params = nil
		break
	case "getnetworkinfo":
		request.Params = nil
		break
	case "getmininginfo":
		request.Params = nil
		break

	}

	requestbyte, err := json.Marshal(&request)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(requestbyte), nil
}

//发起请求。。然后需要
func Dopost(url string, header map[string]string, body string) (*models.Rpcresult, error) {
	client := http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {

		return nil, err
	}

	for k, v := range header {
		request.Header.Add(k, v)
	}
	reponse, err := client.Do(request)
	if err != nil {

		return nil, err
	}
	code := reponse.StatusCode
  fmt.Println(code)

	defer reponse.Body.Close()

	reponsebyte, err := ioutil.ReadAll(reponse.Body)
	if err != nil {

		return nil, err
	}

	var result models.Result
	err = json.Unmarshal(reponsebyte, &result)
	if err != nil {
		return nil, err
	}

	var rpcresult models.Rpcresult
	if code == http.StatusOK {
		rpcresult.Code = code
		rpcresult.Msg = "请求成功"
		rpcresult.Data = result
	} else {
		rpcresult.Code = code
		rpcresult.Msg = "请求失败"
		rpcresult.Data = result
	}
	return &rpcresult, nil
}
func Base64(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

/*
**请求头设置
 */
func CreateMap() map[string]string {
	Map := make(map[string]string)
	Map["Encoding"] = "UTF-8"
	Map["Content-Type"] = "application/json"
	Map["Authorization"] = "Basic " + Base64(RPCUSER+":"+RPCPASSWORD)
	return Map
}
