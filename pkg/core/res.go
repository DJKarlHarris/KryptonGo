package core

import (
	"KryptonGo/pkg/luban"
	"KryptonGo/pkg/res"
)

func (app *App) InitRes() {
	var err error

	if app.tables, err = res.NewTables(luban.Loader); err != nil {
		println(err.Error())
		return
	}
}

func GetResTable() *res.Tables {
	return app.tables
}
