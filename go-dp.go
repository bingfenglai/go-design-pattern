package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./README.md")
	defer file.Close()

	if err != nil {
		log.Default().Println(err.Error())
	}

	input := bufio.NewReader(file)
	for {
		readString, err := input.ReadString('\n')
		fmt.Println(readString)
		if err == io.EOF {
			//fmt.Println(err)
			return
		}
	}
}
