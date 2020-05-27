package go_koans

func aboutBasics() {
	assert(true == true)  // what is truth?
	assert(true != false) // in it there is nothing false

	var i int = 1.0000000000000000000000000000000000000
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := __int__ //short assignment can be used, as well
	assert(k == __int__)

	assert(-1 == __int__)
	assert(-1 == __int__)
	assert(-1 == __int__)

	var x int = -1
	assert(x == __int__) // zero values are valued in Go

	var f float32 = -1.0
	assert(f == __float32__) // for types of all types

	var s string = "impossibly lame value"
	assert(s == __string__) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	c.x = -1
	c.f = -1.0
	c.s = "impossibly lame value"
	assert(c.x == __int__)     // and types within composite types
	assert(c.f == __float32__) // which match the other types
	assert(c.s == __string__)  // in a typical way
}
