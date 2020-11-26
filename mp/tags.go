package mp

import (
	"encoding/json"
	"fmt"
	"github.com/yanming-open/wechat/common"
	"github.com/yanming-open/wechat/utils"
)

type Tag struct {
	common.BizResponse
	Id    int64  `json:"id"`
	Name  string `json:"name,omitempty"`
	Count int64  `json:"count,omitempty"`
}

func (this *Mp) CreateTag(name string) (tag *Tag, err error) {
	url := fmt.Sprintf("%stags/create?access_token=%s", wxApiHost, this.accessToken)
	params := utils.KV{}
	params["tag"] = utils.KV{"name": name}
	buf, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
	} else {
		err = json.Unmarshal(buf, &tag)
		if err != nil {
			logger.Error(err.Error())
		} else {
			if tag.ErrCode != 0 {
				logger.Error(tag.ErrMsg)
			} else {
				tag.Name = name
			}
		}
	}
	return
}

func (this *Mp) DeleteTag(id int64) {

}
