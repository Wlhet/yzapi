//客户API
package yzapi

import (
	"encoding/json"
)

//创建客户，手机号必填
//https://open.youzanyun.com/api/youzan.scrm.customer.create/3.0.0
//@params:mobile    			String    客户手机号    必填
//@params:remark    			String    客户信息备注
//@params:birthday  			String    生日(日期格式:yyyy-MM-dd HH:mm:ss)   1988-05-13 00:00:00
//@params:name    				String    客户姓名
//@params:gender    			Int  	  客户性别
//@params:ascription_kdt_id    	Float64   归属分店(使用的是有赞的店铺id)
func (this *Server) CustomerCreate(mobile string, remark, birthday, name string, gender int) (CustomerCreateMsg, error) {
	params := make(map[string]interface{})
	params["mobile"] = mobile
	params["customer_create"] = map[string]interface{}{
		"remark":   remark,
		"birthday": birthday,
		"name":     name,
		"gender":   gender,
	}
	rs, err := this.Call("youzan.scrm.customer.create", "3.0.0", params)
	var msg CustomerCreateMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//根据店铺信息、身份、成为客户/会员的时间等条件获取客户列表
//https://open.youzanyun.com/api/youzan.scrm.customer.search/3.1.2
//@params:is_member   		int       是否为会员    必填 0表示非会员，1表示会员
//@params:page    			Int    	  页码，最多支持500页(500页是以每页默认值20为单位，客户查询最大为10000)
//@params:page_size 		Int       每页数量，最多支持50个
func (this *Server) CustomerSearch(is_member, page, page_size int) (CustomerSearchMsg, error) {
	params := make(map[string]interface{})
	params["is_member"] = is_member
	params["page"] = page
	params["page_size"] = page_size
	rs, err := this.Call("youzan.scrm.customer.search", "3.1.2", params)
	var msg CustomerSearchMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//更新客户备注，支持多种身份类型，如：手机号，商家自有粉丝（fanstype为1的），yzUid（有赞注册手机号生成的账号ID）
//https://open.youzanyun.com/api/youzan.scrm.customer.remark.update/3.0.0
//@params:remark   		String   更新备注信息    必填
//@params:account_id   	String   帐号ID    必填
//@params:account_type  String   帐号类型。 FansID：自有粉丝ID， Mobile：手机号， YouZanAccount：有赞账号    必填
func (this *Server) CustomerRemarkUpdate(remark, account_id, account_type string) (CustomerRemarkUpdateMsg, error) {
	params := make(map[string]interface{})
	account := map[string]string{
		"account_id":   account_id,
		"account_type": account_type,
	}
	params["remark"] = remark
	params["account"] = account
	rs, err := this.Call("youzan.scrm.customer.search", "3.1.2", params)
	var msg CustomerRemarkUpdateMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//更新客户，支持多种身份类型，如：手机号，商家自有粉丝（fanstype为1的），yzUid（有赞注册手机号生成的账号ID）
//https://open.youzanyun.com/api/youzan.scrm.customer.update/3.0.0
//account
//@params:account_id   	String   帐号ID    必填     若为Mobile 此处就填写手机号
//@params:account_type  String   帐号类型。 FansID：自有粉丝ID， Mobile：手机号， YouZanAccount：有赞账号    必填
//customer_update
//@params:birthday      String   生日      选填     1988-05-13 00:00:00
//@params:gender        int      性别      必填     性别，0：未知；1：男；2：女
//@params:name          String   姓名	   选填
func (this *Server) CustomerUpdate(account_id, account_type string, gender int, customer map[string]interface{}) (CustomerUpdateMsg, error) {
	params := make(map[string]interface{})
	account := map[string]string{
		"account_id":   account_id,
		"account_type": account_type,
	}
	customer_update := map[string]interface{}{
		"gender": gender,
	}
	for k, v := range customer {
		customer_update[k] = v
	}
	params["account"] = account
	params["customer_update"] = customer_update
	rs, err := this.Call("youzan.scrm.customer.update", "3.0.0", params)
	var msg CustomerUpdateMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//@params:account_id   	String   帐号ID    必填     若为Mobile 此处就填写手机号
//@params:account_type  String   帐号类型。 FansID：自有粉丝ID， Mobile：手机号， YouZanAccount：有赞账号    必填
//youzan.scrm.customer.get.3.1.0
func (this *Server) CustomerGet(account_id, account_type string) (CustomerGetMsg, error) {
	params := make(map[string]interface{})
	params["account_id"] = account_id
	params["account_type"] = account_type
	rs, err := this.Call("youzan.scrm.customer.get", "3.1.0", params)
	var msg CustomerGetMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//解密二维码一维码。 场景：第三方解码C端生成的会员码、权益码
//@params:keyword    String 	加密后的关键字
//youzan.scrm.member.code.decode.1.0.0
func (this *Server) MemberCodeDecode(keyword string) (MemberCodeDecodeMsg, error) {
	params := make(map[string]interface{})
	params["keyword"] = keyword
	rs, err := this.Call("youzan.scrm.member.code.decode", "1.0.0", params)
	var msg MemberCodeDecodeMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//
//@params:yz_open_id    String 有赞用户id，用户在有赞的唯一id
//@params:fields    string   需要获得的用户属性，使用,隔开，包含以下user_base（基础信息）,tags（标签）,benefit_cards（权益卡）,benefit_level（等级）,benefit_rights（权益）,credit（积分）,behavior（交易行为）,giftcard（礼品卡）,prepaid（储值）,coupon（优惠券）
//查询客户详情，包含客户基础信息，标签，会员卡，等级，权益，积分，交易行为，礼品卡，储值，优惠券信息，自定义信息项
// func (this *Server) CustomerDetailGet(yz_open_id, fields string) {

// }
