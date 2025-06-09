package main

import (
	"KryptonGo/pkg/core"
	"fmt"
	"reflect"
)

func test(args ...any) {
	for _, v := range args {
		if v == nil {
			break
		}

		t := reflect.TypeOf(v)
		core.SLOG().Infof("%s %s value:%v", t.Name(), t.Kind(), v)
	}
}

func main() {
	var err error

	//core init
	if err = core.Init(); err != nil {
		println(err.Error())
		return
	}
	defer core.Exit()

	//test table
	for idx, reward := range core.GetResTable().TbReward.GetDataList() {
		fmt.Printf("idx:%d id:%d name:%s count:%d\n", idx, reward.Id, reward.Name, reward.Count)
	}

	//var url = "www.4399.com"

	//core.SLOG().Infow("test",
	//	"url", url,
	//	"time", 3,
	//	"backoff", time.Second,
	//)

	//for i := 0; i < 1000; i++ {
	//	core.SLOG().Infof("test url:%s time:%d backoff:%s", url, 3, time.Second)
	//}

	//test db
	results, err := core.QueryMore("SELECT * FROM test_tb")
	if err != nil {
		core.SLOG().Errorf("query fail:%s", err)
		return
	}

	for _, result := range results {

		id, ok := result["id"]
		if !ok {
			core.SLOG().Info("id is not found")
		}

		ext, ok := result["ext"]
		if !ok {
			core.SLOG().Info("ext is not found")
		}

		core.SLOG().Infof("id:%s ext:%s", id, ext)
	}

	//var y interface{} = nil
	//z := y.(int)
	//fmt.Printf("%+v", z)

}
