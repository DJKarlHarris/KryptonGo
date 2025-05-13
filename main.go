package main

import (
	"KryptonGo/pkg/core"
	"fmt"
	"time"
)

func main() {
	var err error

	//core init
	if err = core.Init(); err != nil {
		println(err.Error())
		return
	}

	//test
	for idx, reward := range core.GetResTable().TbReward.GetDataList() {
		fmt.Printf("idx:%d id:%d name:%s count:%d\n", idx, reward.Id, reward.Name, reward.Count)
	}

	var url = "www.4399.com"

	//core.SLOG().Infow("test",
	//	"url", url,
	//	"time", 3,
	//	"backoff", time.Second,
	//)

	for i := 0; i < 1000; i++ {
		core.SLOG().Infof("test url:%s time:%d backoff:%s", url, 3, time.Second)
	}

	core.Exit()
}
