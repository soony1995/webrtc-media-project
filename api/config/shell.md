> "" 와 ''의 차이

### 1. 변수 확장 
    - '' 의 경우     
    # 예시
        USER="John"
        echo 'Hello $USER'  # 출력: Hello $USER    
    ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ
    - "" 의 경우
     # 예시
        USER="John"
        echo "Hello $USER"  # 출력: Hello John
    ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ        
    
### 2. 특수 문자 해석    
    추가적으로, 큰따옴표를 사용하면 몇 가지 특수 문자들이 해석됩니다. 예를 들어, \n은 줄 바꿈 문자로 해석되고, \"는 큰따옴표 문자 자체를 나타냅니다.    

### 3. ssh 
    작은따옴표(')로 둘러싼 부분은 로컬 컴퓨터에서 해당 스크립트를 실행하는 쉘에게서 변수 확장을 비활성화하기 때문에 원격 호스트로 전송되지 않습니다. 따라서 작은따옴표로 둘러싸인 스크립트는 원격 호스트에서 실행되지 않고, 로컬 컴퓨터에서만 실행됩니다. 이것은 원격 호스트로 명령어를 전달할 때 주의해야 하는 중요한 사항입니다. 만약 원격 호스트에서 스크립트를 실행해야 한다면, 해당 스크립트를 큰따옴표(")로 둘러싸야 합니다. 큰따옴표로 둘러싼 스크립트는 원격 호스트로 전달되고 원격 호스트에서 실행됩니다.
    ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ

    '' 을 사용한 경우:
    ssh user@remote_host 'echo "Hello $USER"' --> 값이 전달되지 않음.
    ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ
    
    "" 을 사용한 경우:
    ssh user@remote_host "echo \"Hello \$USER\"" --> 값이 전달 됨.