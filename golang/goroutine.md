### > go routine의 기본적인 예시 

```
package main

import (
	"fmt"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
	wg.Done()  --> 하나의 go routine이 끝나면 하나씩 cnt 감소 
}

func main() {

	1. wg := new(sync.WaitGroup) -> waitGroup 객체 생성 (선언 & 할당)
    2. var wg = &sync.WaitGroup{} -> waitGroup 객체 생성 (선언 & 할당)
    3. var wg *sync.WaitGroup -> waitGroup 객체 생성 (선언)

	wg.Add(4)  -> 3번으로 선언만 한 뒤에 Add 함수를 호출하게 되면 
    "Invalid memory address or nil pointer dereference" 에러를 맛볼 수 있음.
    (실제로 포인터 값을 할당받지 못했기 때문에 참조 할 수 없다는 에러가 나옴.)

	// 함수를 비동기적으로 실행
	go say("Async1", wg)
	go say("Async2", wg)
	go say("Async3", wg)

	wg.Wait() --> 다른 go routine 들이 끝날 때까지 대기
}
```

### > for loop vs for loop in goroutine 
```
func withoutGoroutine() {
	startTime := time.Now()

	for i := 0; i < 100000; i++ {
		// Do some processing here...
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Without Goroutine - Elapsed Time:", elapsedTime)
}

func withGoroutine() {
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			// Do some processing here...
			wg.Done()
		}(i)
	}
	wg.Wait()

	elapsedTime := time.Since(startTime)
	fmt.Println("With Goroutine - Elapsed Time:", elapsedTime)
}

func main() {
	withoutGoroutine()
	withGoroutine()
}

Without Goroutine - Elapsed Time: 144.101µs
With Goroutine - Elapsed Time: 4.5061035s

찾아보니 복잡한 로직일 수록 goroutine의 시간효율성이 올라간다고 한다. 
반대로 말하면 간단한 로직일 수록 goroutine을 사용하는 것이 더 빠르다.
```

### > 동시 접근
```
func main() {
	start := time.Now()
	size := 1000000
	arr := make([]int, 0, size)

	var w sync.WaitGroup
	var m sync.Mutex

	w.Add(size)
	for i := 0; i < size; i++ {
		go func(data int) {
			defer w.Done()
			m.Lock()
			arr = append(arr, data)
			m.Unlock()
		}(i)
	}
	w.Wait()

	if size == len(arr) {
		log.Println(size, time.Since(start))
	} else {
		log.Fatal()
	}
}

하나의 arr를 가지고 동시에 접근해 각 go routine이 가지고 있는 값을 인덱스 주소에 넣다보니 값이 들어가지 않게 된다. 실제로 mutex를 사용하지 않게 되면 size가 1000000이 아니라 999555등 근접한 값을 갖게 된다.
```

### > buffer 
버퍼를 이용해 값을 담아놓고 arr에 담는 방식이 있다. 
```
func main() {
	size := 1000000
	start := time.Now()
	arr := make([]int, 0, size)
	bufferChannel := make(chan int, size) // buffer channel

	var w sync.WaitGroup
	w.Add(size)
	for i := 0; i < size; i++ {
		go func(data int) {
			defer w.Done()
			bufferChannel <- data // 값을 담을 때에는 <-을 사용한다.
		}(i)
	}
	w.Wait()

	close(bufferChannel) <- close를 안하게 되면 deadlock 발생 위험이 있다.
	for data := range bufferChannel {
		arr = append(arr, data)
	}

	if size == len(arr) {
		log.Println(size, time.Since(start))
	} else {
		log.Fatal()
	}
}
```
