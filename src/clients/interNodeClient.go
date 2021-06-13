package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/praateekgupta3991/contraption/entities"
)

type InterNodeClient struct {
	HttpClient *http.Client
}

type InterNodeComm interface {
	Chain(nodeUrl string) ([]entities.Block, error)
}

func InitInterNodeClient(hClient *http.Client) *InterNodeClient {
	client := &InterNodeClient{
		HttpClient: hClient,
	}
	return client
}

func (c *InterNodeClient) Chain(nodeUrl string) ([]entities.Block, error) {
	url := fmt.Sprintf("http://%s/%s", nodeUrl, "chain")
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		fmt.Println(url)
		return nil, err
	} else {
		resp, err := c.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		wh := new([]entities.Block)
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else {
			if err := json.Unmarshal(body, &wh); err != nil {
				return nil, err
			} else {
				return *wh, nil
			}
		}
	}
}
