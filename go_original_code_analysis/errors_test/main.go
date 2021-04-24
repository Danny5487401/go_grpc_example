package main

/* errors 模块
1。 errors.Is 判断是否a错误是否是b错误的后代
2。 errors.Unwrap 将a错误的包装剔除一层
3。 errors.As 将a错误一直剔除到错误类型为 B 类型为止
4。 fmt.Errorf("%w", err) 将err错误包装一层
 */

import (
	"errors"
	"fmt"
)

func doSomethingWrong(o error) error {
	return fmt.Errorf("%w wrapped error", o)
}

func main() {
	// o: original
	// n: new error
	o := errors.New("original error")

	fmt.Printf("error: %s\n", o)
	n := doSomethingWrong(o)
	fmt.Printf("error: %s\n", n)

	fmt.Printf("n is o: %t\n", errors.Is(n, o))
}
