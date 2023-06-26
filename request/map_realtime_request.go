package request

import (
	"encoding/json"

	"github.com/pilafusama/kdniaoGo/enum"
)

func NewMapRealtimeRequest() MapRealtimeRequest {
	req := MapRealtimeRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_MAP_REALTIME)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type MapRealtimeRequest struct {
	KdniaoRequest
	LogisticCode        string     `json:"LogisticCode"`           // 物流单号
	ShipperCode         string     `json:"ShipperCode"`            // 快递公司编码
	OrderCode           string     `json:"OrderCode,omitempty"`    // 订单编号
	CustomerName        string     `json:"CustomerName,omitempty"` // 顺丰必填
	Sender              AddrDetail `json:"Sender,omitempty"`       // 为提升地图准确度，建议传入收发双方省市区
	Receiver            AddrDetail `json:"Receiver,omitempty"`     // 为提升地图准确度，建议传入收发双方省市区
	IsReturnCoordinates int        `json:"IsReturnCoordinates"`    // 是否返回城市经纬度 1：返回，2：不返回，默认1
	IsReturnRouteMap    int        `json:"IsReturnRouteMap"`       // 是否返回轨迹地图 1：返回，2：不返回，默认2
}

type AddrDetail struct {
	ProvinceName string
	CityName     string
	ExpAreaName  string
	Address      string
}

func (req *MapRealtimeRequest) SetLogisticCode(logisticCode string) *MapRealtimeRequest {
	req.LogisticCode = logisticCode
	return req
}

func (req MapRealtimeRequest) GetLogisticCode() string {
	return req.LogisticCode
}

func (req *MapRealtimeRequest) SetShipperCode(shipperCode string) *MapRealtimeRequest {
	req.ShipperCode = shipperCode
	return req
}

func (req MapRealtimeRequest) GetShipperCode() string {
	return req.ShipperCode
}

func (req *MapRealtimeRequest) SetOrderCode(orderCode string) *MapRealtimeRequest {
	req.OrderCode = orderCode
	return req
}

func (req MapRealtimeRequest) GetOrderCode() string {
	return req.OrderCode
}

func (req *MapRealtimeRequest) UpdateRequestData() *MapRealtimeRequest {
	req.requestData = req.ToJson()
	return req
}

func (req MapRealtimeRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}

func (req *MapRealtimeRequest) SetCustomerName(customerName string) *MapRealtimeRequest {
	req.CustomerName = customerName
	return req
}
