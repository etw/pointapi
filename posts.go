package pointapi

import (
	"fmt"
	"strings"
)

func (api *PointAPI) GetAll(before int) (map[string]interface{}, error) {
	var path = fmt.Sprintf("/all?before=%i", before)
	return api.metaGet(path)
}

func (api *PointAPI) GetTags(before int, tags []string) (map[string]interface{}, error) {
	var path = fmt.Sprintf("/tags?before=%i&tag=%s", before, strings.Join(tags, "&tag="))
	return api.metaGet(path)
}
