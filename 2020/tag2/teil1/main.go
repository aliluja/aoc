package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"fmt"

	"github.com/inancgumus/prettyslice"
)

type password struct {
	ltr, psw string
	min, max int
}

//parse the input file
func parse() ([]string, error) {
	input, err := ioutil.ReadFile("../inputs/passwords")
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

func convert() int {
	passwords := []*password{}
	a, _  := parse()
	prettyslice.Show("passwords: ", a)
	input := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)`)

	for _ , i := range a {
		min, _ := strconv.Atoi(input.ReplaceAllString(i, "$1"))
		max, _ := strconv.Atoi(input.ReplaceAllString(i, "$2"))
		passwords = append(passwords, &password{
			ltr: input.ReplaceAllString(i, "$3"),
			psw: input.ReplaceAllString(i, "$4"),
			min: min,
			max: max,
		})

	}

	match := 0
	for _, pw := range passwords{
		if i := strings.Count(pw.psw, pw.ltr); i >= pw.min && i <= pw.max {
			match++
		}
	}
	
	return match
}

func main() {
	fmt.Println(convert())
}