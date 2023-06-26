package response

type MapSubscribeResponse struct {
	Success               bool   `json:"Success"`
	Reason                string `json:"Reason"`
	EBusinessId           string `json:"EBusinessID"`
	UpdateTime            string `json:"UpdateTime"`
	EstimatedDeliveryTime string `json:"EstimatedDeliveryTime"`
	SenderCityLatAndLng   string `json:"SenderCityLatAndLng"`
	ReceiverCityLatAndLng string `json:"ReceiverCityLatAndLng"`
	RouteMapUrl           string `json:"RouteMapUrl"`
	Location              string `json:"Location"`
	LogisticCode          string `json:"LogisticCode"`
	ShipperCode           string `json:"ShipperCode"`
}

func (resp *MapSubscribeResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *MapSubscribeResponse) GetError() string {
	return resp.Reason
}
