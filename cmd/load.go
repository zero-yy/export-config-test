package main

import (
	"fmt"
	"github.com/zero-yy/export-config-test/test/go/gen/conf"
)

func main() {
	conf.MustLoad("./test/go/conf_data")

	// 两种方式均可
	fmt.Println(conf.TestData.Records[2].TdoubleAry2S)
	fmt.Println(conf.GetTest(2).TdoubleAry2S)
}
