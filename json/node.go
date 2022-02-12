package json

import "github.com/tidwall/gjson"

func (f *jsonFile) GetNodeData(nodePath string) *[]byte {
	if len(nodePath) == 0 {
		return nil
	}

	node := gjson.GetBytes(*f.GetData(), nodePath)
	if !node.Exists() {
		return nil
	}
	var data []byte
	if node.Index > 0 {
		data = (*f.GetData())[node.Index : node.Index+len(node.Raw)]
	} else {
		data = []byte(node.Raw)
	}
	return &data
}
