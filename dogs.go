// Package dogs provides the most popular dog names
package dogs

import (
	"math/rand"
)

// Females - Top 100 female dog names sorted by popularity
var Females = []string{
	"Bella",
	"Lucy",
	"Daisy",
	"Molly",
	"Lola",
	"Sophie",
	"Sadie",
	"Maggie",
	"Chloe",
	"Bailey",
	"Roxy",
	"Zoey",
	"Lily",
	"Luna",
	"Coco",
	"Stella",
	"Gracie",
	"Abby",
	"Penny",
	"Zoe",
	"Ginger",
	"Ruby",
	"Rosie",
	"Lilly",
	"Ellie",
	"Mia",
	"Sasha",
	"Lulu",
	"Pepper",
	"Nala",
	"Lexi",
	"Lady",
	"Emma",
	"Riley",
	"Dixie",
	"Annie",
	"Maddie",
	"Piper",
	"Princess",
	"Izzy",
	"Maya",
	"Olive",
	"Cookie",
	"Roxie",
	"Angel",
	"Belle",
	"Layla",
	"Missy",
	"Cali",
	"Honey",
	"Millie",
	"Harley",
	"Marley",
	"Holly",
	"Kona",
	"Shelby",
	"Jasmine",
	"Ella",
	"Charlie",
	"Minnie",
	"Willow",
	"Phoebe",
	"Callie",
	"Scout",
	"Katie",
	"Dakota",
	"Sugar",
	"Sandy",
	"Josie",
	"Macy",
	"Trixie",
	"Winnie",
	"Peanut",
	"Mimi",
	"Hazel",
	"Mocha",
	"Cleo",
	"Hannah",
	"Athena",
	"Lacey",
	"Sassy",
	"Lucky",
	"Bonnie",
	"Allie",
	"Brandy",
	"Sydney",
	"Casey",
	"Gigi",
	"Baby",
	"Madison",
	"Heidi",
	"Sally",
	"Shadow",
	"Cocoa",
	"Pebbles",
	"Misty",
	"Nikki",
	"Lexie",
	"Grace",
	"Sierra",
}

// Males - Top 100 male dog names sorted by popularity
var Males = []string{
	"Max",
	"Buddy",
	"Charlie",
	"Jack",
	"Cooper",
	"Rocky",
	"Toby",
	"Tucker",
	"Jake",
	"Bear",
	"Duke",
	"Teddy",
	"Oliver",
	"Riley",
	"Bailey",
	"Bentley",
	"Milo",
	"Buster",
	"Cody",
	"Dexter",
	"Winston",
	"Murphy",
	"Leo",
	"Lucky",
	"Oscar",
	"Louie",
	"Zeus",
	"Henry",
	"Sam",
	"Harley",
	"Baxter",
	"Gus",
	"Sammy",
	"Jackson",
	"Bruno",
	"Diesel",
	"Jax",
	"Gizmo",
	"Bandit",
	"Rusty",
	"Marley",
	"Jasper",
	"Brody",
	"Roscoe",
	"Hank",
	"Otis",
	"Bo",
	"Joey",
	"Beau",
	"Ollie",
	"Tank",
	"Shadow",
	"Peanut",
	"Hunter",
	"Scout",
	"Blue",
	"Rocco",
	"Simba",
	"Tyson",
	"Ziggy",
	"Boomer",
	"Romeo",
	"Apollo",
	"Ace",
	"Luke",
	"Rex",
	"Finn",
	"Chance",
	"Rudy",
	"Loki",
	"Moose",
	"George",
	"Samson",
	"Coco",
	"Benny",
	"Thor",
	"Rufus",
	"Prince",
	"Chester",
	"Brutus",
	"Scooter",
	"Chico",
	"Spike",
	"Gunner",
	"Sparky",
	"Mickey",
	"Kobe",
	"Chase",
	"Oreo",
	"Frankie",
	"Mac",
	"Benji",
	"Bubba",
	"Champ",
	"Brady",
	"Elvis",
	"Copper",
	"Cash",
	"Archie",
	"Walter",
}

// All - Top 200 (fe)male dog names
// Same ranked female comes before male correspondent
var All []string

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(123457))
	All = make([]string, 200, 200)
	for i := 0; i < 100; i++ {
		All[i*2] = Females[i]
		All[i*2+1] = Males[i]
	}
}

func contains(name string, names []string) bool {
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}

// Gender - Female or Male
type Gender bool

// Gender Type returned from random
const (
	Female Gender = true
	Male   Gender = false
)

// IsFemale checks if given name is a female dog name
// Runs in O(n)
func IsFemale(name string) bool {
	return contains(name, Females)
}

// IsMale checks if given name is a male dog name
// Runs in O(n)
func IsMale(name string) bool {
	return contains(name, Males)
}

// Is checks if given name is a dog name
// Runs in O(n)
func Is(name string) bool {
	return contains(name, All)
}

// RandomFemale returns a random Male dog name
func RandomFemale() string {
	return Females[r.Int31n(100)]
}

// RandomMale returns a random male dog name
func RandomMale() string {
	return Males[r.Int31n(100)]
}

// Random returns a random dog name
func Random() (string, Gender) {
	n := r.Int31n(200)
	return All[n], n%2 == 0
}
