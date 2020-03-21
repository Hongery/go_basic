package main
//上台阶面试问题，那个台阶，一次可以走一步，也可以走两步，有多少种走法
import(
	"fmt"
)
func taijie(n uint64) uint64{
	if n ==1 {
		return 1
	}
	if n==2 {
		return 2
	}
	return taijie(n-1)+taijie(n-2)
}
func main(){
	ret :=taijie(3)
	fmt.Println(ret)
	re :=taijie(10)
	fmt.Println(re)
}