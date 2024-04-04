package main

import (
	channelsnip "concurrency-snippets/channel_snip"
	"concurrency-snippets/wait_group"
	"fmt"
)

func main() {
	fmt.Println("Test\n1. Channel\n2. WaitGroup")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		channelsnip.TestChannel()
	case 2:
		wait_group.TestWaitGroup()
	default:
		fmt.Println("Incorrect choice")
	}
}
