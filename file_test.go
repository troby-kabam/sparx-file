package file

import "testing"

var test_file string = "testfile.txt"
var importedFile fileData = ImportFile(test_file)

func TestEncodedData(t *testing.T) {
	expected := "dGhpcyBpcyBhIHRlc3QK"
	actual := importedFile.data
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}

func TestChecksum(t *testing.T) {
	expected := "91751cee0a1ab8414400238a761411daa29643ab4b8243e9a91649e25be53ada"
	actual := importedFile.checksum
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}

func TestFilename(t *testing.T) {
	expected := test_file
	actual := importedFile.name
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}
