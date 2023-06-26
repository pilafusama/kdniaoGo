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

func NewApiMapSubscribe(config kdniaoGo.KdniaoConfig, logger kdniaoGo.KdniaoLoggerInterface) ApiMapSubscribe {
	return ApiMapSubscribe{config, logger}
}

type ApiMapSubscribe struct {
	config kdniaoGo.KdniaoConfig
	logger kdniaoGo.KdniaoLoggerInterface
}

func (obj ApiMapSubscribe) GetRequest(logisticCode, shipperCode string) request.MapSubscribeRequest {
	req := request.NewMapSubscribeRequest()
	req.SetConfig(obj.config)
	req.SetLogisticCode(logisticCode)
	req.SetShipperCode(shipperCode)
	return req
}

func (obj ApiMapSubscribe) GetResponse(req request.MapSubscribeRequest) (response.MapSubscribeResponse, error) {
	url := enum.GATEWAY + enum.URI_API

	req.UpdateRequestData()
	var resp response.MapSubscribeResponse
	httpResp, err := http.HttpPostForm(url, req.ToValues(), obj.logger)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMapSubscribe,http fail")
	} else if !httpResp.IsOk() {
		return resp, util.ErrorNew("ApiMapSubscribe,code:" + strconv.Itoa(httpResp.GetCode()))
	}
	err = json.Unmarshal(httpResp.GetBytes(), &resp)
	if err != nil {
		return resp, util.ErrorWrap(err, "ApiMapSubscribe,json decode fail")
	}
	if !resp.IsSuccess() {
		return resp, util.ErrorNew("ApiMapSubscribe," + resp.GetError())
	}
	return resp, nil
}
