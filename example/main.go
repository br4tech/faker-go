package main

import (
	"fmt"

	"github.com/br4tech/video-aulas-go/faker-go/faker"
)

func main() {
	fmt.Println("Random Name:", faker.Name())
	fmt.Println("Random Name:", faker.Address())
}
