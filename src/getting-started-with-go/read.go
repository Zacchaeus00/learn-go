package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Name struct {
		fname string
		lname string
	}
	fp := ""
	_, err := fmt.Scanf("%s", &fp)
	if err != nil {
		return
	}
	readFile, err := os.Open(fp)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	data := []Name{}
	for fileScanner.Scan() {
		name := strings.Fields(fileScanner.Text())
		fn, ln := name[0], name[1]
		data = append(data, Name{fn, ln})
	}
	for _, n := range data {
		fmt.Println(n)
	}

	err = readFile.Close()
	if err != nil {
		return
	}
}
