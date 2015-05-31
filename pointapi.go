package pointapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (api *PointAPI) metaGet(s string) (PostList, error) {
	var resmap PostList
	resp, err := api.httpClient.Get(fmt.Sprint(apiPrefix, s))
	if err != nil {
		return resmap, err
	}
	if resp.StatusCode != http.StatusOK {
		return resmap, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resmap, err
	}
	err = json.Unmarshal(body, &resmap)
	if err != nil {
		return resmap, err
	}
	return resmap, nil
}
