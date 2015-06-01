package pointapi

import (
	"fmt"
	"strings"
)

// Method to get posts form https://point.im/all
func (api *PointAPI) GetAll(before int) (PostList, error) {
	var path = fmt.Sprintf("/all?before=%i", before)
	return api.metaGet(path)
}

// Method to get posts filtered by tags
func (api *PointAPI) GetTags(before int, tags []string) (PostList, error) {
	var path = fmt.Sprintf("/tags?before=%s&tag=%s", before, strings.Join(tags, "&tag="))
	return api.metaGet(path)
}
