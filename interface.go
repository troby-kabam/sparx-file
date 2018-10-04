package file

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

// create a FileData struct with the specified file
func NewFileData(filename string) (*FileData, error) {
	newData, err := importFile(filename)
	return &newData, err
}
