package faker

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// gerador de nome
func Name() string {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	return firstName + " " + lastName
}

// gerador de endereco
func Address() string {
	street := streets[rand.Intn(len(streets))]
	number := rand.Intn(1000) + 1

	return fmt.Sprintf("%d %s", number, street)
}
