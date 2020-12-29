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
	Map["getblockcount"] = GetBlockCount(json)         //无参数
	Map["getbestblockhash"] = GetBestBlockHash(json)   //无参数
	Map["getblockhash"] = GetBlockHash(json)           // 高度 int
	Map["getblock"] = GetBlock(json)                   //区块hash
	Map["getnewaddress"] = GetNewAddress(json)         //无参数
	Map["getblockchaininfo"] = GetBlockChainInfo(json) //无参数
	Map["getdifficulty"] = GetDifficulty(json)         //无参数
	Map["uptime"] = UpTime(json)                       //无参数
	Map["validateaddress"] = ValidateAddress(json)     //btc地址
	Map["getbalances"] = GetBalances(json)             //无参数
	Map["getmemoryinfo"] = GetMemoryInfo(json)         //无参数
	Map["gettxoutsetinfo"] = Txoutset(json)            //无参数
	Map["getmempoolinfo"] = MemPoolInfo(json)          //无参数
	Map["getchaintxstats"] = ChainTxStats(json)        //（可选参数 nblocks int 第几个区块 blockhash 块的散列hash ）
	Map["getblockheader"] = BlockHeader(json)          //(blockhash 如果verbose是false,则返回一个序列化的、hex编码数据的字符串,用于块头“散列”。
	//如果verbose是true,则返回一个带有块头消息的对象)
	Map["getwalletinfo"] = Getwalletinfo(json)   //无参数
	Map["getnetworkinfo"] = GetNetWorkInfo(json) //无参数
	Map["getmininginfo"] = GetMiningInfo(json)   //无参数
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
		if err != nil {
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
		if err != nil {
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

//获取关于给定比特币地址的信息
func ValidateAddress(json string) *Entity.Validata {
	var block Entity.Validata
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println("", err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &block)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return &block
	}
	fmt.Println("", err)
	return nil
}

//获取一个在BTC中包含所有余额的对象
func GetBalances(json string) *Entity.Balances {
	var balance Entity.Balances
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &balance)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//getmemoryinfo  获取一个包含内存使用信息的对象 无参数
func GetMemoryInfo(json string) *Entity.Locked {
	var locked Entity.Locked
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &locked)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//gettxoutsetinfo  获取关于未使用事务输出集的统计信息 无参数
func Txoutset(json string) *Entity.Txoutsetinfo {
	var txout Entity.Txoutsetinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &txout)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//getmempoolinfo 获取内存池的详细信息 无参数
func MemPoolInfo(json string) *Entity.Mempoollinfo {
	var mempool Entity.Mempoollinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &mempool)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//getchaintxstats 计算链中交易总数和费率的统计数字 （可选参数 nblocks int 第几个区块 blockhash 块的散列hash ）
func ChainTxStats(json string) *Entity.Chaintxstats {
	var chain Entity.Chaintxstats
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &chain)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//getblockheader (blockhash 如果verbose是false,则返回一个序列化的、hex编码数据的字符串,用于块头“散列”。
//如果verbose是true,则返回一个带有块头消息的对象)
func BlockHeader(json string) *Entity.BlockHeader {
	var header Entity.BlockHeader
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &header)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//获取钱包
func Getwalletinfo(json string) *Entity.Walletinfo {
	var walle Entity.Walletinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &walle)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//获取p2p网络的状态 无参数
func GetNetWorkInfo(json string) *Entity.Networkinfo {
	var network Entity.Networkinfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &network)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}

//获取挖矿信息
func GetMiningInfo(json string) *Entity.MiningInfo {
	var mining Entity.MiningInfo
	RpcResult, err := Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), json)
	if err != nil {
		fmt.Println(err)
	}
	if RpcResult.Code == http.StatusOK {
		err = mapstructure.WeakDecode(RpcResult.Data.Resultbit, &mining)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(err)
	return nil
}
