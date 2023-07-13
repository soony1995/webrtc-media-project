## C:\Windows\System32\drivers\etc\hosts    

> 우리가 도메인을 리졸빙을 할 때, hosts에 적힌 도메인과 ip가 있다면 hosts의 ip로 접속하게 됩니다.  즉, hosts 파일에서 찾아보고 없으면 DNS에 질의를 하게 됩니다.

### 알게 된 계기: 회사의 DNS 서버만 내려가 리졸빙을 못해 IP를 못가져오는 상황이였다. 따라서 수동으로 IP를 적어줌으로써 해결.
```
# Copyright (c) 1993-2009 Microsoft Corp.
#
# This is a sample HOSTS file used by Microsoft TCP/IP for Windows.
#
# This file contains the mappings of IP addresses to host names. Each
# entry should be kept on an individual line. The IP address should
# be placed in the first column followed by the corresponding host name.
# The IP address and the host name should be separated by at least one
# space.
#
# Additionally, comments (such as these) may be inserted on individual
# lines or following the machine name denoted by a '#' symbol.
#
# For example:
#
#	102.54.94.97    rhino.acme.com      # source server
#	38.25.63.10     x.acme.com          # x client host

# localhost name resolution is handled within DNS itself.
#	127.0.0.1       localhost
#	::1             localhost
```
의 형식으로 기입. 
<hr>

+++ 추가로 서버가 다운되었는 지 아는 몇 가지의 테스트 코드.

> nslookup
- nslookup은 nameserver를 알아 볼 수 있는 키워드 윈도우에서 해당 명령어를 실행하게 되면 나의 도메인 서버를 알 수 있다.

> dig
- dig 나의서버 @8.8.8.8
- dig 를 이용해 @8.8.8.8(구글 도메인서버)에 나의서버를 리졸빙하는 요청.
<hr>

**성공일 때**

![image](https://github.com/soony1995/tips/assets/59558831/3f212738-e416-4dde-ae4a-7a5eb81dbe59)

>응답 헤더: status: NOERROR는 정상적인 응답임을 나타냅니다. 에러가 발생하지 않았으며, 쿼리를 성공적으로 처리한 것을 의미합니다.


**실패일 때**

![image](https://github.com/soony1995/tips/assets/59558831/aa98c61f-aad6-4ec5-8990-02aa55049c45)

>응답 헤더: status: NXDOMAIN은 도메인이 존재하지 않음을 나타냅니다. 즉, example.invalid은 실제로는 존재하지 않는 도메인임을 의미합니다.

