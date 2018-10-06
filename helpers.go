package file

import (
	"os"
	"fmt"
	"bytes"
	"encoding/base64"
	"crypto/sha256"
	"path/filepath"
)

func importFile(filename string) (FileData, error) {
	var importedFile FileData

	// create buffer from file
	b, err := getBuffer(filename)
	if err != nil {
		return importedFile, err
	}

	// assign values to struct
	importedFile.Data	= base64.StdEncoding.EncodeToString(b)
	importedFile.Name	= filepath.Base(filename)
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

func decodeData(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return []byte{}, err
	}
	return decoded, nil
}

func writeFile(filename string, data []byte) (int, error) {
	fp, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	count, err := fp.Write(data)
	if err != nil {
		return count, err
	}
	fp.Close()
	return count, nil
}
