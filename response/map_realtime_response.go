package response

//type MapRealtimeResponse struct {
//	Success      bool                       `json:"Success"`
//	Reason       string                     `json:"Reason"`
//	EBusinessId  string                     `json:"EBusinessID"`
//	OrderCode    string                     `json:"OrderCode"`    // 订单编号
//	ShipperCode  string                     `json:"ShipperCode"`  // 快递公司编码
//	LogisticCode string                     `json:"LogisticCode"` // 物流运单号
//	State        string                     `json:"State"`        // 物流状态：2-在途中,3-签收,4-问题件
//	StateEx      string                     `json:"StateEx"`      // 增值物流状态： 1-已揽收， 2-在途中， 201-到达派件城市， 202-派件中， 211-已放入快递柜或驿站， 3-已签收， 311-已取出快递柜或驿站， 4-问题件， 401-发货无信息， 402-超时未签收， 403-超时未更新， 404-拒收（退件）， 412-快递柜或驿站超时未取
//	Location     string                     `json:"Location"`     // 增值所在城市
//	Traces       []MapRealtimeResponseTrace `json:"Traces"`
//}

type MapRealtimeResponseTrace struct {
	AcceptTime    string `json:"AcceptTime"`    // 时间
	AcceptStation string `json:"AcceptStation"` // 描述
	Location      string `json:"Location"`      // 当前城市
	Action        string `json:"Action"`        // 当前状态
	Remark        string `json:"Remark"`        // 备注
}

func (resp *MapRealtimeResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *MapRealtimeResponse) GetError() string {
	return resp.Reason
}

// MapRealtimeResponse https://www.yuque.com/kdnjishuzhichi/dfcrg1/ed5y64
type MapRealtimeResponse struct {
	EBusinessID           string                     `json:"EBusinessID"`
	ShipperCode           string                     `json:"ShipperCode"`
	Success               bool                       `json:"Success"`
	LogisticCode          string                     `json:"LogisticCode"`
	SenderCityLatAndLng   string                     `json:"SenderCityLatAndLng"`
	ReceiverCityLatAndLng string                     `json:"ReceiverCityLatAndLng"`
	State                 string                     `json:"State"`
	CityIsNull            string                     `json:"CityIsNull"`
	StateEx               string                     `json:"StateEx"`
	Location              string                     `json:"Location"`
	Traces                []MapRealtimeResponseTrace `json:"Traces"`
	Coordinates           []struct {                 // 当前城市经纬度
		Location  string `json:"Location"`
		LatAndLng string `json:"LatAndLng"`
	} `json:"Coordinates"`
	PreCoordinates []struct { // 预设路径经过城市
		Location  string `json:"Location"`
		LatAndLng string `json:"LatAndLng"`
	} `json:"PreCoordinates"`
	EstimatedDeliveryTime string `json:"EstimatedDeliveryTime"` // 订单预计到达时间
	RouteMapUrl           string `json:"RouteMapUrl"`           // 轨迹地图URL
	RouteMapUrlCode       string `json:"RouteMapUrlCode"`
	Reason                string `json:"Reason"`
}
