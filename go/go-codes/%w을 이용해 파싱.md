> 언제 %w verb를 사용해야 할까? 
==> 함수안에 뎁스가 많아서 어느 뎁스에서 에러가 발생했는지 알기 쉽게 해주는 연산자이다. 따라서 정밀한 tracing을 하고 싶을 때 사용한다.

> Go 1.13부터 사용가능
```
예제

 package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	result := a / b
	return result, nil
}

func process(a, b int) error {
	_, err := divide(a, b)
	if err != nil {
		return fmt.Errorf("failed to process division: %w", err) // 에러가 발생한 시점에 %w을 이용해 wrapping 해준다.
	}
	return nil
}

func main() {
	a, b := 10, 0
	if err := process(a, b); err != nil {
		fmt.Printf("Error: %v\n", err) // 값으로 이용하는 부분에 %v을 이용한다.
	}
}

// "github.com/pkg/errors" 패키지를 이용하면 wrap을 더 잘 이용할 수 있다.
// 참고자료: https://ssup2.github.io/programming/Golang_Error_Wrapping/
```
