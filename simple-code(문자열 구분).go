exampleCode = "hi i'm soony1995, age is 28"

// 해당 상황은 "," 하나만 등장했을 때를 가정했다. 
// exampleCode 문자열 중에 ","을 구분으로 앞의 부분과 뒷 부분을 파싱을 하려고 한다.

> 나였다면 split을 이용해서 나눴을 테지만 아래의 코드로도 구현이 가능했다. 

func main() { 
	// index 함수를 이용한 경우
	if strings.Contains(exampleCode, ",") {
		index := strings.Index(exampleCode, ",")
		if index != -1 {
			jsonTag = jsonTag[:index]
		}
	}

	// split을 이용한 경우 
	exampleCode := "hi i'm soony1995, age is 28"
	// Split the string into a slice on comma
	if strings.Contains(exampleCode, ",") {
		parsed := strings.Split(exampleCode, ",")
	}	
}

// index의 경우 "," 문자를 찾게 되면 중단 하는 반면에
// split 함수는 항상 모든 "," 을 찾아 나눠야 하기 때문에 조금 느릴 수 있다.
// 하지만 차이는 근소하다고 한다. 