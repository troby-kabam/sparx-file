package file

import (
	"os"
	"fmt"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"crypto/sha256"
)

type FileData struct {
	Data string
	Name string
	Checksum string
}

// construct a default FileData struct for manual population
func New() *FileData {
	return &FileData{
		Data:		"aGVsbG8=",
		Checksum:	"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		Name:		"hello.txt",
	}
}

// import file into FileData struct
func NewFileData(filename string) (*FileData, error) {
	newData, err := importFile(filename)
	return &newData, err
}

func importFile(filename string) (FileData, error) {
	var importedFile FileData

	// create buffer from file
	b, err := getBuffer(filename)
	if err != nil {
		return importedFile, err
	}

	// assign values to struct
	importedFile.Data	= base64.StdEncoding.EncodeToString(b)
	importedFile.Name	= filename
	importedFile.Checksum	= makeChecksum(b)

	return importedFile, nil
}

func makeChecksum(b []byte) string {
	h := sha256.New()
	h.Write(b)
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum
}

func getBuffer(filename string) ([]byte, error) {
	b := bytes.Buffer{}
	fp, err := os.Open(filename)
	if err != nil {
		return b.Bytes(), err
	}
	_, err = b.ReadFrom(fp)
	if err != nil {
		return b.Bytes(), err
	}
	err = fp.Close()
	if err != nil {
		return b.Bytes(), err
	}
	return b.Bytes(), nil
}

func (sp *FileData) GetName() string {
	return sp.Name
}

func (sp *FileData) GetData() string {
	return sp.Data
}

func (sp *FileData) GetChecksum() string {
	return sp.Checksum
}

func (sp *FileData) MarshalJSON() ([]byte, error) {
	json_text := FileData{Name: sp.GetName(), Data: sp.GetData(), Checksum: sp.GetChecksum(),}
	b, err := json.Marshal(json_text)
	if err != nil {
		return b, err
	}

	return b, nil
}
