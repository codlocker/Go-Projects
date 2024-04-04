package wait_group

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func getSliceOfRandomNumbers(size int) []int {
	var numbers []int
	rand.Seed(int64(time.Now().UTC().Second()))
	min := 999999999
	max := 9999999999

	for i := 0; i < size; i++ {
		random_number := rand.Intn(max-min+1) + min
		numbers = append(numbers, random_number)
	}

	return numbers
}

func TestWaitGroup() {
	var waitGroup sync.WaitGroup
	fmt.Println("How many numbers to check?")
	var count int
	//count = 5
	fmt.Scan(&count)
	numbers := getSliceOfRandomNumbers(count)

	fmt.Printf("Numbers are %v\n", numbers)
	waitGroup.Add(count)

	go func() {
		for _, value := range numbers {
			fmt.Println(checkPrime(value))
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}

func checkPrime(number int) string {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return strconv.Itoa(number) + " is not a prime"
		}
	}

	return strconv.Itoa(number) + " is a prime"
}
