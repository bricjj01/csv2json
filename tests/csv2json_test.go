package csv2json_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/bricjj01/csv2json"
)

func TestFile2FileWithHeader(t *testing.T) {
	err := csv2json.File2File("test.csv", "TestFile2FileWithHeader.json", true)
	if err != nil {
		t.Errorf("File2File error: %v", err)
	}

	test, err := ioutil.ReadFile("TestFile2FileWithHeader.json")
	if err != nil {
		t.Errorf("ioutil couldn't read contents of TestFile2FileWithHeader.json. error - %v", err)
	}
	verified, err := ioutil.ReadFile("VerifiedWithHeader.json")
	if err != nil {
		t.Errorf("ioutil couldn't read contents of VerifiedWithHeader.json. error - %v", err)
	}

	if bytes.Compare(test, verified) != 0 {
		t.Errorf("TestFile2FileWithHeader failed. bytes.Compare(test,verified) says the []byte of the files are different.")
	}
}

func TestFile2FileNoHeader(t *testing.T) {
	err := csv2json.File2File("test.csv", "TestFile2FileNoHeader.json", false)
	if err != nil {
		t.Errorf("File2File error: %v", err)
	}

	test, err := ioutil.ReadFile("TestFile2FileNoHeader.json")
	if err != nil {
		t.Errorf("ioutil couldn't read contents of TestFile2FileNoHeader.json. error - %v", err)
	}
	verified, err := ioutil.ReadFile("VerifiedNoHeader.json")
	if err != nil {
		t.Errorf("ioutil couldn't read contents of VerifiedNoHeader.json. error - %v", err)
	}

	if bytes.Compare(test, verified) != 0 {
		t.Errorf("TestFile2FileNoHeader failed. bytes.Compare(test,verified) says the []byte of the files are different.")
	}
}

func TestFile2InMemoryHeader(t *testing.T) {
	test, err := csv2json.File2InMemory("test.csv", true)
	if err != nil {
		t.Errorf("File2InMemory error: %v", err)
	}

	verified, err := ioutil.ReadFile("VerifiedWithHeader.json")
	if err != nil {
		t.Errorf("ioutil.Readfile(VerifiedWithHeader.json) error: %v", err)
	}

	if bytes.Compare(test, verified) != 0 {
		t.Errorf("TestFile2InMemoryHeader failed. bytes.Compare(test, verified) says the []byte of the files are different.")
	}
}

func TestFile2InMemoryNoHeader(t *testing.T) {
	test, err := csv2json.File2InMemory("test.csv", false)
	if err != nil {
		t.Errorf("File2InMemory error: %v", err)
	}

	verified, err := ioutil.ReadFile("VerifiedNoHeader.json")
	if err != nil {
		t.Errorf("ioutil.Readfile(VerifiedNoHeader.json) error: %v", err)
	}

	if bytes.Compare(test, verified) != 0 {
		t.Errorf("TestFile2InMemoryNoHeader failed. bytes.Compare(test, verified) says the []byte of the files are different.")
	}
}
