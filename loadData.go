package rwanda

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
)

//go:embed data.json
var dataFS embed.FS

type AllData struct {
	Provinces []Province
}
var PreprocessedData *[]AllData
var Data *AllData

func LoadData() {
    file, err := OpenEmbeddedFile("data.json")
    if err != nil {
        panic(err)
    }

    byteValue, err := io.ReadAll(file)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(byteValue, &PreprocessedData)
    Data = &(*PreprocessedData)[0]
    if err != nil {
        panic(err)
    }
}
func OpenEmbeddedFile(name string) (io.Reader, error) {
    data, err := dataFS.ReadFile(name)
    if err != nil {
        return nil, fmt.Errorf("failed to read embedded file: %w", err)
    }
    return bytes.NewReader(data), nil
}
