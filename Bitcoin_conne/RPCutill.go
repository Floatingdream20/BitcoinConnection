package Bitcoin_conne

import (
	"BitcoinConnection/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func PrepareJSON(method string, param interface{}) (string, error) {
	request := models.RPCrequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: RPCVERSION,
		Param:   param,
	}
	requestbyte, err := json.Marshal(&request)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(requestbyte), nil
}

//发起请求。。然后需要
func Dopost(url string, header map[string]string, body io.Reader) (*models.Rpcresult, error) {
	client := http.Client{}
	request, err := http.NewRequest("POST", url, body)
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
func CreateMap() map[string]string {
	Map := make(map[string]string)
	Map["Encoding"] = "UTF-8"
	Map["Content-Type"] = "application/json"
	Map["Authorization"] = "Basic " + Base64(RPCUSER+":"+RPCPASSWORD)
	return Map
}
