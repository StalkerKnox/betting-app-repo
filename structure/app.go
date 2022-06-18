package structure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Parsing JSON
func GetJSON(URL string, structure interface{}) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(resBody, structure)
}
