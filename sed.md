sed를 이용하면 파일을 읽어온 다음 필요한 부분만 정규표현식으로 찾고 바로 변경할 수 있다.   
예시) 
```
cmd := fmt.Sprintf("sed -i 's/file\\s*=\\s*.*/file= %s'/ %s", "test", "./default.jcfg")
err := execShellByRun(cmd)
fmt.Println(err)

func execShellByRun(cmd string) error {
	var stderr bytes.Buffer
	execCmd := exec.Command("/bin/sh", "-vc", cmd)
	execCmd.Stderr = &stderr
	if err := execCmd.Run(); err != nil {
		return errors.New(stderr.String())
	}
	return nil
}

```
## sed는 공백문자 \s를 받을 수 없다. 
```
cmd += fmt.Sprintf("-e 's/^\\([[:space:]]*[^#]*%s[[:space:]]*=[[:space:]]*\\).*/\\1%v/' ", jsonTag, value)
와 같이 지저분하게 정규 표현식을 써야 한다.
```
> 대안으로는 perl , awk를 사용하는 방법이 있다.
```
cmd += fmt.Sprintf("-e 's/^\s*([^#]*%s\s*=\s*).*/$1%v/", jsonTag, value)
로 변경할 수 있다. 
```
