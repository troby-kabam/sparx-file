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

func EncodeFile(filename string) (string, error) {
	var encodedData string
	fp, err := os.Open(filename)
	if err != nil {
		msg := fmt.Sprintf("cannot open %s", filename)
		return msg, err
	}

	b := bytes.Buffer{}
	count, err := b.ReadFrom(fp)
	if err != nil {
		msg := fmt.Sprintf("ReadFrom error")
		return msg, err
	}
	msg := fmt.Sprintf("bytes read: %d\n", count)
	os.Stderr.WriteString(msg)
	err = fp.Close()
	if err != nil {
		msg := fmt.Sprintf("error closing file")
		return msg, err
	}
	encodedData = base64.StdEncoding.EncodeToString(b.Bytes())
	return encodedData, nil
}

func GetChecksum(encodedData string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		msg := "decoding failed"
		return msg, err
	}
	h := sha256.New()
	h.Write(decodedData)
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum, nil
}
