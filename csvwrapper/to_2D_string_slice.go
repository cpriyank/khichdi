package csvwrapper

// package csvwrapper is a wrapper around go's encoding/csv package to read and
// write csv files

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
)

// read a csv file specified by filepath as a slice of strings
// e.g. If the file contains:
// first_name,last_name,username
// "Rob","Pike",rob
// Ken,Thompson,ken
// "Robert","Griesemer","gri"
//
// it returns:
// [first_name last_name username]
// [Rob Pike rob]
// [Ken Thompson ken]
// [Robert Griesemer gri]

func CSVTo2DStringSlice(fileName string) ([][]string, error) {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(in)
	r := csv.NewReader(buf)
	// encoding/csv allows specifying alternatives if fields are not separated by
	// comma, like so:
	// r.Comma = ';'
	// r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
