package bitcoin_server

import (
	"BitcoinConnection/Bitcoin_conne"
	"BitcoinConnection/models"
	"strings"
)

var prejson models.Prejson

/*获取区块数量*/
func GetBlockCount() interface{} {
	json, err := Bitcoin_conne.PrepareJSON("getblockcount", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}

/*获取帮助*/
func GetHelp() interface{} {
	json, err := Bitcoin_conne.PrepareJSON("help", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}

/*退出*/
func GetStop() interface{} {
	json, err := Bitcoin_conne.PrepareJSON("stop", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}

/*获取难度值*/
func GetDifficult() interface{} {
	json, err := Bitcoin_conne.PrepareJSON("getdifficulty", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}

/*获取钱包*/
func GetWalletInfo() interface{} {
	json, err := Bitcoin_conne.PrepareJSON("getwalletinfo", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}


/*获取（计算机等的）正常运行时间*/
func GetUptime()interface{}  {
	json, err := Bitcoin_conne.PrepareJSON("uptime", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}

/*获取矿工信息*/
func GetMiningInfo()interface{}  {
	json, err := Bitcoin_conne.PrepareJSON("getmininginfo", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}


/*获取区块状态*/
func GetBlockStatus()interface{}  {
	json, err := Bitcoin_conne.PrepareJSON("getblockstatus", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}


/*生成新的比特币地址*/
func GetNewAddress()interface{}  {
	json, err := Bitcoin_conne.PrepareJSON("getnewaddress", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}


/*根据区块hash获取区块信息*/
func GetBlock()interface{}  {
	json, err := Bitcoin_conne.PrepareJSON("getblock", nil)
	if err != nil {
		return ""
	}
	var RPCresult *models.Rpcresult
	RPCresult, err = Bitcoin_conne.Dopost(Bitcoin_conne.RPCURL, Bitcoin_conne.CreateMap(), strings.NewReader(json))
	if err != nil {
		return nil
	}
	return RPCresult.Data.Resultbit
}