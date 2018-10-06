package file

import "encoding/json"

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

	return b, nil
}

func UnmarshalJSON(b []byte) (*FileData, error) {
	unmarshaledData := New()
	err := json.Unmarshal(b, unmarshaledData)
	if err != nil {
		return nil, err
	}
	return unmarshaledData, nil
}

func (sp *FileData) RestoreFile(filename string) (int, error) {
	if filename == "" {
		filename = sp.Name
	}
	decoded, err := decodeData(sp.GetData())
	if err != nil {
		return 0, err
	}
	count, err := writeFile(filename, decoded)
	if err != nil {
		return 0, err
	}
	return count, nil
}
