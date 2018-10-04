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
	data string
	name string
	checksum string
}

// construct a default FileData struct for manual population
func New() *FileData {
	return &FileData{
		data:		"aGVsbG8=",
		checksum:	"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		name:		"hello.txt",
	}
}

// import file into FileData struct
func NewFileData(filename string) (*FileData, error) {
	newData, err := importFile(filename)
	return &newData, err
}

func importFile(filename string) (FileData, error) {
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

func (sp *FileData) MarshalJSON() ([]byte, error) {
	json_text := fmt.Sprintf("{\"name\": \"%s\", \"data\": \"%s\", \"checksum\": \"%s\"}", sp.GetName(), sp.GetData(), sp.GetChecksum())
	b, err := json.Marshal(json_text)
	if err != nil {
		return b, err
	}
	return b, nil
}
