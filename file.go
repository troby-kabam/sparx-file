package file

import (
	"os"
	"fmt"
	"bytes"
	"encoding/base64"
	"crypto/sha256"
)

type fileData struct {
	data string
	name string
	checksum string
}

func New() *fileData {
	return &fileData{}
}

func NewFileData(filename string) (fileData, error) {
	newData, err := ImportFile(filename)
	return newData, err
}

func ImportFile(filename string) (fileData, error) {
	var importedFile fileData
	fp, err := os.Open(filename)
	if err != nil {
		return importedFile, err
	}

	// create encoded data
	b := bytes.Buffer{}
	_, err = b.ReadFrom(fp)
	if err != nil {
		return importedFile, err
	}
	err = fp.Close()
	if err != nil {
		return importedFile, err
	}
	importedFile.data = base64.StdEncoding.EncodeToString(b.Bytes())

	// create checksum
	h := sha256.New()
	h.Write(b.Bytes())
	sum := fmt.Sprintf("%x", h.Sum(nil))
	importedFile.checksum = sum

	// set filename
	importedFile.name = filename

	return importedFile, nil
}
