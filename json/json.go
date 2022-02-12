package json

type JsonFile interface {
	ValidatePath() error
	Read() (*[]byte, error)
	GetPath() string
	GetData() *[]byte
	GetNodeData(nodePath string) *[]byte
}

type jsonFile struct {
	path string
	data *[]byte
}

func NewJsonFile(path string) JsonFile {
	return &jsonFile{
		path: path,
		data: nil,
	}
}
