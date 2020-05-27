package go_koans

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == fruits[0]) // slices seem like arrays
	assert(len(fruits) == 3)       // in nearly all respects

	tasty_fruits := fruits[1:3]                // we can even slice slices
	assert(tasty_fruits[0] == tasty_fruits[0]) // slices of slices also share the underlying data

	pregnancy_slots := []string{"baby", "baby", "lemon"}
	assert(cap(pregnancy_slots) == cap(pregnancy_slots)) // the capacity is initially the length

	pregnancy_slots = append(pregnancy_slots, "baby!")
	assert(len(pregnancy_slots) == len(pregnancy_slots)) // slices can be extended with append(), much like realloc in C
	assert(cap(pregnancy_slots) == cap(pregnancy_slots)) // but with better optimizations

	pregnancy_slots = append(pregnancy_slots, "another baby!?", "yet another, oh dear!", "they must be Catholic")

	assert(len(pregnancy_slots) == len(pregnancy_slots)) // append() can take N arguments to append to the slice
	assert(cap(pregnancy_slots) == cap(pregnancy_slots)) // the capacity optimizations have a guessable algorithm
}
