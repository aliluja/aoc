package main

import (
	"io/ioutil"
	"strings"
	"unicode"
	"fmt"
)

//parse the input file
func parse() ([]string, error) {
	input, err := ioutil.ReadFile("../inputs/way")
	if err != nil {
		fmt.Println("Beim Auslesen von der Text-Datei ist ein Fehler aufgetreten", err)
		return nil, err
	}
	data := string(input)

	f := func(r rune) bool {
		return unicode.IsSpace(r) && r == '\n'
	}

	splitted := strings.FieldsFunc(data, f)

	return splitted, nil
}
