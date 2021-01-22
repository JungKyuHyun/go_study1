package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// 로그 속성 설정
	log.SetPrefix("greetings.Hello :")
	log.SetFlags(0);

	names := []string{"jacob", "cob", "cobe"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
} 

/*
- main 패키지 선언. 
- Go 애플리케이션으로 실행되는 코드는 main 패키지에 있어야 한다.
- hello 패키지에 대한 새 모듈을 만듬


$ go build
- Go가 모듈을 찾고 go.mod 파일에 종속성으로 추가

window 
$ .\hello.exe

mac or linux
$ ./hello
*/