package file

import (
	"testing"
	"encoding/json"
)

var test_file string		= "testfile.txt"
var test_file_data string	= "dGhpcyBpcyBhIHRlc3QK"
var test_file_checksum string	= "91751cee0a1ab8414400238a761411daa29643ab4b8243e9a91649e25be53ada"
var importedFile *FileData
var tErr error

func TestImportFile(t *testing.T) {
	testData, err := importFile(test_file)
	if err != nil {
		t.Errorf("importFile failed")
	}
	if testData.name != test_file {
		t.Errorf("TestImportFile: functional test failed on testData.name")
	}
	if testData.checksum != test_file_checksum {
		t.Errorf("TestImportFile: functional test failed on testData.checksum")
	}
	if testData.data != test_file_data {
		t.Errorf("TestImportFile: functional test failed on testData.data")
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
}
