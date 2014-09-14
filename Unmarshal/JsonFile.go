package Unmarshal

import (
	"encoding/json"
	"io/ioutil"
)

// JsonFile parses the JSON-encoded file and stores the result
// in the value pointed to by dest.
func JsonFile(fileName string, dest interface{}) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, dest)

	if err != nil {
		return err
	}

	return nil
}
