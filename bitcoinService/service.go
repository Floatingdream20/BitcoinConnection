package bitcoinService

import (
	"Bitcoin_core_web/Bitcoin_conne"
	"Bitcoin_core_web/Entity"
	"fmt"
	"github.com/mitchellh/mapstructure"

	"net/http"
)

func Order(method string, json string) interface{} {
	Map := make(map[string]interface{})
	Map["getblockcount"] = GetBlockCount(json)
	Map["getbestblockhash"] = GetBestBlockHash(json)
	Map["getblockhash"] = GetBlockHash(json)
	Map["getblock"] = GetBlock(json)
	Map["getnewaddress"] = GetNewAddress(json)
	Map["getblockchaininfo"] = GetBlockChainInfo(json)
	Map["getdifficulty"] = GetDifficulty(json)
	Map["uptime"] = UpTime(json)
	Map["validateaddress"] = ValidateAddress(json)
	Map["getbalances"] = GetBalances(json)
	Map["help"] = GetHelp(json)
	Map["stop"] = GetStop(json)
	Map["getblockstatus"] = GetBlockStatus(json)
	Map["getwalletinfo"] = GetWalletInfo(json)
	Map["getrawchangeaddress"] = GetRawChangeAddress(json)
	Map["getnetworkinfo"] = GetNetWorkInfo(json)
	Map["gettxoutsetinfo"] = GetTxOutSetInfo(json)
	Map["getmininginfo"] = GetMiningInfo(json)
	for k, _ := range Map {
		if method == k {
			return Map[k]
		}

	}
	return nil
}

//1.查看当前节点同步的区块个数
func GetBlockCount(json string) interface{} {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//2.获取最新区块的hash
func GetBestBlockHash(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//3.通过区块高度查看区块hash
func GetBlockHash(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//4.根据区块hash获取区块信息
func GetBlock(json string) *Entity.Getblock {
	var block Entity.Getblock

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	//var block Entity.Getblock
	if RpcResult.Code == http.StatusOK {
		result := RpcResult.Data.Resultbit

		err = mapstructure.WeakDecode(result, &block)
		return &block
	}
	fmt.Println("", err)
	return nil
}

//5.生成一个新的比特币地址
func GetNewAddress(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//6.获取区块链的数据
func GetBlockChainInfo(json string) *Entity.ChainInfo {
	var block Entity.ChainInfo

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		return &block
	}
	fmt.Println("", err)
	return nil
}

//7.以最小难度的倍数返回工作证明难度
func GetDifficulty(json string) interface{} {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//8.返回服务器的总运行时间
func UpTime(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//9.返回关于给定比特币地址的信息
func ValidateAddress(json string) *Entity.Validata {
	var block Entity.Validata

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		return &block
	}
	fmt.Println("", err)
	return nil
}

//10.返回一个在BTC中包含所有余额的对象
func GetBalances(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//11.获取帮助
func GetHelp(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//12.退出bitcoin
func GetStop(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//13.获取区块状态
func GetBlockStatus(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//14.获取钱包
func GetWalletInfo(json string) *Entity.Getwalletinfo {
	var getwalletinfo Entity.Getwalletinfo

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		return &getwalletinfo
	}
	fmt.Println("", err)
	return nil
}

//15.获取原始的地址
func GetRawChangeAddress(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		return err
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	return err
}

//16.获取网络节点的信息
func GetNetWorkInfo(json string) *Entity.NetWorkInfo {
	var getnetworkinfo Entity.NetWorkInfo

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		return &getnetworkinfo
	}
	fmt.Println("", err)
	return nil
}

//17.获取交易输出集合信息
func GetTxOutSetInfo(json string) *Entity.GetTxOutSetInfo {
	var gettxoutsetinfo Entity.GetTxOutSetInfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		return &gettxoutsetinfo
	}
	fmt.Println("", err)
	return nil
}

//18.获取挖矿信息
func GetMiningInfo(json string) *Entity.GetMiningInfo {
	var getmininginfo Entity.GetMiningInfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		return &getmininginfo
	}
	fmt.Println("", err)
	return nil
}

//19.获取网络算力
//
//func GetNetWorkHashPs(json string, height int64) interface{} {
//	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
//	if err != nil {
//		return err
//	}
//	if RpcResult.Code == http.StatusOK {
//		return RpcResult.Data.Resultbit
//	}
//	return err
//}
