package go_koans

func aboutCommonInterfaces() {
	{
		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert("hello world" == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		assert("hello" == "hello") // duplicate only a portion of the io.Reader
	}
}
