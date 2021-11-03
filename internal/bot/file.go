package bot

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetUrls() ([]string, error) {
	jsonFile, err := os.ReadFile("urls.json")
	if err != nil {
		return nil, err
	}

	var urls []string

	err = json.Unmarshal(jsonFile, &urls)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func AddUrl(url string) error {
	jsonFile, err := os.ReadFile("urls.json")
	if err != nil {
		return err
	}

	var urls []string

	err = json.Unmarshal(jsonFile, &urls)
	if err != nil {
		return err
	}

	for _, urlInt := range urls {
		if urlInt == url {
			return nil
		}
	}

	newUrls := append(urls, url)

	rawDataOut, err := json.MarshalIndent(&newUrls, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("urls.json", rawDataOut, 0)
	if err != nil {
		return err
	}
	return nil
}
