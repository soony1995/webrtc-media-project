> 콜백 함수

다른 함수에 인수로 전달되는 함수로, 특정 작업을 완료하기 위해 콜백 함수를 실행해야 할 떄마다 콜백 함수를 실행.

- 동기 콜백
- 비동기 콜백

> 사용법
```
package main

import "fmt"

func main() {

	calc := func(a int, b int) int {
		return a + b
	}
    
	fmt.Println(calc)	//output : 0x467fa0
	fmt.Println(calc(1, 2))	//output : 3
}
```

