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
