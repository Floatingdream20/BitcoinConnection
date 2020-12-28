package bitcoinService

import (
	"Bitcoin_core_web/Bitcoin_conne"
	"Bitcoin_core_web/Entity"
	"fmt"
	"github.com/mapstructure-master"
	"net/http"
)

func Order(method string, json string) interface{} {
	Map := make(map[string]interface{})
	Map["getblockcount"] = GetBlockCount(json)//无参数
	Map["getbestblockhash"] = GetBestBlockHash(json)//无参数
	Map["getblockhash"] = GetBlockHash(json)// 高度 int
	Map["getblock"] = GetBlock(json) //区块hash
	Map["getnewaddress"] = GetNewAddress(json)//无参数
	Map["getblockchaininfo"] = GetBlockChainInfo(json)//无参数
	Map["getdifficulty"] = GetDifficulty(json)//无参数
	Map["uptime"] = UpTime(json)//无参数
	Map["validateaddress"] = ValidateAddress(json)//btc地址
	Map["getbalances"] = GetBalances(json)//无参数

	for k, _ := range Map {
		if method == k {
			return Map[k]
		}
	}
	return nil
}

//查看当前节点同步的区块个数
func GetBlockCount(json string) int {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit.(int)
	}
	fmt.Println(err)
	return -1
}

//获取最新区块的hash
func GetBestBlockHash(json string) string {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit.(string)
	}
	fmt.Println(err)
	return ""
}

//通过区块高度查看区块hash
func GetBlockHash(json string) string {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit.(string)
	}
	fmt.Println(err)
	return ""
}

//根据区块hash获取区块信息
func GetBlock(json string) *Entity.Getblock {
	var block Entity.Getblock

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &block)
		if err!=nil {
			fmt.Println(err)
			return nil
		}
		return &block
	}
	fmt.Println(err)
	return nil
}

//生成一个新的比特币地址
func GetNewAddress(json string) string {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit.(string)
	}
	fmt.Println(err)
	return ""
}

//获取区块链的数据
func GetBlockChainInfo(json string) *Entity.ChainInfo {
	var block Entity.ChainInfo

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &block)
		if err!=nil {
			fmt.Println(err)
			return nil
		}
		return &block
	}
	fmt.Println("", err)
	return nil
}

//以最小难度的倍数返回工作证明难度
func GetDifficulty(json string) int {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit.(int)
	}
	fmt.Println(err)
	return -1
}

//返回服务器的总运行时间
func UpTime(json string) int {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit.(int)
	}
	fmt.Println(err)
	return -1
}

//返回关于给定比特币地址的信息
func ValidateAddress(json string) *Entity.Validata {
	var block Entity.Validata
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		err=mapstructure.WeakDecode(RpcResult.Data.Resultbit,&block)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return &block
	}
	fmt.Println("", err)
	return nil
}

//返回一个在BTC中包含所有余额的对象
func GetBalances(json string) *Entity.Balances {
	var balance Entity.Balances
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit,&balance)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}
//getmemoryinfo
func GetMemoryInfo()  {
	json,err:=Bitcoin_conne.PrepareJSON("getmemoryinfo",nil)
	if err != nil {
		fmt.Println(err)
	}
	result,err:=Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL,Bitcoin_conne.CreateMap(),json)
	if err != nil {
		fmt.Println(err)
	}
	if result.Code==http.StatusOK {
		fmt.Println("getmemoryinfo",result.Data.Resultbit)
	}
}
//gettxoutsetinfo
//getmempoolinfo
//getchaintxstats
//getchaintips
//getblockheader