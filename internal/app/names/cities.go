package names

import (
	"math/rand"
)

var names = []string{"Ur", "Uruk", "lagash", "nippur", "akkad", "Wakad", "Kad", "Kadu", "Lilum"}

func GenerateName() string {
	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})
	return names[rand.Intn(len(names))]
}
