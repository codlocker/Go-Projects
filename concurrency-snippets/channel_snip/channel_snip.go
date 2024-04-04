package channelsnip

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getSliceOfRandomNumbers(size int) []int {
	var numbers []int
	rand.Seed(int64(time.Now().UTC().Second()))
	min := 9999999
	max := 99999999

	for i := 0; i < size; i++ {
		random_number := rand.Intn(max-min+1) + min
		numbers = append(numbers, random_number)
	}

	return numbers
}

func TestChannel() {
	channel := make(chan string)
	go checkPrime(channel)

	for msg := range channel {
		fmt.Println(msg)
	}
}

func checkPrime(channel chan string) {
	fmt.Println("What is number of elements?")
	var count int
	fmt.Scan(&count)

	numbers := getSliceOfRandomNumbers(count)
	fmt.Println("Numbers are : ", numbers)

	for _, number := range numbers {
		var i int
		for i = 2; i < number; i++ {
			if number%i == 0 {
				channel <- strconv.Itoa(number) + " is not prime"
				break
			}
		}

		if i == number {
			channel <- strconv.Itoa(number) + " is prime"
		}
	}

	close(channel)
}
