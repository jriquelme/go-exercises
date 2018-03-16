package dictionaryquiz

func ExamplePrintNotFoundWords() {
	err := PrintNotFoundWords("testdata/dictionary.txt", "testdata/text.txt")
	if err != nil {
		panic(err)
	}
	// Output:
	// are
	// Would
	// be
	// Hola
	// Cómo
	// estás
}
