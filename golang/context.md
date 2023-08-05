## context를 이용한 상태관리

context package? 
>  취소 신호 및 API 경계를 넘어서는 기타 요청 범위 값취소 신호 및 API 경계를 넘어선 기타 요청 범위 그리고 프로세스 사이.

모든 context는 interface를 구현을 요구하고 있다.

``` 
type Context interace {
}
```

```
func main() {
    ctx := context.Background()
    ctx := context.TODO()
}
```

```
withCancel(), withDeadline(), withValue()

```

참고 링크: https://inspirit941.tistory.com/425

##  예제) 
1. 작업 영역에 대한 정의 및 제한 설정
- 웹 서버 전역에 대한 컨텍스트 생성
- 로그인 핸들러에 상위 컨텍스트 전달
- 유저 정보 조회 작업에 대한 타임아웃 컨텍스트 생성
- 레디스 조회

2. 요청 범위 안에서의 데이터 전파
- jwt 이용  

### context.AfterFunc()

3. 작업 종료 후 소멸자 역할
- 간단한 TCP 프로토콜 구성
- 모든 연결은 수시로 연결 되었다가 끊어질 수 있으며, 데이터를 읽을 때 최소 하나 이상의 헤더를 가져야함.
- 헤더 구조체를 sync.Pool을 풀링하면 메모리 할당 조절 가능.

### context.Cause()

4. 같은 작업 영역 안에서의 에러 전파




