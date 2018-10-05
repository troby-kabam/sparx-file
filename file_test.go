package file

import (
	"testing"
	"reflect"
	"regexp"
	"encoding/json"
)

var test_file string		= "testfile.txt"
var test_file_data string	= "dGhpcyBpcyBhIHRlc3QK"
var test_file_checksum string	= "91751cee0a1ab8414400238a761411daa29643ab4b8243e9a91649e25be53ada"
var test_data_json []byte
var importedFile *FileData
var tErr error

func TestImportFile(t *testing.T) {
	testData, err := importFile(test_file)
	if err != nil {
		t.Errorf("importFile failed")
	}
	if testData.Name != test_file {
		t.Errorf("TestImportFile: functional test failed on testData.Name")
	}
	if testData.Checksum != test_file_checksum {
		t.Errorf("TestImportFile: functional test failed on testData.Checksum")
	}
	if testData.Data != test_file_data {
		t.Errorf("TestImportFile: functional test failed on testData.Data")
	}
}

func TestNewFileData(t *testing.T) {
	importedFile, tErr = NewFileData(test_file)
	if tErr != nil {
		t.Errorf("TestNewFileData failed")
	}
}

func TestGetName(t *testing.T) {
	if importedFile.GetName() != test_file {
		t.Errorf("GetName method failed")
	}
}

func TestGetData(t *testing.T) {
	if importedFile.GetData() != test_file_data {
		t.Errorf("GetData method failed")
	}
}

func TestGetChecksum(t *testing.T) {
	if importedFile.GetChecksum() != test_file_checksum {
		t.Errorf("GetChecksum method failed")
	}
}

// return struct as json formatted []byte data
func TestMarshalJSON(t *testing.T) {
	json_data, err := importedFile.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON generated error")
	}
	if json.Valid(json_data) != true {
		t.Errorf("TestMarshalJSON: invalid json data")
	}
	test_data_json = json_data
}

func TestGetBuffer(t *testing.T) {
	bExpected := []byte{}
	b, err := getBuffer(test_file)
	if err != nil {
		t.Errorf("getBuffer failed")
	}
	if reflect.TypeOf(b) != reflect.TypeOf(bExpected) {
		t.Errorf("type mismatch")
	}
}

func TestMakeChecksum(t *testing.T) {
	b := []byte{}
	expected := "^[a-f0-9]{64}$"
	sum := makeChecksum(b)
	match, err := regexp.MatchString(expected, sum)
	if err != nil {
		t.Errorf("regexp.Match error")
	}
	if match != true {
		t.Errorf("checksum mismatch")
	}
}
