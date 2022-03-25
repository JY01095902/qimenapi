package singleitem

import (
	"fmt"

	"github.com/jy01095902/topapi/request"
)

type SynchronizeRequest struct {
	BaseURL   string
	AppKey    string
	OwnerCode string // 货主编码
	Data      struct {
		ActionType    string `json:"actionType"`
		WarehouseCode string `json:"warehouseCode"`
		OwnerCode     string `json:"ownerCode"`
		Item          struct {
			ItemCode string `json:"itemCode"`
			ItemType string `json:"itemType"`
			ItemId   string `json:"itemId"`
			ItemName string `json:"itemName"`
			BarCode  string `json:"barCode"`
		} `json:"item"`
	}
}

type SynchronizeResponse struct {
	Flag    string `json:"flag"`
	Code    string `json:"code"`
	Message string `json:"message"`
	ItemId  string `json:"itemId"`
}

func Synchronize(req SynchronizeRequest) (SynchronizeResponse, error) {
	apiReq := request.New(request.Config{
		BaseURL:            req.BaseURL,
		AppKey:             req.AppKey,
		AppSecret:          "",
		MsgType:            "taobao.qimen.singleitem.synchronize",
		LogisticProviderId: req.OwnerCode,
		ToCode:             "",
		PartnerCode:        "",
		FromCode:           "",
	})
	vals, err := apiReq.Post(req.Data)
	if err != nil {
		return SynchronizeResponse{}, err
	}

	result, err := vals.GetResult(SynchronizeResponse{})
	if err != nil {
		return SynchronizeResponse{}, fmt.Errorf("%w: %s", request.ErrCallQiMenAPIFailed, err.Error())
	}

	resp, ok := result.(*SynchronizeResponse)
	if !ok {
		return SynchronizeResponse{}, fmt.Errorf("%w: result is not SynchronizeResponse", request.ErrCallQiMenAPIFailed)
	}

	if resp.Flag == "failure" {
		return *resp, fmt.Errorf("%w: (%s)%s", request.ErrQiMenAPIBizError, resp.Code, resp.Message)
	}

	return *resp, nil
}
