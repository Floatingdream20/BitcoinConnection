package main

import (
	"BitcoinConnection/bitcoin_server"
	"BitcoinConnection/dbMysqlSV"
	_ "BitcoinConnection/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	dbMysqlSV.ConDB()

	//方法

	//1.获取区块数量
	GetBlockCount:=bitcoin_server.GetBlockCount()
	fmt.Println("获取区块数量",GetBlockCount)
    //2.获取帮助
    GetHelp:=bitcoin_server.GetHelp()
    fmt.Println("获取帮助",GetHelp)
    ////3.退出
    //GetStop:=bitcoin_server.GetStop()
    // defer fmt.Println("退出",GetStop)
    //4.获取当前难度值
    GetDifficult:=bitcoin_server.GetDifficult()
    fmt.Println("获取当前难度值",GetDifficult)
    //5.获取钱包
	GetWalletInfo:=bitcoin_server.GetWalletInfo()
	fmt.Println("获取钱包",GetWalletInfo)
	//6.获取（计算机等的）正常运行时间
    GetUptime:=bitcoin_server.GetUptime()
	fmt.Println("正常运行时间",GetUptime)
    //7.获取矿工信息
	GetMiningInfo:=bitcoin_server.GetMiningInfo()
    fmt.Println("获取矿工信息",GetMiningInfo)
	//8.获取区块状态
	GetBlockStatus:=bitcoin_server.GetBlockStatus()
	fmt.Println("获取区块状态",GetBlockStatus)
	//9.生成新的比特币地址
	GetNewAddress:=bitcoin_server.GetNewAddress()
	fmt.Println("生成新的比特币地址",GetNewAddress)
    //10.根据区块hash获取区块信息
    GetBlock:=bitcoin_server.GetBlock()
    fmt.Println("根据区块hash获取区块信息",GetBlock)


    
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
