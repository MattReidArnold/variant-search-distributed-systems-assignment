package tsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const DateFormat = "1/2/06"
const Tab = '\t'

type RowCallback func(row []string) error
type FileCallback func(rowCount int, errs []error) error

func Read(filename string, rowCallback RowCallback, fileCallback FileCallback) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = Tab

	errs := []error{}

	headerRow, err := r.Read()
	if err == io.EOF {
		return fileCallback(0, errs)
	}
	if err != nil {
		return err
	}
	fmt.Printf("Processing file with headers: %s\n", headerRow)
	rowsRead := 1

	for i := 0; i < 20; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		rowsRead++
		err = rowCallback(row)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return fileCallback(rowsRead, errs)
}

func ParseTime(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ParseInt(s string) (*int64, error) {
	if s == "" || s == "NULL" {
		return nil, nil
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &i, nil

}
