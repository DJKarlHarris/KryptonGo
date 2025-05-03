package main

import (
	"KryptonGo/pkg/luban"
	"KryptonGo/pkg/res"
	"fmt"
)

func main() {
	var tables *res.Tables
	var err error

	//load table
	if tables, err = res.NewTables(luban.Loader); err != nil {
		println(err.Error())
		return
	}

	for idx, reward := range tables.TbReward.GetDataList() {
		fmt.Printf("idx:%d id:%d name:%s count:%d\n", idx, reward.Id, reward.Name, reward.Count)
	}

}
