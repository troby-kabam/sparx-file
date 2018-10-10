package file

import (
	"errors"
	"encoding/json"
)

func (sp *FileData) GetName() string {
	return sp.Name
}

func (sp *FileData) SetName(name string) {
	sp.Name = name
}

func (sp *FileData) GetData() string {
	return sp.Data
}

func (sp *FileData) SetData(data string) {
	sp.Data = data
}

func (sp *FileData) GetChecksum() string {
	return sp.Checksum
}

func (sp *FileData) SetChecksum(sum string) {
	sp.Checksum = sum
}

func (sp *FileData) MarshalJSON() ([]byte, error) {
	json_text := FileData{Name: sp.GetName(), Data: sp.GetData(), Checksum: sp.GetChecksum(),}
	b, err := json.Marshal(json_text)
	if err != nil {
		return b, err
	}

	if json.Valid(b) != true {
		err := errors.New("json.Marshal() returned invalid json")
		return b, err
	}
	return b, nil
}

func UnmarshalJSON(b []byte) (*FileData, error) {
	if json.Valid(b) != true {
		err := errors.New("no valid json found")
		return nil, err
	}

	unmarshaledData := New()
	err := json.Unmarshal(b, unmarshaledData)
	if err != nil {
		return nil, err
	}
	return unmarshaledData, nil
}

func UnmarshalFile(filename string) (*FileData, error) {
	b, err := getBuffer(filename)
	if err != nil {
		return nil, err
	}
	return UnmarshalJSON(b)
}

func (sp *FileData) RestoreFile() (int, error) {
	decoded, err := decodeData(sp.GetData())
	if err != nil {
		return 0, err
	}
	count, err := writeFile(sp.GetName(), decoded)
	if err != nil {
		return 0, err
	}
	return count, nil
}
