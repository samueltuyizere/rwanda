package rwanda

import (
	"encoding/json"
	"io"
	"os"
)

type AllData struct {
	Provinces []Province
}

var Data *AllData

func LoadData() {
	jsonFile, err := os.Open("./data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &Data)
}
