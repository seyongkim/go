package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	seed := time.Now().Unix()
	rand.Seed(seed)
	isPrime := true
	number := rand.Intn(150) + 2 //0과 1제외, 2~151까지의 숫자가 출력
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i <= number; i++ { //1과 num일때 반복문이 돌지 않는다
		if number%i == 0 {
			//count++
			isPrime = false
			//count = count + 1

		}
	}

	if isPrime == true {
		fmt.Println(number, "는 소수입니다!")
	} else {
		fmt.Println(number, "는 소수가 아닙니다!")

	}

}
