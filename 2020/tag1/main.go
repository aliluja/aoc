package main

import (
	"io/ioutil"
	"strconv"
	"fmt"
	"strings"
)

//parse the input file
func parse() (nums string) {
	data, err := ioutil.ReadFile("../inputs/nums")
	if err != nil {
		fmt.Println("Beim Auslesen von der Text-Datei ist ein Fehler aufgetreten", err)
		return
	}
	return string(data)
}

func stringToSlice() ([]int, error) {
	data := parse()
	cleaned :=  strings.ReplaceAll(data, "\n", " ")
	splitted := strings.Fields(cleaned)

	var slice = []int{}

	for _, i := range splitted {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		slice = append(slice, j)
	}

	return slice, nil
}

func sortNum() (a,b,c int) {
	rowOutput, _ := stringToSlice()
	fmt.Println("objects in slice: ",len(rowOutput))
	for _, a := range rowOutput {
		for _, b := range rowOutput {
			for _, c := range rowOutput{
				if a + b + c  == 2020 {
					return a, b, c
				}
			}
		}
	}
	return a,b,c
}

func main() {
	a,b,c := sortNum()
	fmt.Println("Res: ", a*b*c)
}
