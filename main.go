package main

import (
	"fmt"
	"libration-query-generator/datamining"
	"os"
	"log"
)

func main() {
	var filepath = os.Args[1]
	var ints = datamining.Build(filepath)
	const planet1 = "JUPITER"
	const planet2 = "SATURN"
	for _, value := range ints {
		var asteroidNumbers = datamining.FetchLibrations(value, planet1, planet2)

		if len(asteroidNumbers) == 0 {
			continue
		}

		var outFilePath = fmt.Sprintf(
			"%s-%s_%d_%d_%d", planet1, planet2, value.First,
			value.Second, value.Asteroid,
		)

		_, err := os.Stat(outFilePath)
		var outfile *os.File
		if os.IsNotExist(err) {
			outfile, err = os.Create(outFilePath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			outfile, err = os.OpenFile(outFilePath, os.O_WRONLY, 0)
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, v := range asteroidNumbers {
			outfile.WriteString(v+"\n")
		}
	}
}
