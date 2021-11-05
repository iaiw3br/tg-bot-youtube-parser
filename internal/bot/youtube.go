package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
	"tg-bot-youtube-parser/internal/models"
)

func GetLastVideoId(channelUrl string) (string, error) {
	items, err := getVideoByFilter(channelUrl)
	if err != nil {
		return "", err
	}
	if len(items) < 1 {
		return "", errors.New("video not found")
	}

	return items[0].Id.VideoId, nil
}

func getVideoByFilter(channelUrl string) ([]models.Item, error) {
	req, err := makeRequest(channelUrl)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var restResponse models.RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Items, nil
}

func makeRequest(channelUrl string) (*http.Request, error) {
	youtubeSearchUrl := viper.GetString("YOUTUBE_SEARCH_URL")
	lastIndex := strings.LastIndex(channelUrl, "/")
	channelId := channelUrl[lastIndex+1:]
	req, err := http.NewRequest("GET", youtubeSearchUrl, nil)
	if err != nil {
		return nil, err
	}

	youtubeApiToken := viper.GetString("YOUTUBE_API_TOKEN")
	query := req.URL.Query()
	query.Add("part", "id")
	query.Add("channelId", channelId)
	query.Add("maxResults", "1")
	query.Add("order", "date")
	query.Add("key", youtubeApiToken)

	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())

	return req, nil
}
