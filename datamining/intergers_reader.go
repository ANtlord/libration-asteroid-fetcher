// Package datamining provides getting intergers satisfying D'Alambert Rule from file
package datamining

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []string {
	var file, err = os.Open(path)
	if err != nil {
		fmt.Printf("Error during openning file %s\n", err.Error())
		os.Exit(-1)
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	var res = make([]string, 10)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

// Integers represents intergers satisfying D'Alambert Rule from file
type Integers struct {
	First, Second, Asteroid int
}

func newIntegers(line string) *Integers {
	var data = strings.Split(line, " ")
	const N = 3
	var vals = []int{0,0,0}
	for i := 0; i < N; i++ {
		val, err := strconv.Atoi(data[i])
		if err != nil {
			fmt.Println("Error during parsing " + data[i])
			os.Exit(-1)
		}

		vals[i] = val
	}
	var res = &Integers{ vals[0], vals[1], vals[2] }
	return res
}

// Build builds set of Integers
func Build(fromFilePath string) []*Integers {
	var fileData = readFile(fromFilePath)
	var res = make([]*Integers, 0)
	for _, line := range fileData {
		if len(line) > 0 {
			var integers = newIntegers(line)
			res = append(res, integers)
		}
	}
	return res
}
