package pointapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"errors"
)

const apiPrefix = "https://point.im/api"

type PointAPI struct {
	httpClient *http.Client
	auth       *string
}

func New(c *http.Client, a *string) *PointAPI {
	api := new(PointAPI)
	api.httpClient = c
	api.auth = a
	return api
}

func (api *PointAPI) metaGet(s string) (map[string]interface{}, error) {
	var resmap map[string]interface{}
	resp, err := api.httpClient.Get(fmt.Sprint(apiPrefix, s))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &resmap)
	if err != nil {
		return nil, err
	}
	return resmap, nil
}
