package split

import "strings"
//匹配单个字符
func Split(str, sep string) (result []string) {
	result = make([]string, 0, strings.Count(str, sep)+1)
	i := strings.Index(str, sep)
	for i > -1 {
		result = append(result, str[:i])
		str = str[i+1:]
		i = strings.Index(str, sep)
	}
	result = append(result, str)
	return
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
