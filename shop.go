package yzapi

import "encoding/json"

//https://doc.youzanyun.com/doc#/content/API/1-331/detail/api/0/1390
//https://open.youzanyun.com/api/youzan.shop.staff.query/1.0.0
//分页查询店铺内所有员工唯一id，支持单店铺或连锁模式查询
//分页查询员工详情信息，支持单店（mode="SELF_SHOP"，默认）和连锁范围(mode="CHAIN_SHOP")查询，单店时无需分页
func (this *Server) YouzanShopStaffQuery(page_no, page_size int) (YouzanShopStaffQueryMsg, error) {
	params := map[string]interface{}{
		"chain_query_manage_param_d_t_o": map[string]interface{}{
			"mode":      "SELF_SHOP",
			"page_no":   page_no,
			"page_size": page_size,
		},
	}
	rs, err := this.Call("youzan.shop.staff.query", "1.0.0", params)
	// fmt.Print("返回数据", string(rs))
	var msg YouzanShopStaffQueryMsg
	err = json.Unmarshal(rs, &msg)
	return msg, err
}
