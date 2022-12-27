package names

import (
	"log"
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// alphabet
var consonants = []string{"m", "n", "k", "g", "p", "t", "sh", "r", "l"}
var finals = []string{"n", "sh", "r", "l", "k", "i"}
var vowels = []string{"a", "e", "i", "u"}
var schemes = []string{"V", "VC", "CV", "CVC"}

type NameGenerator struct {
	consonants []string
	vowels     []string
	syllables  []string
	schemas    []string
}

func NewNameGenerator() *NameGenerator {
	g := &NameGenerator{
		consonants: consonants,
		vowels:     vowels,
		syllables:  []string{},
		schemas:    schemes,
	}
	g.generateSyllables()

	return g
}

func (g *NameGenerator) generateSyllables() {
	// g.syllables = append(g.syllables, g.vowels...)

	for _, C := range g.consonants {
		for _, V := range g.vowels {
			for _, schema := range g.schemas {
				switch schema {
				case "VC":
					for _, F := range finals {
						if V == F {
							continue
						}
						g.syllables = append(g.syllables, V+F)
					}
				case "CV":
					g.syllables = append(g.syllables, C+V)
				case "CVC":
					for _, F := range finals {
						if C == F || V == F {
							continue
						}
						g.syllables = append(g.syllables, C+V+F)
					}
				}
			}
		}
	}

	log.Println(g.syllables)
}

func (g *NameGenerator) GenerateName() string {
	rand.Shuffle(len(g.syllables), func(i, j int) {
		g.syllables[i], g.syllables[j] = g.syllables[j], g.syllables[i]
	})
	sylcount := 1 + rand.Intn(3)
	var name string = ""

	for i := 0; i < sylcount; i++ {
		randsyl := g.syllables[rand.Intn(len(g.syllables))]
		name += randsyl
		if sylcount == 3 && i == 0 {
			name += "-"
		}
	}

	caser := []cases.Caser{
		cases.Title(language.English),
	}
	name = caser[0].String(name)

	log.Println(name)

	return name
}
