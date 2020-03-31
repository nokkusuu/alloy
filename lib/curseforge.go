package lib

import (
	"fmt"
	"errors"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const baseURL = "https://ddph1n5l22.execute-api.eu-central-1.amazonaws.com/dev/"

func GetModFiles(id string) ([]Mod, error) {
	var requrl string = baseURL + "mod/" + id + "/files/"
	var cresp CurseResp
	var err error

	req, err := http.NewRequest("GET", requrl, nil)
	req.Header.Add("User-Agent", UserAgent)

	resp, err := HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return []Mod{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Mod{}, err
	}

	err = json.Unmarshal(body, &cresp)
	if err != nil {
		return []Mod{}, err
	}
	if cresp.Status != "ok" {
		fmt.Println(string(body))
		return []Mod{}, errors.New("got non-ok response from curseforge api")
	}

	return cresp.Result, nil
}