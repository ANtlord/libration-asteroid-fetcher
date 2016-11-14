package main

import (
	"fmt"
	"libration-query-generator/datamining"
	"os"
	"log"
	"github.com/alexflint/go-arg"
	fp "path/filepath"
)

const OUT_FOLDER = "librated-asteroids"
func main() {
	var args struct {
		Filepath string `arg:"-f,required,help:Path to file stores integers and axises for pair of planets."`
		OnlyPure bool `arg:"-o,help:if 1 it will save numbers of asteroids that has only pure librations."`
		User string `arg:"-u,help:database user,help:by default is postgres."`
		Password string `arg:"-p,required,help:database user's password"`
		Database string `arg:"-d,help:database nane"`
		Host string `arg:"-s,required,help:database host"`
		Port string `arg:"-r,help:database port. By default is 5432"`
	}

	args.User = "postgres"
	args.Port = "5432"
	args.Database = "resonances"
	arg.MustParse(&args)

	var filepath = args.Filepath
	var ints = datamining.Build(filepath)
	const planet1 = "JUPITER"
	const planet2 = "SATURN"
	var miner = datamining.Miner{args.User, args.Password, args.Database, args.Host, args.Port}

	if _, err := os.Stat(OUT_FOLDER); os.IsNotExist(err) {
		var err = os.Mkdir(OUT_FOLDER, 0774)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, value := range ints {
		var asteroidNumbers = miner.FetchLibrations(value, planet1, planet2, args.OnlyPure)

		if len(asteroidNumbers) == 0 {
			continue
		}

		var outFilePath = fmt.Sprintf(
			"%s-%s_%d_%d_%d", planet1, planet2, value.First,
			value.Second, value.Asteroid,
		)
		var cwd, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		outFilePath = fp.Join(cwd, OUT_FOLDER, outFilePath)

		_, err = os.Stat(outFilePath)
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
