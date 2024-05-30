package parser

import (
	"bufio"
	"bytes"
	"io"
)

const (
	START = "CREATE TABLE"
	END   = ";"
)

func ParseCreateTable(file io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(file)
	var parsedArr [][]byte
	parsing := false

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.HasPrefix(line, []byte(START)) {
			parsing = true
		}
		if !parsing {
			continue
		}
		parsedArr = append(parsedArr, line)
		if bytes.HasSuffix(line, []byte(";")) {
			parsing = false
		}
	}

	return bytes.Join(parsedArr, []byte("\n")), nil
}
