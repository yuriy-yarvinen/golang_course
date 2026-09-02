package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"net/http"
	"os"
)

type Data struct {
	Header   string
	FirstRow []string
}

func ParseData(source io.Reader) (*Data, error) {
	bufReader := bufio.NewReader(source)
	header, err := bufReader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(bufReader)
	row, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	return &Data{
		Header:   header,
		FirstRow: row,
	}, nil
}

func ProcessFile() {
	file, _ := os.Open("some.dat")
	defer file.Close()
	ParseData(file)
}

func ProcessHttp() {
	res, _ := http.Get("http://yarvinen.ru/file.dat")
	defer res.Body.Close()
	ParseData(res.Body)
}
