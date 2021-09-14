//소수 판정 프로그램 version 0.1
//소수는 1과 자기자신외에는 나누어 떨어지지 않는 수(0과 1은 제외)
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	seed := time.Now().Unix()
	rand.Seed(seed)
	count := 0
	number := rand.Intn(150) + 2 //0과 1제외, 2~151까지의 숫자가 출력
	fmt.Println("임의로 추출된 수 : ", number)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 2 {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {

		fmt.Println(number, "는(은) 소수가 아닙니다!")
	}
}
