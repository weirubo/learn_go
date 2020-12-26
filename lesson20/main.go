package main

import (
	"fmt"
	"reflect"
)

// reflect
// type A int
/* var a A = 1
t := reflect.TypeOf(a)
fmt.Println(t.Name(), t.Kind()) */

/* b := 2
t1, t2 := reflect.TypeOf(b), reflect.TypeOf(&b)
fmt.Println(t1, t2, t1 == t2, t2.Elem())
fmt.Println(t1.Kind(), t2.Kind())
fmt.Println(t1 == t2.Elem()) */

/* type student struct {
	name  string
	score int
} */

/* var s student
t := reflect.TypeOf(&s)
if t.Kind() == reflect.Ptr {
	t = t.Elem()
}

fmt.Println(t.NumField())

for i := 0; i < t.NumField(); i++ {
	structField := t.Field(i)
	fmt.Println(structField.Name, structField.Type)
} */

/* type user struct {
	name string `form:"user_name"`
	age  int    `form:"user_age"`
} */

/* var user1 user
t := reflect.TypeOf(user1)
for i := 0; i < t.NumField(); i++ {
	structField := t.Field(i)
	fmt.Println(structField.Name, structField.Tag.Get("form"))
} */

/* type member struct {
}

func (m member) MemberInfo(name string, score int) string {
	mInfo := fmt.Sprintf("姓名：%s 积分：%d\n", name, score)
	return mInfo
} */

/* var m1 member
v := reflect.ValueOf(&m1)
m := v.MethodByName("MemberInfo")

param := []reflect.Value{
	reflect.ValueOf("frank"),
	reflect.ValueOf(88),
}

result := m.Call(param)
for _, v := range result {
	fmt.Println(v)
} */

/* type person struct{}

func (p person) Hobby(name string, hobby ...interface{}) string {
	myHobby := fmt.Sprintf("I am %v, My hobby is %v", name, hobby)
	return myHobby
} */

/* var p1 person
v := reflect.ValueOf(&p1)
m := v.MethodByName("Hobby")

param := []reflect.Value{
	reflect.ValueOf("lucy"),
	reflect.ValueOf("singing"),
	reflect.ValueOf("dancing"),
}

result := m.Call(param)
fmt.Println(result)

param = []reflect.Value{
	reflect.ValueOf("lucy"),
	reflect.ValueOf([]interface{}{"singing", "dancing"}),
}
result = m.CallSlice(param)
fmt.Println(result) */

type animal struct {
	Name string
	age  int
}

func main() {
	var cat animal
	v := reflect.ValueOf(&cat).Elem()
	fmt.Printf("v:%v type:%v kind:%v\n", v, v.Type(), v.Kind())
	v.FieldByName("Name").SetString("kity")
	fmt.Printf("v:%v type:%v kind:%v\n", v, v.Type(), v.Kind())
}
