package dictionaryquiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PrintNotFoundWords prints words from textFilePath not found in dictionaryPath. Each line in dictionaryPath
// is expected to contains one word definition, using the format:
//
// word1:definition of word1
// word2:definition of word2
// ...
// wordN:definition of wordN
//
// Words are case insensitive; also, any duplicated word in the dictionary will replace previous definitions.
func PrintNotFoundWords(dictionaryPath, textFilePath string) error {
	// read dictionary
	dictionaryFile, err := os.Open(dictionaryPath)
	if err != nil {
		return err
	}
	defer dictionaryFile.Close()
	dictionary := make(map[string]string)
	// scan lines
	dictionaryScanner := bufio.NewScanner(dictionaryFile)
	for dictionaryScanner.Scan() {
		entry := dictionaryScanner.Text()
		separator := strings.Index(entry, ":")
		if separator == -1 {
			return fmt.Errorf("invalid entry, separator not found: %s", entry)
		}
		word := entry[:separator]
		word = strings.ToLower(word)
		definition := entry[separator+1:]
		dictionary[word] = definition
	}
	if dictionaryScanner.Err() != nil {
		return dictionaryScanner.Err()
	}

	// process text file
	textFile, err := os.Open(textFilePath)
	if err != nil {
		return err
	}
	defer textFile.Close()
	// punctuation marks are removed
	pmarksReplacer := strings.NewReplacer(",", "", ".", "", "¿", "", "?", "", "¡", "", "!", "", ";", "")
	textFileScanner := bufio.NewScanner(textFile)
	textFileScanner.Split(bufio.ScanWords) // scan by words
	for textFileScanner.Scan() {
		word := pmarksReplacer.Replace(textFileScanner.Text())
		lcWord := strings.ToLower(word)
		_, found := dictionary[lcWord]
		if !found {
			fmt.Printf("%s\n", word)
		}
	}
	if textFileScanner.Err() != nil {
		return textFileScanner.Err()
	}

	return nil
}
