package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// set ititial value
// go init은 전역 변수가 초기화된 후 프로그램 시작시 자동으로 함수를 실행
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// 괄호안에 크기를 생략하여 []string 슬라이스의 기본 배열이 동적으로 크기 조정될 수 있게 함.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
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


// errors 모듈 추가
모든 GO 함수는 여러 값을 반환할 수 있다. 
message, nil -> 성공적인 반환에서 두번째 값으로 (오류 없음을 의미) 추가
호출자가 함수가 성공했음을 알 수 있다.

*/





























