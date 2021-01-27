// Code generated by export-config. DO NOT EDIT.
package conf

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"path"
)

var (
	TestData = &Test{}
	Test2Data = &Test2{}
	Test3Data = &Test3{}
)

func initLoadFunc() {
	DataLoadFunc["test"] = func(dataPath string) {
		fullName := path.Join(dataPath, "test.bytes")
		fmt.Println(fullName)

		f, err := os.Open(fullName)
		panicIfErr(err)
		defer f.Close()

		buf, err := ioutil.ReadAll(f)
		panicIfErr(err)

		err = proto.Unmarshal(buf, TestData)
		panicIfErr(err)
	}

	DataLoadFunc["test2"] = func(dataPath string) {
		fullName := path.Join(dataPath, "test2.bytes")
		fmt.Println(fullName)

		f, err := os.Open(fullName)
		panicIfErr(err)
		defer f.Close()

		buf, err := ioutil.ReadAll(f)
		panicIfErr(err)

		err = proto.Unmarshal(buf, Test2Data)
		panicIfErr(err)
	}

	DataLoadFunc["test3"] = func(dataPath string) {
		fullName := path.Join(dataPath, "test3.bytes")
		fmt.Println(fullName)

		f, err := os.Open(fullName)
		panicIfErr(err)
		defer f.Close()

		buf, err := ioutil.ReadAll(f)
		panicIfErr(err)

		err = proto.Unmarshal(buf, Test3Data)
		panicIfErr(err)
	}

}
func GetTest(id int32) (*Test_Record) {
	if d, ok := TestData.Records[id]; ok {return d} 
	return nil	
}
func GetTest2(id string) (*Test2_Record) {
	if d, ok := Test2Data.Records[id]; ok {return d} 
	return nil	
}
func GetTest3(id int32) (*Test3_Record) {
	if d, ok := Test3Data.Records[id]; ok {return d} 
	return nil	
}
