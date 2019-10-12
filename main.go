package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func main() {
	File2File("test.csv", "test.json", true)
}

// File2File pulls data from an RFC 4180 compliant csv file and uses it to create a json file.
// If the csv file contains an optional header line, set hasHeaderLine to true to use those header values to create a more descriptive json.
func File2File(csvFilename, jsonFilename string, hasHeaderLine bool) error {
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		return errors.Wrapf(err, "could not open specified csv file, filename: %s", csvFilename)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	if hasHeaderLine {
		headers, err := r.Read()
		if err != nil {
			return errors.Wrap(err, "encountered an error attempting to read the header line")
		}

		elements := make(map[string]map[string]string)
		counter := 0
		for line, err := r.Read(); err != io.EOF; line, err = r.Read() {
			innerMap := make(map[string]string)
			for i := 0; i < len(line); i++ {
				innerMap[headers[i]] = line[i]
			}
			elements[fmt.Sprintf("record_%d", counter)] = innerMap
			counter++
		}

		file, err := json.MarshalIndent(elements, "", " ")
		if err != nil {
			return errors.Wrap(err, "json.Marhsal encountered an error attempting to Marshal(elements) into []byte")
		}

		err = ioutil.WriteFile(jsonFilename, file, 0644)
		if err != nil {
			return errors.Wrapf(err, "io.util encountered an error writing the marshalled json data (a []byte) to the filename %s", jsonFilename)
		}
	}

	return nil
}
