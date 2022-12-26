package names

import (
	"math/rand"
)

var initial = []string{"Ur", "La", "Ak", "Nip", "Ak", "Wa", "Ka", "Ba", "Kab", "Li"}
var final = []string{"", "uk", "shak", "pur", "kad", "du", "lum"}

func GenerateName() string {
	rand.Shuffle(len(initial), func(i, j int) {
		initial[i], initial[j] = initial[j], initial[i]
	})
	rand.Shuffle(len(final), func(i, j int) {
		final[i], final[j] = final[j], final[i]
	})
	return initial[rand.Intn(len(initial))] + final[rand.Intn(len(final))]
}
