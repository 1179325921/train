package main

import (
	"errors"
	"fmt"
	"wenhui.yu/models"

	"reflect"
	"strconv"
)

//反射方法
//获取字段信息
func toReflectionField(o interface{}) {
	fmt.Println("======  toReflectionField  ======")
	// 获得目标对象信息
	t := reflect.TypeOf(o)
	//t.Name 获取类型名称
	fmt.Printf("Type: %s \n", t.Name())

	v := reflect.ValueOf(o)
	fmt.Printf("All value: %v \n", v)
	fmt.Println("Fields: ")
	//通过t.NumField() 来获取索引
	for i := 0; i < t.NumField(); i++ {
		//通过索引来获取字段
		f := t.Field(i)
		//获取字段的值
		if v.Field(i).CanInterface() {
			value := v.Field(i).Interface()
			fmt.Printf("Field:%s Value:%v \n", f.Name, value)
		}
	}
}

//反射方法
//获取方法信息
//执行方法
func toReflectionMethod(o interface{}) {
	fmt.Println("======  toReflectionMethod  ======")
	// 获得目标对象信息
	t := reflect.TypeOf(o)
	//t.Name 获取类型名称
	fmt.Printf("Type: %s \n", t.Name())

	//通过t.NumMethod() 来获取索引
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("Method infomation :\n Name: %v, Type: %s \n", m.Name, m.Type)
		fmt.Println("============")
		//	执行方法
		rMethod := reflect.ValueOf(o).MethodByName(m.Name)
		args := make([]reflect.Value, 0)
		rMethod.Call(args)
		fmt.Println("============")
	}
}

//判断类型
func allowReflectKind(o interface{}, kind reflect.Kind) bool {
	fmt.Println("======  allowReflectKind  ======")
	t := reflect.TypeOf(o)
	reflect.ValueOf(o).Index(0).Interface()
	if t.Kind() != kind {
		fmt.Printf("Kind illegality,need %s real %s \n", kind, t.Kind())
		return false
	}
	return true
}

func IsInArray(arr interface{}, ele interface{}) bool {
	v := reflect.ValueOf(arr)
	for i := 0; i < v.Len(); i++  {
		if v.Index(i).Interface() == ele {
			return true
		}
	}

	return false
}

func Contain(target interface{}, obj interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

/*func IsInArray(arr interface{}, ele interface{}) bool {
	intArr, ok := arr.([]interface{})
	if !ok || 0 >= len(intArr) {
		return false
	}
	for _, v := range intArr {
		if ele == v {
			return true
		}
	}
	return false
}*/

//判断类型
//强制类型转换
//执行该类型的方法
func transformClass(o interface{}) {
	fmt.Println("======  transformClass  ======")
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	if t.Name() == "User" {
		user := v.Interface().(models.User)
		user.Hello()
		fmt.Println(user)
	}
}

// 强制类型转换（获取指针）
// 该方法中不会修改指针地址
func transformClassPointer(o interface{}) {
	fmt.Println("======  transformClassPointer  ======")
	reflectType := reflect.TypeOf(o).Elem()
	reflectValue := reflect.ValueOf(o).Elem()
	if reflectType.Name() == "GoodProgrammer" {
		pointer, _ := reflectValue.Interface().(models.GoodProgrammer)
		fmt.Printf("pointer: %x \n", pointer)
		fmt.Println("Name:" + pointer.Name)
		fmt.Println("Age:" + strconv.Itoa(pointer.Age))
		pointer.Run()
		fmt.Printf("Salary: %d \n", pointer.Work(10, 1000))
	}
}

func reflectInterface(person models.Person) {
	fmt.Println("======  reflectInterface  ======")
	person.Run()
	fmt.Printf("Salary: %d \n", person.Work(10, 1000))
}

// 通过反射改变字段的值
// 只有指针才能修改
func reflectChangeValue(in interface{}, filedName string, afterValue reflect.Value) {
	fmt.Println("======  reflectChangeValue  ======")
	valueOf := reflect.ValueOf(in).Elem()
	typeOf := reflect.TypeOf(in).Elem()
	//reflect.Ptr 指针类型
	if valueOf.Kind() != reflect.Ptr && !valueOf.CanSet() {
		fmt.Printf("Con't reset filed：%s value %v \n", filedName, afterValue)
		return
	}
	fmt.Printf("Orgianl:%v \n", in)
	for i := 0; i < valueOf.NumField(); i++ {
		if typeOf.Field(i).Name == filedName {
			fmt.Printf("AfterValue:%v \n", afterValue)
			valueOf.Field(i).Set(afterValue)
		}
	}
	fmt.Printf("Complete:%v \n", in)
}

type User struct {
	Name string
	Age  int
	// 私有属性，对外不可见
	sex bool
}

func (myself User) SayHello() {
	fmt.Println("Hello, my name is " + myself.Name + ", age is " + strconv.Itoa(myself.Age))
}

func main() {

	user1 := User{Name: "XiaoMing", Age: 18}
	fmt.Println(user1)
	toReflectionField(user1)
	//toReflectionMethod(user1)
	obj := []string{"a", "b", "c", "d", "e"}
	b, _ := Contain(obj, "e")
	fmt.Println(b)



	//ints := []int{1, 2, 4, 5}
	//strs := []string{"aa", "bb", "cc"}
	//var a int64
	//a = 1
	//fmt.Println(IsInArray(ints, a))
	//fmt.Println(IsInArray(ints, 1))
	//fmt.Println(IsInArray(ints, 10))
	//fmt.Println(IsInArray(strs, "aa"))
	//fmt.Println(IsInArray(strs, "aaa"))

	//user := models.User{"wenhui", 1, 18}
	//person := models.GoodProgrammer{"zhangsan", 1}
	//
	//toReflectionField(user)
	//
	//toReflectionMethod(user)
	//
	//if allowReflectKind(user, reflect.Struct) {
	//	fmt.Println("Kind right")
	//}
	//
	//transformClass(user)
	//
	//fmt.Println(reflect.TypeOf(person).Name())
	//
	//toReflectionField(person)
	//
	//transformClassPointer(&person)
	//
	//reflectInterface(person)
	//
	//reflectChangeValue(&person, "Name", reflect.ValueOf("lisi"))
	//
	//reflectChangeValue(&person,"Age",reflect.ValueOf(2))
}
