package yzapi

import "fmt"

//服务结构体
type Server struct {
	AuthorizeType string `json:"authorize_type"`
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	GrantID       int64  `json:"grant_id"`
	Refresh       bool   `json:"refresh"`
	AccessToken
}

//token 结构体
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires"`
	Scope        string `json:"scope"`
}

//基础返回信息结构体
type BaseMsg struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	ErrResp struct {
		ErrMsg  string `json:"err_msg"`
		ErrCode int    `json:"err_code"`
	} `json:"gw_err_resp"`
}

//将返回信息格式化处理
func (this *BaseMsg) ToLogString() string {
	return fmt.Sprintf("return//code:%d Msg:%s  Err:%s", this.Code, this.Message, this.ErrResp.ErrMsg)
}

//获取token 返回信息结构体
type AccessTokenMsg struct {
	Success     bool   `json:"success"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	AccessToken `json:"data"`
}

//用户基础信息结构体
type UserBase struct {
	Data struct {
		Mobile      string `json:"mobile"`     //手机号
		Avatar      string `json:"avatar"`     //头像地址
		NickName    string `json:"nick_name"`  //昵称
		YzOpenId    string `json:"yz_open_id"` //有赞OpendID
		CountryCode string `json:"country_code"`
		Gender      int    `json:"gender"`
		Province    string `json:"province"`
		FansType    int    `json:"fans_type"` //java.lang.Integer 平台类型s
	} `json:"data"`
	BaseMsg
}

//创建用户返回消息结构体
type CustomerCreateMsg struct {
	BaseMsg
	Data struct {
		AccoutType string `json:"account_type"`
		AccoutId   string `json:"account_id"`
		YzOpenId   string `json:"yz_open_id"`
	} `json:"data"`
}

//根据店铺信息、身份、成为客户/会员的时间等条件获取客户列表  返回信息结构体
type CustomerSearchMsg struct {
	BaseMsg
	Data struct {
		Total      float64 `json:"total"`
		RecordList []struct {
			WeixinFansId    int64  `json:"weixin_fans_id"`    //微信粉丝ID
			FansId          int64  `json:"fans_id"`           //粉丝ID
			CreatedAt       int64  `json:"created_at"`        //用户创建时间
			MemberCreatedAt int64  `json:"member_created_at"` //用户成为会员时间
			Gender          int    `json:"gender"`            //性别，0:其他 1:男 2:女
			IsMember        int    `json:"is_member"`         //是否是会员，0：不是 1：是
			TradeCount      int    `json:"trade_count"`
			ShowName        string `json:"show_name"`
			Mobile          string `json:"mobile"`
			Name            string `json:"name"`
			YzUid           int64  `json:"yz_uid"`
			Points          int64  `json:"points"`
		} `json:"record_list"`
	} `json:"data"`
}

//更新客户备注返回信息结构体
type CustomerRemarkUpdateMsg struct {
	BaseMsg
	Data bool `json:"data"`
}

//更新客户信息 返回信息结构体
type CustomerUpdateMsg struct {
	BaseMsg
	Data bool `json:"data"`
}

//查询用户优惠券码 返回信息结构体
type VoucherQueryMsg struct {
	BaseMsg
	Total    int `json:"total"`
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
	Data     []struct {
		KdtId                  int64    `json:"kdt_id"`      //店铺在有赞的id标识
		SendSource             string   `json:"send_source"` //发放来源
		SendAt                 int64    `json:"send_at"`
		ValidStartTime         int64    `json:"valid_start_time"` //优惠券/码活动开始时间，Unix时间戳，单位：毫秒
		ValidEndTime           int64    `json:"valid_end_time"`
		VerifiedAt             int64    `json:"verified_at"`
		ActivityId             int64    `json:"activity_id"`              //优惠券/码活动id
		Value                  int64    `json:"value"`                    //面额，单位：分
		VerifyCode             string   `json:"verify_code"`              //买家端优惠券/码/码核销码
		OrderPreferentialValue int64    `json:"order_preferential_value"` //凭证在订单中实际优惠金额，单位：分
		VerifiedInOrderNos     []string `json:"verified_in_order_nos"`    //订单列表
		IsSyncCard             bool     `json:"is_sync_card"`             //是否同步微信卡券 0：否 1：是
		PreferentialMode       int      `json:"preferential_mode"`        //优惠方式，1：代金券，2：折扣券，3：兑换券
		VerifyOrderSource      string   `json:"verify_order_source"`      //核销单来源（"YZ_TRADE", "有赞线上交易订单号"）
		VoucherIdentity        struct {
			CouponId   int64 `json:"coupon_id"`   //优惠券ID
			CouponType int   `json:"coupon_type"` //优惠券：0优惠券，1：优惠码
		} `json:"voucher_identity"`
		UserIdentity struct {
			HasBoundMobile bool   `json:"has_bound_mobile"`
			Mobile         string `json:"mobile"`
			YzOpenId       string `json:"yz_open_id"`
			ProxyFansId    int64  `json:"proxy_fans_id"`
			FansId         int64  `json:"fans_id"`
		} `json:"user_identity"`
	} `json:"data"`
}

type UmpVoucherDisuseMsg struct {
	BaseMsg
	Data bool `json:"data"`
}
type ShopUserMsg struct {
	BaseMsg
}

//获取客户信息返回结构体
type CustomerGetMsg struct {
	BaseMsg
	Data struct {
		Mobile    string //手机号
		Remark    string
		Gender    int
		IsMember  bool   `json:"is_member"`
		Name      string `json:"name"`
		OutUserId string `json:"out_user_id"`
		Birthday  string `json:"birthday"` //2020-02-02
	} `json:"data"`
}

//C方解码获取用户ID返回结构体
type MemberCodeDecodeMsg struct {
	BaseMsg
	Data struct {
		Mobile   string `json:"mobile"`
		YzOpenId string `json:"yz_open_id"`
	} `json:"data"`
}

//查询用户是否存在返回结构体
type UsersAccountCheckMsg struct {
	BaseMsg
	Data bool `json:"data"`
}

//微信opendId换取有赞OPENID 返回结构体
type UsersOpenIdGetByOpenIdMsg struct {
	BaseMsg
	Data struct {
		YzOpenId string `json:"yz_open_id"`
	} `json:"data"`
}

//UmpCouponTake
type UmpCouponTakeMsg struct {
	BaseMsg
	Data struct {
		CouponType string `json:"coupon_type"`
		Promocard  struct {
			StartAt          int64  `json:"start_at"`
			IsUsed           int    `json:"is_used"`
			PreferentialType int    `json:"preferential_type"`
			Value            string `json:"value"`
			IsInvalid        int    `json:"is_invalid"`
			Discount         int    `json:"discount"`
			DetailUrl        string `json:"detail_url"`
			PromocardId      int64  `json:"promocard_id"`
			Condition        string `json:"condition"`
			Title            string `json:"title"`
			BackgroundColor  string `json:"background_color"`
			EndAt            int64  `json:"end_at"`
			IsExpired        int    `json:"is_expired"`
			VerifyCode       string `json:"verify_code"`
		} `json:"promocard"`
		Promocode struct {
			StartAt          int64  `json:"start_at"`
			IsUsed           int    `json:"is_used"`
			PreferentialType int    `json:"preferential_type"`
			Value            string `json:"value"`
			IsInvalid        int    `json:"is_invalid"`
			Discount         int    `json:"discount"`
			DetailUrl        string `json:"detail_url"`
			PromocardId      int64  `json:"promocard_id"`
			Condition        string `json:"condition"`
			Title            string `json:"title"`
			BackgroundColor  string `json:"background_color"`
			EndAt            int64  `json:"end_at"`
			IsExpired        int    `json:"is_expired"`
			VerifyCode       string `json:"verify_code"`
		} `json:"promocode"`
	} `json:"data"`
}

type UmpCouponsUnfinishedSearchMsg struct {
	//https://doc.youzanyun.com/doc#/content/API/1-341/detail/api/0/160
	BaseMsg
	Data struct {
		Coupons []struct {
			GroupId            int64  `json:"group_id"`
			StatFetchNum       int    `json:"stat_fetch_num"`
			UserLevelName      string `json:"user_level_name"`
			IsForbidPreference int    `json:"is_forbid_preference"`
			Total              int    `json:"total"`
			Created            int64  `json:"created"`
			Value              string `json:"value"`
			Description        string `json:"description"`
			WeixinCardId       string `json:"weixin_card_id"`
			Quota              int    `json:"quota"`
			ValueRandomTo      string `json:"value_random_to"`
			Stock              int    `json:"stock"`
			EndAt              int64  `json:"end_at"`
			IsShare            int    `json:"is_share"`
			NeedUserLevel      int    `json:"need_user_level"`
			Updated            int64  `json:"updated"`
			IsAtLeast          int    `json:"is_at_least"`
			Title              string `json:"title"`
			StartAt            int64  `json:"start_at"`
			StatUseNum         int    `json:"stat_use_num"`
			IsRandom           int    `json:"is_random"`
			StatFetchUserNum   int    `json:"stat_fetch_user_num"`
			IsSyncWeixin       int    `json:"is_sync_weixin"`
			RangeType          string `json:"range_type"`
			Status             int    `json:"status"`
			AtLeast            string `json:"at_least"`
			ExpireNotice       int    `json:"expire_notice"`
			FetchUrl           string `json:"fetch_url"`
			CouponType         string `json:"coupon_type"`
		} `json:"coupons"`
	} `json:"data"`
}

type UmpVoucheractivitySuspendsendMsg struct {
	BaseMsg
	Data bool `json:"data"`
}

type YouzanShopStaffQueryMsg struct {
	BaseMsg
	Data struct {
		TotalCount int64   `json:"total_count"`
		AdminIds   []int64 `json:"admin_ids"`
	} `json:"data"`
}

//为用户增加积分返回信息
type CrmCustomerPointsIncreaseMsg struct {
	BaseMsg
	Data struct {
		IsSuccess string `json:"is_success"` //是否成功  "false/true"
	} `json:"data"`
}

//为用户减少积分返回信息
type CrmCustomerPointsDecreaseMsg struct {
	BaseMsg
	Data struct {
		IsSuccess string `json:"is_success"` //是否成功  "false/true"
	} `json:"data"`
}

//用户积分变动日志
type CrmCustomerPointsChangelogSearchMsg struct {
	BaseMsg
	Items struct {
		GoodsTitle      string `json:"goods_title"`
		StatementNo     string `json:"statement_no"`
		CostPrice       int64  `json:"cost_price"`
		GoodsCount      int    `json:"goods_count"`
		CreatedTime     int64  `json:"created_time"`
		ReducedPrice    int    `json:"reduced_price"`
		CouponGroupType int    `json:"coupon_group_type"`
		GoodsType       int    `json:"goods_type"`
		GoodsId         int64  `json:"goods_id"`
		Mobile          string `json:"mobile"`
		GoodsImg        string `json:"goods_img"`
		CostPoints      int    `json:"cost_points"`
	} `json:"items"`
	Paginator struct {
		TotalCount int `json:"total_count"`
		PageSize   int `json:"page_size"`
		Page       int `json:"page"`
	} `json:"paginator"`
}

//获取用户积分
type CrmCustomerPointsGetMsg struct {
	BaseMsg
	Data struct {
		Point int64 `json:"point"`
	} `json:"data"`
}

//同步用户积分
type CrmCustomerPointsSyncMsg struct {
	BaseMsg
	Data struct {
		IsSuccess string `json:"is_success"` //是否成功  "false/true"
	} `json:"data"`
}
type UserGetWxOpenIdByPhoneMsg struct {
	BaseMsg
	Data struct {
		OpenId  string `json:"open_id"`
		Unionid string `json:"union_id"`
	} `json:"data"`
}

type UmpPromoCardAddMsg struct {
	BaseMsg
	Data struct {
		Promocard struct {
			DateType int64  `json:"date_type"`
			EndAt    int64  `json:"end_at"`
			Title    string `json:"title"`
			StartAt  int64  `json:"start_at"`
			Total    int64  `json:"total"`
			GroupId  int64  `json:"group_id"`
			FetchUrl string `json:"fetch_url"`
			Stock    int64  `json:"stock"`
		} `json:"promocard"`
	} `json:"data"`
}
