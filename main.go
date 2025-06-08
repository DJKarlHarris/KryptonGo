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

func test2(args []any) []any {
	for _, v := range args {
		if v == nil {
			break
		}

		t := reflect.TypeOf(v)
		core.SLOG().Infof("%s %s value:%v", t.Name(), t.Kind(), v)
	}

	args = append(args, 10)

	return args
}

func main() {
	var err error

	//core init
	if err = core.Init(); err != nil {
		println(err.Error())
		return
	}

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

	//test cfg
	core.SLOG().Infof("test:%d", core.GetConfig().ServerConfig.Test)

	for _, value := range core.GetConfig().ServerConfig.Table {
		core.SLOG().Infof("%s", value)
	}

	//test db
	core.QueryMore()

	argsAny := make([]any, 0, 5)
	argsAny = append(argsAny, 1)
	argsAny = append(argsAny, "2")
	argsAny = append(argsAny, 1.2)

	//test(argsAny...)
	argsAny = test2(argsAny)

	for _, v := range argsAny {
		core.SLOG().Infof("%v", v)
	}

	//var y interface{} = nil
	//z := y.(int)
	//fmt.Printf("%+v", z)

	core.Exit()
}
