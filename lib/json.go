package lib

import (
	"io/ioutil"
	"encoding/json"
)

func OpenPackFile(dest string) (JSONPack, error) {
	var pack JSONPack
	var err error

	f, err := ioutil.ReadFile(dest)
	if err != nil {
		return JSONPack{}, err
	}

	err = json.Unmarshal(f, &pack)
	if err != nil {
		return JSONPack{}, err
	}

	return pack, nil
}