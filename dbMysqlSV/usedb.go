package dbMysqlSV

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)
 var   Db *sql.DB
func ConDB()  {
	config:=beego.AppConfig
	dbDriveername:=config.String("db_dirveerName")
	dbUser:=config.String("db_user")
	dbPassword:=config.String("db_password")
	dbip:=config.String("db_ip")
	dbName:=config.String("db_name")
	condbUrl:=dbUser+":"+dbPassword+"@tcp"+dbip+"/"+dbName+"?charset=utf8"
	DBtype,err:=sql.Open(dbDriveername,condbUrl)
	if err != nil {
		panic("数据库连接错误")
	}
   Db=DBtype
   fmt.Println("数据库连接成功")
}
