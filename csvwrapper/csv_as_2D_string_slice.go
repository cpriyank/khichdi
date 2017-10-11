package csvwrapper

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
)

func CSVTo2DStringSlice(file string) ([][]string, error) {
	in, err1 := ioutil.ReadFile(file)
	buf := bytes.NewBuffer(in)
	r := csv.NewReader(buf)
	// r.Comma = ';'
	// r.Comment = '#'

	records, err2 := r.ReadAll()
	return records, fmt.Errorf("%v : %v", err1, err2)
}
