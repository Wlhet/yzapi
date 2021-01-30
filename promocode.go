package yzapi

import (
	"encoding/json"
)

//根据凭证状态查询用户凭证列表，凭证状态：0.全部1.生效=未使用and未过期2.已使用3.已过期=未使用and已过期4.已失效=已使用or已过期
//其中 1、yz_open_id；2、mobile；3、fans_id和fans_type三组参数中必须要传一个。（fans_id和fans_type组成一个唯一的有赞用户标识。）
//https://open.youzanyun.com/api/youzan.ump.voucher.query/3.0.0
//@parmas: activity_type_group   int   活动类型分组，1：优惠券，2：优惠码
//@parmas: page_num  int   起始页码：1
//@parmas: page_size  int   每页数量最大值：200
//
func (this *Server) UmpVoucherQuery(activity_type_group, status, page_num, page_size int, mobile string) (VoucherQueryMsg, error) {
	params := make(map[string]interface{})
	params["activity_type_group"] = activity_type_group
	params["status"] = status
	params["mobile"] = mobile
	params["page_num"] = page_num
	params["page_size"] = page_size
	rs, err := this.Call("youzan.ump.voucher.query", "3.0.0", params)
	// fmt.Print("返回数据", string(rs))
	var msg VoucherQueryMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//作废用户领取的优惠凭证
//https://open.youzanyun.com/api/youzan.ump.voucher.disuse/1.0.0
//https://doc.youzanyun.com/doc#/content/API/1-341/detail/api/0/928
//@parmas coupon_id   优惠券/码Id
//@parmas operator_id 操作人账号id
//@parmas activity_id  优惠券/码活动id
//@parmas coupon_type  券码类型0券1码
func (this *Server) UmpVoucherDisuse(coupon_id, operator_id, activity_id int64, coupon_type int) (UmpVoucherDisuseMsg, error) {
	params := make(map[string]interface{})
	params["coupon_id"] = coupon_id
	params["operator_id"] = operator_id
	params["activity_id"] = activity_id
	params["coupon_type"] = coupon_type
	rs, err := this.Call("youzan.ump.voucher.disuse", "1.0.0", params)

	var msg UmpVoucherDisuseMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//youzan.ump.coupon.take
//发放优惠券
//youzan.ump.coupon.take.3.0.0
//https://doc.youzanyun.com/doc#/content/API/1-341/detail/api/0/155
func (this *Server) UmpCouponTake(mobile string, coupon_group_id int64) (UmpCouponTakeMsg, error) {
	params := make(map[string]interface{})
	params["mobile"] = mobile
	params["coupon_group_id"] = coupon_group_id
	rs, err := this.Call("youzan.ump.coupon.take", "3.0.0", params)
	var msg UmpCouponTakeMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//获取所有未结束的优惠卷/码活动列表
//youzan.ump.coupons.unfinished.search.3.0.0
//https://doc.youzanyun.com/doc#/content/API/1-341/detail/api/0/160
func (this *Server) UmpCouponsUnfinishedSearch() (UmpCouponsUnfinishedSearchMsg, error) {
	params := make(map[string]interface{})
	rs, err := this.Call("youzan.ump.coupons.unfinished.search", "3.0.0", params)
	var msg UmpCouponsUnfinishedSearchMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//优惠券/码活动停止发放(失效)
//https://open.youzanyun.com/api/youzan.ump.voucheractivity.suspendsend/3.0.0
//https://doc.youzanyun.com/doc#/content/API/1-341/detail/api/0/351
func (this *Server) UmpVoucheractivitySuspendsend(operator_type int, app_name string, operator_id, id int64) (UmpVoucheractivitySuspendsendMsg, error) {
	params := make(map[string]interface{})
	request := make(map[string]interface{})
	request["operator_type"] = operator_type
	request["app_name"] = app_name
	request["operator_id"] = operator_id
	request["id"] = id
	params["request"] = request
	rs, err := this.Call("youzan.ump.coupons.unfinished.search", "3.0.0", params)
	var msg UmpVoucheractivitySuspendsendMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}

//创建优惠券活动
//https://open.youzanyun.com/api/youzan.ump.promocard.add/3.0.1
// https://doc.youzanyun.com/doc#/content/API/1-341/detail/api/0/605
func (this *Server) UmpPromoCardAdd(title, start, end string, at_least float64, price int64) (UmpPromoCardAddMsg, error) {
	// params := make(map[string]interface{})
	request := make(map[string]interface{})
	var if_at_least int64
	if at_least > 0 {
		if_at_least = 1
	} else {
		if_at_least = 0
	}
	request["is_at_least"] = if_at_least //是否有门槛
	request["end_at"] = end              //结束时间
	request["date_type"] = 1             //优惠券有效时间类型，1为固定有效期（配合start_at和end_at这两个字段使用）；2为相对有效期
	request["preferential_type"] = 1     //优惠方式，1表示代金，2表示折扣。
	request["title"] = title             //优惠券标题，不超过10个字符。
	request["start_at"] = start          //开始时间
	request["total"] = 10000             //表示优惠券活动的总库存，即可以发放的最大数量，范围为1-100000000
	request["quota"] = 1                 //每个用户能够领取该活动下优惠券的最大次数，为 0 时则表示无次数限制。系统默认上限为500
	request["cent_value"] = price        //优惠券面额  100=1元
	request["range_type"] = "ALL"        //可选值：ALL，全部商品可用；PART，部分指定商品可用；EXCLUDED，部分指定商品不可用。
	request["at_least"] = at_least       //表示该优惠券的门槛。单位：元，精确到分
	request["need_user_level"] = 0       //是否限制领取优惠券用户的会员等级
	request["is_share"] = 0              //是否允许分享

	// params["request"] = request
	rs, err := this.Call("youzan.ump.promocard.add", "3.0.1", request)
	var msg UmpPromoCardAddMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}
