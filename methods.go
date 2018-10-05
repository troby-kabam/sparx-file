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
