package core

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

func (this *Common) MysqlPingCheck(host, username, password, command string) error {
	en, err := xorm.NewEngine("mysql", username+":"+password+"@tcp("+host+")/mysql?charset=utf8&parseTime=true")
	defer en.Close()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if command == "" {
		if err = en.PingContext(ctx); err != nil {
			return err
		}
		return nil
	} else {
		if result, err := en.Exec(command); err == nil {
			count, _ := result.RowsAffected()
			fmt.Sprintf("rowsAffected number: %d", count)
			return nil
		} else {
			return err
		}
	}

}
