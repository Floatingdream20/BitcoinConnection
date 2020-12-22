package bitcoinService

import (
	"Bitcoin_core_web/Bitcoin_conne"
	"fmt"
	"net/http"
)

func Order() {
	switch {

	}

}

func GetBlockCount() (interface{}, error) {
	json,err := Bitcoin_conne.PrepareJSON("getblockcount",nil)
	if err != nil {
		fmt.Println("service:21行")
		return nil, err
	}
	fmt.Println(json)
	 RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		fmt.Println("service:26行")
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil,err
}
