package order

import (
	"fmt"

	"github.com/jy01095902/qimenapi/request"
)

type CallbackRequest struct {
	BaseURL            string
	AppKey             string
	LogisticProviderId string
	ToCode             string
	Data               struct {
		WarehouseCode string `json:"warehouseCode"`
		OrderId       string `json:"orderId"`
	}
}

type CallbackResponse struct {
	Flag    string `json:"flag"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// 配送拦截接口
func Callback(req CallbackRequest) (CallbackResponse, error) {
	apiReq := request.New(request.Config{
		BaseURL:            req.BaseURL,
		AppKey:             req.AppKey,
		AppSecret:          "",
		MsgType:            "taobao.qimen.order.callback",
		LogisticProviderId: req.LogisticProviderId,
		ToCode:             req.ToCode,
		PartnerCode:        "",
		FromCode:           "",
	})

	vals, err := apiReq.Post(req.Data)
	if err != nil {
		return CallbackResponse{}, err
	}

	result, err := vals.GetResult(CallbackResponse{})
	if err != nil {
		return CallbackResponse{}, fmt.Errorf("%w: %s", request.ErrCallQiMenAPIFailed, err.Error())
	}

	resp, ok := result.(*CallbackResponse)
	if !ok {
		return CallbackResponse{}, fmt.Errorf("%w: result is not CallbackResponse", request.ErrCallQiMenAPIFailed)
	}

	if resp.Flag == "failure" {
		return *resp, fmt.Errorf("%w: (%s)%s", request.ErrQiMenAPIBizError, resp.Code, resp.Message)
	}

	return *resp, nil
}
