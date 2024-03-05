package request

import (
	"encoding/json"

	"github.com/pilafusama/kdniaoGo/enum"
)

func NewMapSubscribeRequest() MapSubscribeRequest {
	req := MapSubscribeRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_MAP_SUBSCRIBE)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type MapSubscribeRequest struct {
	KdniaoRequest
	CallBack            string     `json:"Callback,omitempty"`     // 用户自定义回调信息
	CallbackUrl         string     `json:"CallbackUrl,omitempty"`  // 自定义回调地址
	MemberId            string     `json:"MemberID,omitempty"`     // 会员标识
	WareHouseId         string     `json:"WareHouseID,omitempty"`  // 仓库标识
	CustomerName        string     `json:"CustomerName,omitempty"` // 电子面单客户号
	CustomerPwd         string     `json:"CustomerPwd,omitempty"`  // 电子面单密码
	SendSite            string     `json:"SendSite,omitempty"`     // 收件网点标识(名称)
	ShipperCode         string     `json:"ShipperCode"`            // 快递公司编码
	LogisticCode        string     `json:"LogisticCode"`           // 快递单号
	OrderCode           string     `json:"OrderCode,omitempty"`    // 订单编号
	MonthCode           string     `json:"MonthCode,omitempty"`    // 月结编码
	PayType             string     `json:"PayType,omitempty"`      // 邮费支付方式: 1-现付，2-到付，3-月结，4-第三方支付)
	ExpType             string     `json:"ExpType,omitempty"`      // 快递类型：1-标准快件
	Cost                string     `json:"Cost,omitempty"`         // 快递运费
	OtherCost           string     `json:"OtherCost,omitempty"`    // 快递运费
	Receiver            AddrDetail `json:"Receiver,omitempty"`
	Sender              AddrDetail `json:"Sender,omitempty"`
	IsNotice            string     `json:"IsNotice,omitempty"`  // 是否通知快递员上门揽件：0-通知；1-不通知；不填则
	StartDate           string     `json:"StartDate,omitempty"` // 上门揽件时间段，格式：YYYY-MM-DD HH24:MM:SS
	EndDate             string     `json:"EndDate,omitempty"`
	Weight              string     `json:"Weight,omitempty"`              // 包裹总重量kg
	Quantity            string     `json:"Quantity,omitempty"`            // 包裹数，一个包裹对应一个运单号，如果是大于1个包裹，返回则按照子母件的方式返回母运单号和子运单号
	Volume              string     `json:"Volume,omitempty"`              // 包裹总体积m3
	Remark              string     `json:"Remark,omitempty"`              // 备注
	IsSendMessage       string     `json:"IsSendMessage,omitempty"`       // 是否订阅短信：0-不需要；1-需要
	IsReturnCoordinates int        `json:"IsReturnCoordinates,omitempty"` // 是否返回城市经纬度 1：返回，2：不返回，不填则默认为1
	IsReturnRouteMap    int        `json:"IsReturnRouteMap,omitempty"`    // 是否返回轨迹地图 1：返回，2：不返回，不填则默认为2

	//AddService    MapSubscribeRequestAddService  `json:"AddService,omitempty"`
	//Commodity     []MapSubscribeRequestCommodity `json:"Commodity,omitempty"`
}

type MapSubscribeRequestCommodity struct {
	GoodsName     string `json:"GoodsName"`     // 商品名称
	GoodsCode     string `json:"GoodsCode"`     // 商品编码
	Goodsquantity string `json:"Goodsquantity"` // 商品件数
	GoodsPrice    string `json:"GoodsPrice"`    // 商品价格
	GoodsWeight   string `json:"GoodsWeight"`   // 商品重量kg
	GoodsDesc     string `json:"GoodsDesc"`     // 商品描述
	GoodsVol      string `json:"GoodsVol"`      // 商品体积m3
}

func (req *MapSubscribeRequest) SetLogisticCode(logisticCode string) *MapSubscribeRequest {
	req.LogisticCode = logisticCode
	return req
}

func (req MapSubscribeRequest) GetLogisticCode() string {
	return req.LogisticCode
}

func (req *MapSubscribeRequest) SetShipperCode(shipperCode string) *MapSubscribeRequest {
	req.ShipperCode = shipperCode
	return req
}

func (req MapSubscribeRequest) GetShipperCode() string {
	return req.ShipperCode
}

func (req *MapSubscribeRequest) SetCustomerName(customerName string) *MapSubscribeRequest {
	req.CustomerName = customerName
	return req
}

func (req MapSubscribeRequest) GetCustomerName() string {
	return req.CustomerName
}

func (req *MapSubscribeRequest) UpdateRequestData() *MapSubscribeRequest {
	req.requestData = req.ToJson()
	return req
}

func (req MapSubscribeRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}
