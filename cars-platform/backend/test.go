package main

import (
	"fmt"
	"reflect"
)

type Rsp struct {
	code *int64
}

func (rsp *Rsp) GetCode() {
	fmt.Println("type:", reflect.TypeOf(*(rsp.code)))
	fmt.Println("value:", *(rsp.code))
}

func (rsp *Rsp) SetCode(v int64) {
	rsp.code = &v
	fmt.Println(" set code success -> ", v)
}

func Reflect(rsp interface{}, code int64) {
	refRsp := reflect.ValueOf(rsp)
	methodValue := refRsp.MethodByName("SetCode")
	args := []reflect.Value{reflect.ValueOf(code)}
	methodValue.Call(args)

}

func main() {
	a := int64(4)
	rsp := &Rsp{code: &a}
	fmt.Println("before")
	rsp.GetCode()
	Reflect(rsp, 10001)
	fmt.Println("after")
	rsp.GetCode()

}
