package core

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (app *App) InitDB() (err error) {

	dsn := "root:123456@tcp(192.168.1.3:3306)/test"

	app.db, err = sql.Open("mysql", dsn)
	if err != nil {
		SLOG().Errorf("open mysql fail:%s", err.Error())
		return err
	}

	if err = app.db.Ping(); err != nil {
		SLOG().Errorf("db ping fail:%s", err.Error())
		return err
	}

	SLOG().Info("db connect succ.")

	return err
}

type testtb struct {
	id int
}

func QueryMore() (err error) {
	rows, err := getApp().db.Query("SELECT * FROM testtb")
	if err != nil {
		SLOG().Info("query fail:%s", err.Error())
		return err
	}
	defer rows.Close()

	for rows.Next() {
		//var test testtb
		values := make([]interface{}, 1)
		values[0] = new(interface{})

		rows.Scan(values...)

		SLOG().Infof("id:%v", *(values[0].(*interface{})))
	}

	return err
}
