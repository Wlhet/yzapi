package yzapi

//https://doc.youzanyun.com/doc#/content/API/1-308
import (
	"encoding/json"
)

//依据有赞openid 或者手机号 获取用户简要信息
func (this *Server) UserBasicInfoGet(phone string) (UserBase, error) {
	rs, err := this.Call("youzan.user.basic.get", "3.0.1", map[string]interface{}{"mobile": phone})
	var user UserBase
	if err != nil {
		return user, err
	}
	err_decode := json.Unmarshal(rs, &user)
	return user, err_decode
}

//通过手机号获取微信用户OPEN ID
func (this *Server) GetWxOpenIdByPhone(phone string) (UserGetWxOpenIdByPhoneMsg, error) {
	rs, err := this.Call("youzan.user.weixin.openid.get", "3.0.0", map[string]interface{}{"mobile": phone})
	var user UserGetWxOpenIdByPhoneMsg
	if err != nil {
		return user, err
	}
	err_decode := json.Unmarshal(rs, &user)
	return user, err_decode
}

//根据union_id查询yz_open_id。 union_id是微信生态下（公众号、小程序、微信登录）的用户唯一标识，
//只有当在同一微信开放平台帐号下用户对应的union_id才相同。 该接口用于三方开发者基于union_id实现多个平台的帐号体系打通
//，如商家A开发了应用B，可基于在应用B上获取到的用户union_id，通过该接口访问用户在有赞上的数据。、
// 使用的前提：用户在union_id在有赞有保存，多种情况下无法获得微信用户的union_id，详见微信开放平台相关文档。
func (this *Server) UserUnionidGet(union_id string) (UserBase, error) {
	rs, err := this.Call("youzan.user.unionid.get", "1.0.0", map[string]interface{}{"union_id": union_id})
	var user UserBase
	if err != nil {
		return user, err
	}
	err_decode := json.Unmarshal(rs, &user)
	return user, err_decode
}

//根据微信openid查询有赞openid
func (this *Server) UserOpenidGetByOpenid(open_id, wechat_type string) (UserBase, error) {
	rs, err := this.Call("youzan.user.openid.getbyopenid", "1.0.0", map[string]interface{}{"wechat_type": wechat_type, "open_id": open_id})
	var user UserBase
	if err != nil {
		return user, err
	}
	err_decode := json.Unmarshal(rs, &user)
	return user, err_decode
}

//获取平台帐号信息
//youzan.user.platform.get.1.0.0
//@parmas:fans_type   平台类型，即粉丝类型，必传，不能传1，查询微信平台信息时，请传9
func (this *Server) UserPlatformGet(yz_open_id string, fans_type int) (UserBase, error) {
	rs, err := this.Call("youzan.user.platform.get", "1.0.0", map[string]interface{}{"yz_open_id": yz_open_id, "fans_type": fans_type})
	var user UserBase
	if err != nil {
		return user, err
	}
	err_decode := json.Unmarshal(rs, &user)
	return user, err_decode
}

//查询是否存在有赞帐号
//youzan.users.account.check.1.0.0
func (this *Server) UsersAccountCheck(account_id, account_type string) (UsersAccountCheckMsg, error) {
	rs, err := this.Call("youzan.users.account.check", "1.0.0", map[string]interface{}{"account_id": account_id, "account_type": account_type})
	var msg UsersAccountCheckMsg
	if err != nil {
		return msg, err
	}
	err_decode := json.Unmarshal(rs, &msg)
	return msg, err_decode
}

//通过手机号查询是否存在有赞帐号
//youzan.users.account.check.1.0.0
func (this *Server) UsersCheckByPhone(account_id string) (UsersAccountCheckMsg, error) {
	rs, err := this.Call("youzan.users.account.check", "1.0.0", map[string]interface{}{"account_id": account_id, "account_type": 2})
	var msg UsersAccountCheckMsg
	if err != nil {
		return msg, err
	}
	err_decode := json.Unmarshal(rs, &msg)
	return msg, err_decode
}

//youzan.scrm.customer.openuser.create.3.0.0
//使用open_user_id创建客户【该接口仅限App开店商家使用】

// /根据微信openid查询有赞openid
//youzan.user.openid.getbyopenid.1.0.0
func (this *Server) UsersOpenIdGetByOpenId(open_id, wechat_type string) (UsersOpenIdGetByOpenIdMsg, error) {
	rs, err := this.Call("youzan.user.openid.getbyopenid", "1.0.0", map[string]interface{}{"open_id": open_id, "wechat_type": wechat_type})
	var msg UsersOpenIdGetByOpenIdMsg
	if err != nil {
		return msg, err
	}
	err_decode := json.Unmarshal(rs, &msg)
	return msg, err_decode
}
