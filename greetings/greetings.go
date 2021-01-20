package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

/*
$ go mod init example.com/greetings
모듈을 경로로 지정. 프로덕션 코드에서 모듈을 다운로드 할 수 있는 URL이 됨.

Go에서 := 연산자는 한줄로 변수를 선언하고 초기화 하는 방법. 오른쪽에 있는 값이 변수의 유형을 결정한다.

만약 위 방법을 사용하지 않는다면 아래와 같이 작성할 수 있다.
func Hello(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
*/





























