package models

import (
	"Bitcoin_core_web/dbMysqlSV"
	"fmt"
)

type User struct {
	Id           int  `form:"id"`
	UserName     string `form:"userName"`
	UserPassword string `form:"password"`
	PhoneNumber  string `form:"phone_number"`
	//Email        string `form:"userEmail"`
}

func (u User)Quereuser()(*User,error){

    row:=dbMysqlSV.Db.QueryRow("select id,user_name ,user_password,user_phone from bccw_user where user_name=? and user_password =?",u.UserName,u.UserPassword)
    err:= row.Scan(&u.Id,&u.UserName,&u.UserPassword,&u.PhoneNumber)
	if err != nil {
		return nil,err
	}
    return &u,nil

}

func (u User) Inseruser()(int64,error){
	result,err:=dbMysqlSV.Db.Exec("insert into bccw_user(user_name ,user_password,user_phone ) values(?,?,?) ",u.UserName,u.UserPassword,u.PhoneNumber)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	 id,err:=result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id,err
}