package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	e, err := ioutil.ReadFile("bonjour.txt")
	check(err)
	fmt.Print(string(e))

}
