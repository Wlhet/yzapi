//积分API
package yzapi

import (
	"encoding/json"
)

//youzan.crm.customer.points.increase.4.0.0
//给用户加积分 支持的用户账号类型 1-有赞粉丝id(有赞不同的合作渠道会生成不同渠道对应在有赞平台下的fans_id) ; 2-手机号; 3-三方帐号(原open_user_id:三方App用户ID，该参数仅限购买App开店插件的商家使用) ; 5-有赞用户id，用户在有赞的唯一id。推荐使用）
//帐号类型（1-有赞粉丝id(有赞不同的合作渠道会生成不同渠道对应在有赞平台下的fans_id) ; 2-手机号; 3-三方帐号(原open_user_id:三方App用户ID，该参数仅限购买App开店插件的商家使用) ;5-有赞用户id，用户在有赞的唯一id。推荐使用）
func (this *Server) CrmCustomerPointsIncrease(reason string, biz_value string, points int, account_type int, account_id string) (CrmCustomerPointsIncreaseMsg, error) {
	params := map[string]interface{}{
		"params": map[string]interface{}{
			"reason":    reason,
			"biz_value": biz_value,
			"points":    points,
			"user": map[string]interface{}{
				"account_type": account_type,
				"account_id":   account_id,
			},
		},
	}
	rs, err := this.Call("youzan.crm.customer.points.increase", "4.0.0", params)
	var msg CrmCustomerPointsIncreaseMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.points.increase.4.0.0
//给用户加积分 支持的用户账号类型 1-有赞粉丝id(有赞不同的合作渠道会生成不同渠道对应在有赞平台下的fans_id) ; 2-手机号; 3-三方帐号(原open_user_id:三方App用户ID，该参数仅限购买App开店插件的商家使用) ; 5-有赞用户id，用户在有赞的唯一id。推荐使用）
//reasno:加分理由,biz_value:加分记录单号--自定义，points:积分,phone:手机
func (this *Server) PointAddByPhone(reason string, biz_value string, points int, phone string) (CrmCustomerPointsIncreaseMsg, error) {
	params := map[string]interface{}{
		"params": map[string]interface{}{
			"reason":    reason,
			"biz_value": biz_value,
			"points":    points,
			"user": map[string]interface{}{
				"account_type": 2,
				"account_id":   phone,
			},
		},
	}
	rs, err := this.Call("youzan.crm.customer.points.increase", "4.0.0", params)
	var msg CrmCustomerPointsIncreaseMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.points.changelog.search.4.0.0
//查询用户积分日志
//https://doc.youzanyun.com/doc#/content/API/1-371/detail/api/0/885
func (this *Server) CrmCustomerPointsChangelogSearch(end_time, begin_time string, account_type int, account_id string, page, page_size int) (CrmCustomerPointsChangelogSearchMsg, error) {
	params := make(map[string]interface{})
	params["end_time"] = end_time
	params["begin_time"] = begin_time
	params["page"] = page
	params["page_size"] = page_size
	params["user"] = map[string]interface{}{
		"account_type": account_type,
		"account_id":   account_id,
	}
	rs, err := this.Call("youzan.crm.customer.points.changelog.search", "4.0.0", params)
	var msg CrmCustomerPointsChangelogSearchMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.pointstore.exchange.records.2.0.0
//获取用户积分兑换记录
func (this *Server) CrmCustomerPointstoreExchangeRecords(end_time, begin_time int64, page, page_size, goods_type int, account_type int, account_id string) {
	params := make(map[string]interface{})
	params["end_time"] = end_time
	params["begin_time"] = begin_time
	params["page"] = page
	params["page_size"] = page_size
	params["user"] = map[string]interface{}{
		"account_type": account_type,
		"account_id":   account_id,
	}
}

//youzan.crm.customer.points.decrease.4.0.0
//https://doc.youzanyun.com/doc#/content/API/1-371/detail/api/0/873
func (this *Server) CrmCustomerPointsDecrease(reason, biz_value string, points int, account_type int, account_id string) (CrmCustomerPointsDecreaseMsg, error) {
	params := map[string]interface{}{
		"params": map[string]interface{}{
			"reason":    reason,
			"biz_value": biz_value,
			"client_id": this.ClientID,
			"points":    points,
			"user": map[string]interface{}{
				"account_type": account_type,
				"account_id":   account_id,
			},
		},
	}
	rs, err := this.Call("youzan.crm.customer.points.decrease", "4.0.0", params)
	var msg CrmCustomerPointsDecreaseMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.points.decrease.4.0.0
//https://doc.youzanyun.com/doc#/content/API/1-371/detail/api/0/873
//reasno:加分理由,biz_value:加分记录单号--自定义，points:积分,phone:手机
func (this *Server) PointDesByPhone(reason, biz_value string, points int, phone string) (CrmCustomerPointsDecreaseMsg, error) {
	params := map[string]interface{}{
		"params": map[string]interface{}{
			"reason":    reason,
			"biz_value": biz_value,
			"client_id": this.ClientID,
			"points":    points,
			"user": map[string]interface{}{
				"account_type": 2,
				"account_id":   phone,
			},
		},
	}
	rs, err := this.Call("youzan.crm.customer.points.decrease", "4.0.0", params)
	var msg CrmCustomerPointsDecreaseMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.points.get.1.0.0
//查询用户当前积分
func (this *Server) CrmCustomerPointsGet(account_type int, account_id string) (CrmCustomerPointsGetMsg, error) {
	params := map[string]interface{}{}
	params["user"] = map[string]interface{}{
		"account_type": account_type,
		"account_id":   account_id,
	}
	rs, err := this.Call("youzan.crm.customer.points.get", "1.0.0", params)
	var msg CrmCustomerPointsGetMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.points.get.1.0.0
//查询用户当前积分
func (this *Server) PointQueryByPhone(phone string) (CrmCustomerPointsGetMsg, error) {
	params := map[string]interface{}{}
	params["user"] = map[string]interface{}{
		"account_type": 2,
		"account_id":   phone,
	}
	rs, err := this.Call("youzan.crm.customer.points.get", "1.0.0", params)
	var msg CrmCustomerPointsGetMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.points.sync.4.0.0
//同步客户积分（根据传参覆盖掉用户当前积分值，例 A有200积分，传参需要同步后的用户积分为10，则会扣除A 190个积分，将A的积分修改为10） 限制条件：同一店铺同一用户10s内只能操作一次同步积分
func (this *Server) CrmCustomerPointsSync(reason, biz_value string, account_type int, account_id string, points int) (CrmCustomerPointsSyncMsg, error) {
	params := map[string]interface{}{}
	params["reason"] = reason
	params["biz_value"] = biz_value
	params["points"] = points
	params["user"] = map[string]interface{}{
		"account_id": account_id,
		"points":     points,
	}
	rs, err := this.Call("youzan.crm.customer.points.sync", "4.0.0", params)
	var msg CrmCustomerPointsSyncMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.crm.customer.points.sync.4.0.0
//同步客户积分（根据传参覆盖掉用户当前积分值，例 A有200积分，传参需要同步后的用户积分为10，则会扣除A 190个积分，将A的积分修改为10） 限制条件：同一店铺同一用户10s内只能操作一次同步积分
func (this *Server) PointSyncByPhone(reason, biz_value string, account_id string, points int) (CrmCustomerPointsSyncMsg, error) {
	params := map[string]interface{}{}
	params["reason"] = reason
	params["biz_value"] = biz_value
	params["points"] = points
	params["user"] = map[string]interface{}{
		"account_type": 2,
		"account_id":   account_id,
	}
	rs, err := this.Call("youzan.crm.customer.points.sync", "4.0.0", params)
	var msg CrmCustomerPointsSyncMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}
