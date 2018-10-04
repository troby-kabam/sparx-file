package file

import (
	"os"
	"fmt"
	"bytes"
	"encoding/base64"
	"crypto/sha256"
)

type FileData struct {
	data string
	name string
	checksum string
}

func New() *FileData {
	return &FileData{
		data:		"aGVsbG8=",
		checksum:	"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		name:		"hello.txt",
	}
}

func NewFileData(filename string) (*FileData, error) {
	newData, err := ImportFile(filename)
	return &newData, err
}

func ImportFile(filename string) (FileData, error) {
	var importedFile FileData
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

func (sp *FileData) GetName() string {
	return sp.name
}

func (sp *FileData) GetData() string {
	return sp.data
}

func (sp *FileData) GetChecksum() string {
	return sp.checksum
}
