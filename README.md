This go library is designed to convert an arbitrary file type into a transportable payload.
File data is base64-encoded and stored with its checksum and original filename.

A file can be restored, but the original path will be ignored.
The data struct can also be Marshaled into JSON and Unmarshaled back to a Go data struct.

The following interface functions are available:
func New() *FileData
func NewFileData(filename string) (*FileData, error)

The following functions require an input:
UnmarshalJSON(b []byte) (*FileData, error)
func UnmarshalFile(filename string) (*FileData, error)

The following methods are accessible via FileData:
GetName() string
SetName(name string)
GetData() string
SetData(data string)
GetChecksum() string
SetChecksum(sum string)
MarshalJSON() ([]byte, error)
RestoreFile() (int, error)
