package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func main() {
	for i := 0; i < 10000000; i++ {
		fmt.Println(string(fake.IPv4()))
	}
}
