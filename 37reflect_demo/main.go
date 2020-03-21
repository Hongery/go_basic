package main

import (
"fmt"
"reflect"
)
type Cat struct{

}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)//返回类型int64/float32
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("type:%v\n", v)//返回的是值3.14/100
	k := v.Kind() //值得种类
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

//通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	var c =Cat{}
	//Type
	reflectType(a) // type:float32
	reflectType(b) // type:int64
	reflectType(c) // type:main.Cat
	//Value
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	reflectValue(c) // type:{}为空
	//设置值
	reflectSetValue1(b)  //传值，传指针 不会修改
	fmt.Println(b)	//100
	reflectSetValue2(b)  //得调用Elem()方法设置
	fmt.Println(b)	//200 修改成功


}
