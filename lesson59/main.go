package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func main() {
	// 创建 Engine
	engine, err := xorm.NewEngine("mysql", "root:root@/example?charset=utf8")
	defer func() {
		err = engine.Close()
		if err != nil {
			fmt.Printf("engine close err=%v\n", err)
			return
		}
	}()
	if err != nil {
		fmt.Printf("init xorm engine fail, err=%v\n", err)
		return
	}

	// 更改数据
	sql := "UPDATE example SET title=?, view=?, created=? WHERE id=?"
	res, err := engine.Exec(sql, "Python", 60, time.Now().Unix(), 2)
	if err != nil {
		fmt.Printf("Update err=%v\n", err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected err=%v\n", err)
		return
	}
	fmt.Printf("affected=%d\n", affected)

	//example := &Example{
	//	Title: "JavaScript",
	//	View:  98,
	//}
	//
	//condi := &Example{
	//	Id: 2,
	//}
	//
	//affected, err := engine.Update(example, condi)
	//if err != nil {
	//	fmt.Printf("Update err=%v\n", err)
	//	return
	//}
	//fmt.Printf("affected=%d\n", affected)

	// 插入数据
	//example := &Example{
	//	Title: "Vue",
	//	View:  85,
	//}
	//affected, err := engine.Insert(example)
	//if err != nil {
	//	fmt.Printf("Insert err=%v\n", err)
	//	return
	//}
	//id := example.Id
	//fmt.Printf("affected=%v || id=%d\n", affected, id)

	// 更新数据
	//example := &Example{
	//	Title: "Java",
	//	View:  0,
	//}
	//condi := &Example{
	//	Id: 2,
	//}
	////affected, err := engine.Update(example, condi)
	//affected, err := engine.Cols("title", "view").Update(example, condi)
	//if err != nil {
	//	fmt.Printf("Update err=%v\n", err)
	//	return
	//}
	//fmt.Printf("affected=%d\n", affected)

}

type Example struct {
	Id      int    `json:"id" form:"id" xorm:"autoincr"`
	Title   string `json:"title" form:"title"`
	View    int    `json:"view" form:"view"`
	Created int    `json:"created" form:"created" xorm:"created"`
	Updated int    `json:"updated" form:"updated" xorm:"updated"`
}
