package main

import (
	"Btcconnection/Server"
	"Btcconnection/Utill"
	"Btcconnection/entity"
	"fmt"
	"github.com/mapstructure-master"
)

func Getblockcount() {
	//getbalances
	requeststr := Utill.PrepareJSON("getbalances", nil)
	result := Utill.Dopost(entity.RPCURL, Server.Createmap(), requeststr)
	if result.Code == 200 {
		resultadss := *(result.Data)
		b:=resultadss.Result
		a:=entity.Getbalances{}

		err:= mapstructure.WeakDecode(a,b)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(a.Mine.Trusted)
	}
}