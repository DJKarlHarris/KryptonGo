package core

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func (app *App) InitDB() (err error) {

	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/test_db",
		GetConfig().ServerConfig.Db.Passwd,
		GetConfig().ServerConfig.Db.Addr,
		GetConfig().ServerConfig.Db.Port)

	SLOG().Info(dsn)

	app.db, err = sql.Open("mysql", dsn)
	if err != nil {
		SLOG().Errorf("open mysql fail:%s, dsn:%s", err.Error(), dsn)
		return err
	}

	if err = app.db.Ping(); err != nil {
		SLOG().Errorf("db ping fail:%s", err.Error())
		return err
	}

	SLOG().Info("db connect succ.")

	return err
}

func QueryMore(query string, args ...any) (results []map[string]string, err error) {
	rows, err := getApp().db.Query(query, args...)
	if err != nil {
		SLOG().Info("query fail:%s", err.Error())
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		SLOG().Errorf("get rows column fail:%s", err)
		return nil, err
	}

	//可否不用make
	results = make([]map[string]string, 0)
	vals := make([]interface{}, len(cols))
	//pointers := make([]interface{}, len(cols))
	pointers := make([]interface{}, len(cols))

	for i := range vals {
		pointers[i] = &vals[i]
	}

	for rows.Next() {
		if err = rows.Scan(pointers...); err != nil {
			SLOG().Errorf("scan fail, err:%s", err)
			return nil, err
		}

		result := make(map[string]string, 0)

		for i, col := range cols {
			val := vals[i]
			switch v := val.(type) {
			case []byte:
				result[col] = string(v)
			case nil:
				result[col] = ""
			default:
				result[col] = fmt.Sprintf("%v", v)
			}
		}

		results = append(results, result)
	}

	return results, err
}

func (app *App) FreeDb() {
	if app.db != nil {
		app.db.Close()
	}
}
