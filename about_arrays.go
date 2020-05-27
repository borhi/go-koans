package go_koans

func aboutArrays() {
	fruits := [4]string{"apple", "orange", "mango"}

	assert("impossibly lame value" == __string__) // indexes begin at 0
	assert("impossibly lame value" == __string__) // one is indeed the loneliest number
	assert("impossibly lame value" == __string__) // it takes two to ...tango?
	assert("impossibly lame value" == __string__) // there is no spoon, only an empty value

	assert(-1 == __int__) // the length is what the length is
	assert(-1 == __int__) // it can hold no more

	assert(fruits == fruits) // comparing arrays is not like comparing apples and oranges

	tasty_fruits := fruits[1:3]                   // defining oneself as a variation of another
	assert("impossibly lame value" == __string__) //and get not a simple array as a result
	assert("impossibly lame value" == __string__) // slices of arrays share some data
	assert("impossibly lame value" == __string__) // albeit slightly askewed

	assert(-1 == __int__) // its length is manifest
	assert(-1 == __int__) // but its capacity is surprising!

	tasty_fruits[0] = "lemon" // are their shared roots truly identical?

	assert("impossibly lame value" == __string__) // has this element remained the same?
	assert("impossibly lame value" == __string__) // how about the second?
	assert("impossibly lame value" == __string__) // surely one of these must have changed
	assert("impossibly lame value" == __string__) // but who can know these things

	veggies := [...]string{"carrot", "pea"}

	assert(len(veggies) == 2) // array literals need not repeat an obvious length
}
