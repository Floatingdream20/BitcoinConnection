package bitcoinService

import (
	"Bitcoin_core_web/Bitcoin_conne"
	"Bitcoin_core_web/Entity"
	"Bitcoin_core_web/models"
	"fmt"
	"github.com/mapstructure-master"
	"net/http"
)

func Mapbymethod(method string, json string) map[string]interface{} {
	Map := make(map[string]interface{})
	switch method {
	case "getblockcount":
		Map["Alltype"] = GetBlockCount(json) //无参数
		break
	case "getblock":
		Map["Alltype"] = GetBlock(json) //区块hash
		break
	case "getbestblockhash":
		Map["Alltype"] = GetBestBlockHash(json) //无参数
		break
	case "getblockhash":
		Map["Alltype"] = GetBlockHash(json) // 高度 int
		break
	case "getnewaddress":
		Map["Alltype"] = GetNewAddress(json) //无参数
		break
	case "getblockchaininfo":
		Map["Alltype"] = GetBlockChainInfo(json) //无参数
		break
	case "getdifficulty":
		Map["Alltype"] = GetDifficulty(json) //无参数
		break
	case "uptime":
		Map["Alltype"] = UpTime(json) //无参数
		break
	case "validateaddress":
		Map["Alltype"] = ValidateAddress(json) //btc地址
		break
	case "getbalances":
		Map["Alltype"] = GetBalances(json) //无参数
		break
	case "getmemoryinfo":
		Map["Alltype"] = GetMemoryInfo(json) //无参数
		break
	case "gettxoutsetinfo":
		Map["Alltype"] = Txoutset(json) //无参数
		break
	case "getmempoolinfo":
		Map["Alltype"] = MemPoolInfo(json) //无参数
		break
	case "getchaintxstats":
		Map["Alltype"] = ChainTxStats(json) //（可选参数 nblocks int 第几个区块 blockhash 块的散列hash ）
		break
	case "getblockheader":
		Map["Alltype"] = BlockHeader(json) //(blockhash 如果verbose是false,则返回一个序列化的、hex编码数据的字符串,用于块头“散列”。
		break
	case "getwalletinfo":
		Map["Alltype"] = Getwalletinfo(json) //无参数
		break
	case "getnetworkinfo":
		Map["Alltype"] = GetNetWorkInfo(json) //无参数
		break
	case "getmininginfo":
		Map["Alltype"] = GetMiningInfo(json) //无参数
		break
	}
	return Map
}

func Order(method string, json string) *models.AllType {

	Map := Mapbymethod(method, json)
	var alltype models.AllType
	err := mapstructure.WeakDecode(Map, &alltype)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(alltype)
	return &alltype
}

//查看当前节点同步的区块个数
func GetBlockCount(json string) interface{} {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	fmt.Println(err)
	return -1
}

//获取最新区块的hash
func GetBestBlockHash(json string) interface{} {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	fmt.Println(err)
	return ""
}

//通过区块高度查看区块hash
func GetBlockHash(json string) interface{} {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	fmt.Println(err)
	return ""
}

//根据区块hash获取区块信息
func GetBlock(json string) Entity.Getblock {
	var block Entity.Getblock

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)

	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &block)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(err)
	return block
}

//生成一个新的比特币地址
func GetNewAddress(json string) interface{} {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	fmt.Println(err)
	return ""
}

//获取区块链的数据
func GetBlockChainInfo(json string) Entity.ChainInfo {
	var block Entity.ChainInfo

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)

	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &block)
	if err != nil {
		fmt.Println(err)

	}

	return block
}

//以最小难度的倍数返回工作证明难度
func GetDifficulty(json string) interface{} {
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	fmt.Println(err)
	return -1
}

//返回服务器的总运行时间
func UpTime(json string) interface{} {

	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if RpcResult.Code == http.StatusOK {
		return RpcResult.Data.Resultbit
	}
	fmt.Println(err)
	return -1
}

//获取关于给定比特币地址的信息
func ValidateAddress(json string) Entity.Validata {
	var block Entity.Validata
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)

	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &block)
		if err != nil {
			fmt.Println(err)

		}

	}
	return block
}

//获取一个在BTC中包含所有余额的对象
func GetBalances(json string) Entity.Balances {
	var balance Entity.Balances
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)

	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &balance)
	if err != nil {
		fmt.Println(err)
	}

	return balance
}

//getmemoryinfo  获取一个包含内存使用信息的对象 无参数
func GetMemoryInfo(json string) Entity.Locked {
	var locked Entity.Locked
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &locked)
		if err != nil {
			fmt.Println(err)

		}
	}

	return locked
}

//gettxoutsetinfo  获取关于未使用事务输出集的统计信息 无参数
func Txoutset(json string) Entity.Txoutsetinfo {
	var txout Entity.Txoutsetinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &txout)
	if err != nil {
		fmt.Println(err)

	}

	return txout
}

//getmempoolinfo 获取内存池的详细信息 无参数
func MemPoolInfo(json string) Entity.Mempoollinfo {
	var mempool Entity.Mempoollinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &mempool)
	if err != nil {
		fmt.Println(err)

	}

	return mempool
}

//getchaintxstats 计算链中交易总数和费率的统计数字 （可选参数 nblocks int 第几个区块 blockhash 块的散列hash ）
func ChainTxStats(json string) Entity.Chaintxstats {
	var chain Entity.Chaintxstats
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &chain)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(err)
	return chain
}

//getblockheader (blockhash 如果verbose是false,则返回一个序列化的、hex编码数据的字符串,用于块头“散列”。
//如果verbose是true,则返回一个带有块头消息的对象)
func BlockHeader(json string) Entity.BlockHeader {
	var header *Entity.BlockHeader
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &header)
	if err != nil {
		fmt.Println(err)

	}

	return *header
}

//获取钱包
func Getwalletinfo(json string) Entity.Walletinfo {
	var walle Entity.Walletinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &walle)
	if err != nil {
		fmt.Println(err)

	}

	return walle
}

//获取p2p网络的状态 无参数
func GetNetWorkInfo(json string) Entity.Networkinfo {
	var network Entity.Networkinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &network)
	if err != nil {
		fmt.Println(err)

	}

	return network
}

//获取挖矿信息
func GetMiningInfo(json string) Entity.MiningInfo {
	var mining Entity.MiningInfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}

	err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &mining)
	if err != nil {
		fmt.Println(err)

	}

	return mining
}
