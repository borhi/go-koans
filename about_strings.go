package go_koans

import "fmt"

func aboutStrings() {
	assert("a"+__string__ == "aimpossibly lame value") // string concatenation need not be difficult
	assert(len("abc") == 3)                            // and bounds are thoroughly checked

	assert(255 == __byte__) // their contents are reminiscent of C

	assert("impossibly lame value" == __string__) // slicing may omit the end point
	assert("impossibly lame value" == __string__) // or the beginning
	assert("impossibly lame value" == __string__) // or neither
	assert("impossibly lame value" == __string__) // or both

	assert("impossibly lame value" == __string__) // they can be compared directly
	assert("impossibly lame value" == __string__) // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	assert("impossibly lame value" == __string__) // strings can be created from byte-slices

	bytes[0] = 'z'
	assert("impossibly lame value" == __string__) // byte-slices can be mutated, although strings cannot

	assert(fmt.Sprintf("hello %s", __string__) == "hello impossibly lame value") // our old friend sprintf returns
	assert("impossibly lame value" == __string__)                                // quoting is familiar
	assert("impossibly lame value" == __string__)                                // although it can be done more easily

	assert("impossibly lame value" == __string__) // "the root of all evil" is actually a misquotation, by the way
}
