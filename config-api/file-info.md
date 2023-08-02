# 파일에 대한 메타데이터를 얻어오는 방식

```
func TimeParser(filePath string, cmd string) (string, error){
    cmd += fmt.Sprintf("stat -c %%y %s", filePath) // %y는 형식지정자로 쓰이는 게 아니기 때문에 escape 처리를 해줘야 한다. %의 경우 \가 아닌 %을 써서 escape를 시켜준다.

	modTime, _ := w.ExecShellByOutput(cmd)

	if len(modTime) == 0 {
		return "", fmt.Errorf("TimeParser | modTime is empty value")
	}

	return modTime, nil
}

// shell을 이용해 err를 받게 되면 자세하게 나오지 않아, 상세한 에러가 나오게 구현했다.
func ExecShellByOutput(cmd string) (string, error) {
	var stderr bytes.Buffer
	execCmd := exec.Command("/bin/sh", "-vc", cmd)
	execCmd.Stderr = &stderr
	output, _ := execCmd.Output()
	if len(output) == 0 {
		return "", errors.New(stderr.String())
	}
	return string(output), nil
}
```
### 이런식으로 stat을 이용해 해당 파일의 메타 데이터를 얻어 올 수 있다.

명령어를 입력한 결과이다.
> stat server.sh
```
File: server.sh
  Size: 1779            Blocks: 8          IO Block: 4096   regular file
Device: ca11h/51729d    Inode: 131444      Links: 1
Access: (0775/-rwxrwxr-x)  Uid: ( 1000/  nimbus)   Gid: ( 1000/  nimbus)
Access: 2023-05-19 02:20:58.006262185 +0000
Modify: 2023-03-08 02:12:53.414514201 +0000
Change: 2023-03-08 02:12:53.414514201 +0000
 Birth: -
```
- Blocks : 하드디스크에 저장하기 위해 할당된 블럭 수 
- IO Block : 할당된 블럭 크기 (바이트)
- Inode : 파일 Inode 값
- Links : 파일 링크 수
- Access : 접근 권한
- Access : 최종 파일 접근 시간
- Modify : 최종 파일 수정 시간
- Change : 속성 또는 컨텐츠 최종 변경 시간
- Birth : 파일 생성 시간 (파일 시스템에 따라서 -로 나옴)

### 내가 썼던 flag들 
> stat -f server.sh 

파일에 대한 정보를 가져온다. 만약에 가져오지 못한다면 파일이 없는 것으로 간주하고 에러처리.
```
File: "server.sh"
ID: 17162432fb97d488 Namelen: 255     Type: ext2/ext3
Block size: 4096       Fundamental block size: 4096
Blocks: Total: 12835704   Free: 9694722    Available: 9035279
Inodes: Total: 3276800    Free: 2856238
```

> stat -c %y server.sh
```
2023-08-02 01:41:28.466702340 +0000
가장 최근에 수정한 날짜 
```
### 더 많은 flag를 알고 싶다면 
> stat --help