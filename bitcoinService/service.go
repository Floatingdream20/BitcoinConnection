package bitcoinService

import (
	"Bitcoin_core_web/Bitcoin_conne"
	"Bitcoin_core_web/Entity"
	"github.com/mapstructure-master"
	"net/http"
)

func Order(method string, param interface{}) (interface{}, error) {
	var data interface{}
	var err error
	switch method {
	case "getblockcount":
		data, err = GetBlockCount(method, param)
		return data, err
	case "getbestblockhash":
		data, err = GetBestBlockHash(method, param)
		return data, err
	case "getblockhash":
		data, err = GetBlockHash(method, param)
		return data, err
	case "getblock":
		data, err = GetBlock(method, param)
		return data, err
	case "getnewaddress":
		data, err = GetNewAddress(method, param)
		return data, err
	case "getblockchaininfo":
		data, err = GetBlockChainInfo(method, param)
		return data, err
	case "getdifficulty":
		data, err = GetDifficulty(method, param)
		return data, err
	case "uptime":
		data, err = UpTime(method, param)
		return data, err
	case "validateaddress":
		data, err = ValidateAddress(method, param)
		return data, err
	case "getbalances":
		data, err = GetBalances(method, param)
		return data, err
	}

	return nil, nil
}
//查看当前节点同步的区块个数
func GetBlockCount(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}
//获取最新区块的hash
func GetBestBlockHash(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}

//通过区块高度查看区块hash
func GetBlockHash(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}

//根据区块hash获取区块信息
func GetBlock(method string, param interface{}) (Entity.Getblock, error) {
	var block Entity.Getblock
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return block, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return block, err
	}
	//var block Entity.Getblock
	if RpcResult.Code == http.StatusOK {
		result:= RpcResult.Data.Resultbit

		err=mapstructure.WeakDecode(result,&block)
		return block,err
	}
	return block, err
}

//生成一个新的比特币地址
func GetNewAddress(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}

//获取区块链的数据
func GetBlockChainInfo(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}

//以最小难度的倍数返回工作证明难度
func GetDifficulty(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}

//返回服务器的总运行时间
func UpTime(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}

//返回关于给定比特币地址的信息
func ValidateAddress(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}

//返回一个在BTC中包含所有余额的对象
func GetBalances(method string, param interface{}) (interface{}, error) {
	json, err := Bitcoin_conne.PrepareJSON(method, param)
	if err != nil {
		return nil, err
	}
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return nil, err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit, nil
	}
	return nil, err
}
