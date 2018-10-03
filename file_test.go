package file

import "testing"

var test_file string = "testfile.txt"
var encoded_data string

func TestEncodeFile(t *testing.T) {
	expected := "dGhpcyBpcyBhIHRlc3QK"
	actual, _ := EncodeFile(test_file)
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
	encoded_data = actual
}

func TestGetChecksum(t *testing.T) {
	expected := "91751cee0a1ab8414400238a761411daa29643ab4b8243e9a91649e25be53ada"
	actual, _ := GetChecksum(encoded_data)
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}
