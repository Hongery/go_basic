package main
import(
	"fmt"
)
 func main(){
	 var s1 = []int{}   //声明切片并进行初始化
	 fmt.Println(s1==nil ) //false

	 var s2 []int //声明int类型切片，并未初始化  
	 fmt.Println(s2==nil)//true

	 var s3 =[]int{1,2,3}
	fmt.Println(len(s3),cap(s3))//3 3
	s4 :=s3[1:2]//左包含右不包含
	fmt.Println(s4)
	s5 :=s3[:]//数组切片，全部的数据 从0开始
	fmt.Println(s5)

	var s6 = []int{1,2,3,4,5,6,7}
	s7 :=s6[1:4]
	fmt.Printf("len(s7)%d   cap(s7)%d",len(s7),cap(s7)) //3 ,6
	 //var s3 = [2]int{} 
	//  fmt.Println(s3==nil)

 }