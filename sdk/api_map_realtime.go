package sdk

import (
	"encoding/json"
	"github.com/pilafusama/kdniaoGo"
	"github.com/pilafusama/kdniaoGo/enum"
	"github.com/pilafusama/kdniaoGo/request"
	"github.com/pilafusama/kdniaoGo/response"
	"github.com/pilafusama/kdniaoGo/util"
	"github.com/pilafusama/kdniaoGo/util/http"
	"strconv"
)

func NewApiMapRealtime(config kdniaoGo.KdniaoConfig, logger kdniaoGo.KdniaoLoggerInterface) ApiMapRealtime {
	return ApiMapRealtime{config, logger}
}

type ApiMapRealtime struct {
	config kdniaoGo.KdniaoConfig
	logger kdniaoGo.KdniaoLoggerInterface
}

func (obj ApiMapRealtime) GetRequest(logisticCode, shipperCode string) request.MapRealtimeRequest {
	req := request.NewMapRealtimeRequest()
	req.SetConfig(obj.config)
	req.SetLogisticCode(logisticCode)
	req.SetShipperCode(shipperCode)
	return req
}

func (obj ApiMapRealtime) GetResponse(req request.MapRealtimeRequest) (response.MapRealtimeResponse, error) {
	url := enum.GATEWAY + enum.URI_BUSINESS

	req.UpdateRequestData()
	var resp response.MapRealtimeResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMapRealtime,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiMapRealtime,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMapRealtime,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiMapRealtime," + resp.GetError())
	}
	return resp, nil
}
