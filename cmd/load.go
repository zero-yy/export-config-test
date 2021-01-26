package main

import (
	"fmt"
	"github.com/zero-yy/export-config-test/test/go/gen/conf"
)

func main() {
	conf.MustLoad("./test/go/conf_data")
	fmt.Println(conf.TestData.Records[2].TdoubleAry2S)
}
