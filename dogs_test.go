package dogs

import "testing"

func TestIsFemale(t *testing.T) {
	if !IsFemale("Bella") {
		t.Fatal("Should be female")
	}
	if IsFemale("Sam") {
		t.Fatal("Should be male")
	}
}

func TestRandomMale(t *testing.T) {
	for i := 0; i < 10; i++ {
		name := RandomMale()
		if !IsMale(name) {
			t.Fatal("Wrong random male", name)
		}
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		name, gender := Random()
		if gender == Male && !IsMale(name) {
			t.Fatal("Wrong random", name, gender)
		} else if gender == Female && !IsFemale(name) {
			t.Fatal("Wrong random", name, gender)
		}
	}
}
