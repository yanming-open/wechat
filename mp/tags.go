package mp

import (
	"encoding/json"
	"errors"
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

// 创建tag
func (m *mp) CreateTag(name string) (tag *Tag, err error) {
	url := fmt.Sprintf("%stags/create?access_token=%s", wxApiHost, m.accessToken)
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

// 删除标签
func (m *mp) DeleteTag(id int) (err error) {
	url := fmt.Sprintf("%stags/delete?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["tag"] = utils.KV{"id": id}
	buf, err := utils.DoPost(url, params)
	if err != nil {
		return err
	} else {
		var resp = common.BizResponse{}
		json.Unmarshal(buf, &resp)
		if resp.ErrCode != 0 {
			return errors.New(resp.ErrMsg)
		}
		return
	}
}

// 修改标签
func (m *mp) UpdateTag(id int, name string) (err error) {
	url := fmt.Sprintf("%stags/update?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["tag"] = utils.KV{"id": id, "name": name}
	buf, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
		return
	} else {
		var bizResp = common.BizResponse{}
		err = json.Unmarshal(buf, &bizResp)
		if err != nil {
			logger.Error(err.Error())
			return
		} else {
			if bizResp.ErrCode != 0 {
				err = errors.New(bizResp.ErrMsg)
			}
		}
	}
	return
}

// 获取全部标签列表
func (m *mp) GetTags() (list []Tag) {
	url := fmt.Sprintf("%stags/get?access_token=%s", wxApiHost, m.accessToken)
	buf, err := utils.DoGet(url)
	if err != nil {
		logger.Error(err.Error())
		return
	} else {
		var tagsResp = tagsResponse{}
		json.Unmarshal(buf, &tagsResp)
		list = tagsResp.Tags
	}
	return
}

// 获取标签下粉丝列表
func (m *mp) GetTagUsers(id int, nextOpenId string) (tagusers TagUsersResponse, err error) {
	url := fmt.Sprintf("%suser/tag/get?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["tag"] = utils.KV{"tagid": id, "next_openid": nextOpenId}
	buf, err := utils.DoPost(url, params)
	logger.Info(string(buf))
	if err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &tagusers)
		if err != nil {
			return
		} else {
			if tagusers.ErrCode != 0 {
				err = errors.New(tagusers.ErrMsg)
			}
			return
		}
	}
}

// 批量为用户打标签
func (m *mp) BatchTagGing(id int, openIdList []string) (err error) {
	url := fmt.Sprintf("%stags/members/batchtagging?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["tagid"] = id
	params["openid_list"] = openIdList
	buf, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
		return
	} else {
		var bizResp = common.BizResponse{}
		err = json.Unmarshal(buf, &bizResp)
		if err != nil {
			logger.Error(err.Error())
			return
		} else {
			if bizResp.ErrCode != 0 {
				err = errors.New(bizResp.ErrMsg)
			}
		}
	}
	return
}

// 批量为用户取消标签
func (m *mp) BatchUnTagGing(id int, openIdList []string) (err error) {
	url := fmt.Sprintf("%stags/members/batchuntagging?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["tagid"] = id
	params["openid_list"] = openIdList
	buf, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
		return
	} else {
		var bizResp = common.BizResponse{}
		err = json.Unmarshal(buf, &bizResp)
		if err != nil {
			logger.Error(err.Error())
			return
		} else {
			if bizResp.ErrCode != 0 {
				err = errors.New(bizResp.ErrMsg)
			}
		}
	}
	return
}

// 获取用户身上的标签列表
func (m *mp) GetUserTagIdList(openId string) (idList []int, err error) {
	url := fmt.Sprintf("%stags/getidlist?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["openid"] = openId
	buf, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
		return
	} else {
		var resultResp = tagsIdListResponse{}
		err = json.Unmarshal(buf, &resultResp)
		if err != nil {
			logger.Error(err.Error())
			return
		} else {
			if resultResp.ErrCode != 0 {
				err = errors.New(resultResp.ErrMsg)
			} else {
				idList = resultResp.TagIdList
			}
		}
	}
	return
}
