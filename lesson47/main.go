package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

func main() {
	// 三方返回普通 json 字符串
	// jsonRes := `{
	// 	"Id": 1001,
	// 	"Name": "frank"
	// }`
	// data := new(User)
	// err := json.Unmarshal([]byte(jsonRes), &data)
	// if err != nil {
	// 	log.Printf("json Unmarshal err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("data=%+v", data)

	// data2 := make(map[string]interface{})
	// err := json.Unmarshal([]byte(jsonRes), &data2)
	// if err != nil {
	// 	log.Printf("json Unmarshal err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("data2=%+v", data2)

	// 三方返回嵌套 json 字符串
	jsonRes := `{
	"Id": 1001,
	"Name": "frank",
	"Details": {
	"Gender": "man",
	"Age": 18,
	"Phone": "13800138000",
	"address": "Beijing"
	}
	}`
	// data := new(User)
	// err := json.Unmarshal([]byte(jsonRes), &data)
	// if err != nil {
	// 	log.Printf("json Unmarshal err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("data=%+v\n", data)

	// data := make(map[string]interface{})
	// err := json.Unmarshal([]byte(jsonRes), &data)
	// if err != nil {
	// 	log.Printf("json Unmarshal err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("data=%+v\n", data)
	// fmt.Printf("age=%v\n", data["Details"]["Age"])

	tmpData := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonRes), &tmpData)
	if err != nil {
		log.Printf("json Unmarshal err:%v\n", err)
		return
	}
	data2 := new(User)
	err = mapstructure.Decode(tmpData, data2)
	if err != nil {
		log.Printf("decode err:%v\n", err)
		return
	}
	fmt.Printf("data2=%+v\n", data2)
	fmt.Printf("age=%v\n", data2.Details.Age)
}

type User struct {
	Id      int
	Name    string
	Details Details
}

type Details struct {
	Gender  string
	Age     interface{}
	Phone   string
	Address string
}
