package file

import "testing"

var test_file string		= "testfile.txt"
var test_file_data string	= "dGhpcyBpcyBhIHRlc3QK"
var test_file_checksum string	= "91751cee0a1ab8414400238a761411daa29643ab4b8243e9a91649e25be53ada"
var importedFile *FileData
var tErr error

func TestImportFile(t *testing.T) {
	testData, err := ImportFile(test_file)
	if err != nil {
		t.Errorf("ImportFile failed")
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
