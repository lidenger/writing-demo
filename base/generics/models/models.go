package models

type Resp[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type AccountReq struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,len=11"`
	Avatar   string `json:"avatar"`    // 头像URL
	Nickname string `json:"nickname"`  // 昵称
	Gender   int    `json:"gender"`    // 性别：0-未知 1-男 2-女
	Birthday string `json:"birthday"`  // 生日
	RealName string `json:"real_name"` // 真实姓名
	IDCard   string `json:"id_card"`   // 身份证号
}

type AccountCond struct {
	Gender   *int  `json:"gender"`
	IsEnable *bool `json:"is_enable"`
}

type Account struct {
	UserID     int64   `json:"user_id"`     // 用户ID
	Username   string  `json:"username"`    // 用户名
	Email      string  `json:"email"`       // 邮箱
	Phone      string  `json:"phone"`       // 手机号
	Avatar     string  `json:"avatar"`      // 头像URL
	Nickname   string  `json:"nickname"`    // 昵称
	Gender     int     `json:"gender"`      // 性别
	Birthday   string  `json:"birthday"`    // 生日
	RealName   string  `json:"real_name"`   // 真实姓名（脱敏）
	Status     int     `json:"status"`      // 账号状态：0-正常 1-冻结
	CreateTime string  `json:"create_time"` // 创建时间
	Balance    float64 `json:"balance"`     // 账户余额
	Points     int     `json:"points"`      // 积分
	VIPLevel   int     `json:"vip_level"`   // VIP等级
}

type OrderReq struct {
	UserID    int64        `json:"user_id" binding:"required"`
	GoodsList []OrderGoods `json:"goods_list" binding:"required,min=1"`
	AddressID int64        `json:"address_id" binding:"required"`
	CouponID  int64        `json:"coupon_id"` // 优惠券ID
	Remark    string       `json:"remark"`    // 订单备注
	PayType   int          `json:"pay_type"`  // 支付方式：1-微信 2-支付宝
	Source    int          `json:"source"`    // 订单来源：1-APP 2-PC
}

// 订单商品项
type OrderGoods struct {
	GoodsID   int64   `json:"goods_id" binding:"required"`
	SkuID     int64   `json:"sku_id"` // SKU ID
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	UnitPrice float64 `json:"unit_price" binding:"required"` // 单价
}

type Order struct {
	OrderID     string       `json:"order_id"`     // 订单号
	OrderStatus int          `json:"order_status"` // 订单状态：0-待支付 1-已支付 2-已发货 3-已完成 4-已取消
	TotalAmount float64      `json:"total_amount"` // 订单总额
	Discount    float64      `json:"discount"`     // 优惠金额
	PayAmount   float64      `json:"pay_amount"`   // 实付金额
	GoodsList   []OrderGoods `json:"goods_list"`   // 商品列表
	Address     OrderAddress `json:"address"`      // 收货地址
	PayType     int          `json:"pay_type"`     // 支付方式
	PayTime     string       `json:"pay_time"`     // 支付时间
	CreateTime  string       `json:"create_time"`  // 创建时间
	ExpireTime  string       `json:"expire_time"`  // 支付过期时间
	LogisticsNo string       `json:"logistics_no"` // 物流单号
	LogisticsCo string       `json:"logistics_co"` // 物流公司
}

// 订单地址信息
type OrderAddress struct {
	Name       string `json:"name"`        // 收货人姓名
	Phone      string `json:"phone"`       // 收货人电话
	Province   string `json:"province"`    // 省
	City       string `json:"city"`        // 市
	District   string `json:"district"`    // 区
	Detail     string `json:"detail"`      // 详细地址
	PostalCode string `json:"postal_code"` // 邮政编码
}

type GoodsReq struct {
	CategoryID int64   `json:"category_id"`                       // 分类ID
	BrandID    int64   `json:"brand_id"`                          // 品牌ID
	Keyword    string  `json:"keyword"`                           // 搜索关键词
	MinPrice   float64 `json:"min_price"`                         // 最低价格
	MaxPrice   float64 `json:"max_price"`                         // 最高价格
	SortType   int     `json:"sort_type"`                         // 排序方式：0-默认 1-价格升序 2-价格降序 3-销量降序
	Page       int     `json:"page" binding:"min=1"`              // 页码
	PageSize   int     `json:"page_size" binding:"min=1,max=100"` // 每页数量
	Status     int     `json:"status"`                            // 商品状态：0-全部 1-上架 2-下架
}

type Goods struct {
	Total      int64       `json:"total"`       // 总数量
	TotalPages int         `json:"total_pages"` // 总页数
	List       []GoodsItem `json:"list"`        // 商品列表
}

// 商品信息
type GoodsItem struct {
	GoodsID     int64      `json:"goods_id"`       // 商品ID
	Name        string     `json:"name"`           // 商品名称
	Title       string     `json:"title"`          // 商品标题
	Description string     `json:"description"`    // 商品描述
	MainImage   string     `json:"main_image"`     // 主图URL
	Images      []string   `json:"images"`         // 商品图集
	CategoryID  int64      `json:"category_id"`    // 分类ID
	Category    string     `json:"category"`       // 分类名称
	BrandID     int64      `json:"brand_id"`       // 品牌ID
	Brand       string     `json:"brand"`          // 品牌名称
	Price       float64    `json:"price"`          // 商品价格
	MarketPrice float64    `json:"market_price"`   // 市场价
	Stock       int        `json:"stock"`          // 库存
	Sales       int        `json:"sales"`          // 销量
	Status      int        `json:"status"`         // 状态：0-下架 1-上架
	CreateTime  string     `json:"create_time"`    // 创建时间
	UpdateTime  string     `json:"update_time"`    // 更新时间
	SKUs        []GoodsSKU `json:"skus,omitempty"` // SKU列表
	Tags        []string   `json:"tags"`           // 标签
}

// 商品SKU
type GoodsSKU struct {
	SkuID      int64   `json:"sku_id"`     // SKU ID
	Attributes []Attr  `json:"attributes"` // 属性列表
	Price      float64 `json:"price"`      // SKU价格
	Stock      int     `json:"stock"`      // SKU库存
	Image      string  `json:"image"`      // SKU图片
	Code       string  `json:"code"`       // SKU编码
}

// 属性键值对
type Attr struct {
	Key   string `json:"key"`   // 属性名：如"颜色"
	Value string `json:"value"` // 属性值：如"红色"
}

type AccountResp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []Account `json:"data"`
}

type OrderResp struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data []Order `json:"data"`
}

type GoodsResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Goods  `json:"data"`
}
