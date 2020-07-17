package utils

import (
	"fmt"
	"monkeyClient/logUtils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func SqlxCli() *sqlx.DB {

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", Config.DbUser,Config.DbPass,Config.DbHost,Config.DbName)
	if dbClient, err := sqlx.Connect("mysql", dns); err != nil{
		logUtils.Errorf("mysql connect error=%v",err)
	}else {
		return dbClient
	}
	return nil
}
