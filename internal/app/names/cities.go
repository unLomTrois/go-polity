package names

import (
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// alphabet
var consonants = []string{"m", "n", "k", "p", "t", "sh", "r", "l", "v"}
var vowels = []string{"a", "e", "i", "u", "o"}
var schemes = []string{"V", "VC", "CV"}

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

	for _, V := range g.vowels {
		for _, C := range g.consonants {
			for _, schema := range g.schemas {
				if schema == "VC" {
					g.syllables = append(g.syllables, V+C)
				}
				if schema == "CV" {
					g.syllables = append(g.syllables, C+V)
				}
				if schema == "CVC" {
					for _, C2 := range consonants {
						g.syllables = append(g.syllables, C+V+C2)
					}
				}
			}
		}
	}

	// log.Println(g.syllables)
}

func (g *NameGenerator) GenerateName() string {
	// rand.Shuffle(len(g.syllables), func(i, j int) {
	// 	g.syllables[i], g.syllables[j] = g.syllables[j], g.syllables[i]
	// })
	syllablescount := 1 + rand.Intn(3)
	var name string = ""
	// log.Print("count ", syllablescount)

	for i := 0; i < syllablescount; i++ {
		randsyl := g.syllables[rand.Intn(len(g.syllables))]
		name += randsyl
	}

	// name = cases.(name)
	caser := []cases.Caser{
		cases.Title(language.English),
	}
	name = caser[0].String(name)

	// log.Println(name)

	return name
}
