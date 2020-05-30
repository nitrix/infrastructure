package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	output := bytes.Buffer{}

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)
	i := 0

	for s.Scan() {
		word := s.Text()
		i++

		if strings.HasPrefix(word, "-----") {
			if i != 1 {
				output.WriteString("\n")
			}

			output.WriteString(word + " ")
			continue
		}

		if strings.HasSuffix(word, "-----") {
			output.WriteString(word + " \n")
			continue
		}

		output.WriteString(word)
	}

	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Truncate(os.Args[1], 0)
	if err != nil {
		log.Fatalln(err)
	}

	file, err = os.OpenFile(os.Args[1], os.O_RDWR, 0600)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.WriteString(output.String())
	if err != nil {
		log.Fatalln(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
}