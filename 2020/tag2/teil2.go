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

type key struct {
	ltr, psw string
	first, second int
}

//parse the input file
func parseData() ([]string, error) {
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

func operate() int {
	keys := []*key{}
	a, _ := parse()
	prettyslice.Show("passwords: ", a)
	input := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)`)

	for _, i := range a {
		first, _ := strconv.Atoi(input.ReplaceAllString(i, "$1"))
		second, _ := strconv.Atoi(input.ReplaceAllString(i, "$2"))
		keys = append(keys, &key{
			ltr: input.ReplaceAllString(i, "$3"),
			psw: input.ReplaceAllString(i, "$4"),
			first: first,
			second: second,
		})
	}

	valid := 0

	for _, k := range keys{
		if (string(k.psw[k.first-1]) == k.ltr) != (string(k.psw[k.second-1]) == k.ltr) {
			valid++
		}
	}

	return valid
}

func main() {
	fmt.Println(operate())
}