package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

//parse the input file
func parse() (nums string) {
	data, err := ioutil.ReadFile("../inputs/passwords")
	if err != nil {
		fmt.Println("Beim Auslesen von der Text-Datei ist ein Fehler aufgetreten", err)
		return
	}
	return string(data)
}

func linesToSlice() ([]string, error) {
	data := parse()
	cleaned :=  strings.ReplaceAll(data, "\n", ",")
	splitted := strings.Fields(cleaned)
	fmt.Println(len(splitted))
	return splitted, nil
}

func main() {
	a , _ := linesToSlice()
	fmt.Println(len(a))

}