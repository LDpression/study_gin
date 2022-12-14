package mysql

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=ytf8mb4&pasrseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"))
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		//fmt.Printf("connext db failed,err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return nil
}

func Close() {
	db.Close()
}
