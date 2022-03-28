package deliveryorder

import (
	"fmt"

	"github.com/jy01095902/qimenapi/request"
)

type CancelRequest struct {
	BaseURL     string
	AppKey      string
	FromCode    string
	PartnerCode string
	Data        struct {
		WarehouseCode string `json:"warehouseCode"`
		OrderCode     string `json:"orderCode"`
		OrderId       string `json:"orderId"`
	}
}

type CancelResponse struct {
	Flag    string `json:"flag"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Cancel(req CancelRequest) (CancelResponse, error) {
	apiReq := request.New(request.Config{
		BaseURL:            req.BaseURL,
		AppKey:             req.AppKey,
		AppSecret:          "",
		MsgType:            "taobao.qimen.order.cancel",
		LogisticProviderId: "",
		ToCode:             "",
		PartnerCode:        req.PartnerCode,
		FromCode:           req.FromCode,
	})

	vals, err := apiReq.Post(req.Data)
	if err != nil {
		return CancelResponse{}, err
	}

	result, err := vals.GetResult(CancelResponse{})
	if err != nil {
		return CancelResponse{}, fmt.Errorf("%w: %s", request.ErrCallQiMenAPIFailed, err.Error())
	}

	resp, ok := result.(*CancelResponse)
	if !ok {
		return CancelResponse{}, fmt.Errorf("%w: result is not CancelResponse", request.ErrCallQiMenAPIFailed)
	}

	if resp.Flag == "failure" {
		return *resp, fmt.Errorf("%w: (%s)%s", request.ErrQiMenAPIBizError, resp.Code, resp.Message)
	}

	return *resp, nil
}
