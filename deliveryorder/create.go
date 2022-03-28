package deliveryorder

import (
	"fmt"

	"github.com/jy01095902/qimenapi/request"
)

type Contact struct {
	Name          string `json:"name"`
	Mobile        string `json:"mobile"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Area          string `json:"area"`
	Town          string `json:"town"`
	DetailAddress string `json:"detailAddress"`
}

type DeliveryOrder struct {
	DeliveryOrderCode    string  `json:"deliveryOrderCode"`
	PreDeliveryOrderCode string  `json:"preDeliveryOrderCode"`
	OrderType            string  `json:"orderType"` // 出库单类型，JYCK=一般交易出库单, HHCK=换货出库单, BFCK=补发出库单，QTCK=其他出库单
	WarehouseCode        string  `json:"warehouseCode"`
	SourcePlatformCode   string  `json:"sourcePlatformCode"`
	CreateTime           string  `json:"createTime"`     // 发货单创建时间 2015-06-12 20:26:32
	PlaceOrderTime       string  `json:"placeOrderTime"` // 前台订单 (店铺订单) 创建时间 (下单时间) 2015-06-12 20:26:32
	OperateTime          string  `json:"operateTime"`
	ShopNick             string  `json:"shopNick"`
	LogisticsCode        string  `json:"logisticsCode"` // 物流公司编码
	SenderInfo           Contact `json:"senderInfo"`
	ReceiverInfo         Contact `json:"receiverInfo"`
}

type OrderLine struct {
	OrderLineNo string `json:"orderLineNo"`
	OwnerCode   string `json:"ownerCode"`
	ItemCode    string `json:"itemCode"`
	ItemId      string `json:"itemId"`
	PlanQty     string `json:"planQty"`
}

type CreateRequest struct {
	BaseURL     string
	AppKey      string
	FromCode    string
	PartnerCode string
	Data        struct {
		DeliveryOrder DeliveryOrder `json:"deliveryOrder"`
		OrderLines    []OrderLine   `json:"orderLines"`
	}
}

type CreateResponse struct {
	Flag            string `json:"flag"`
	Code            string `json:"code"`
	Message         string `json:"message"`
	DeliveryOrderId string `json:"deliveryOrderId"`
	WarehouseCode   string `json:"warehouseCode"`
	LogisticsCode   string `json:"logisticsCode"`
	DeliveryOrder   []struct {
		DeliveryOrderId string `json:"deliveryOrderId"`
		WarehouseCode   string `json:"warehouseCode"`
		LogisticsCode   string `json:"logisticsCode"`
		OrderLines      []struct {
			OrderLineNo string `json:"orderLineNo"`
			ItemCode    string `json:"itemCode"`
			ItemId      string `json:"itemId"`
			Quantity    string `json:"quantity"`
		} `json:"orderLines"`
	} `json:"deliveryOrders"`
}

func Create(req CreateRequest) (CreateResponse, error) {
	apiReq := request.New(request.Config{
		BaseURL:            req.BaseURL,
		AppKey:             req.AppKey,
		AppSecret:          "",
		MsgType:            "taobao.qimen.deliveryorder.create",
		LogisticProviderId: "",
		ToCode:             "",
		PartnerCode:        req.PartnerCode,
		FromCode:           req.FromCode,
	})

	vals, err := apiReq.Post(req.Data)
	if err != nil {
		return CreateResponse{}, err
	}

	result, err := vals.GetResult(CreateResponse{})
	if err != nil {
		return CreateResponse{}, fmt.Errorf("%w: %s", request.ErrCallQiMenAPIFailed, err.Error())
	}

	resp, ok := result.(*CreateResponse)
	if !ok {
		return CreateResponse{}, fmt.Errorf("%w: result is not CreateResponse", request.ErrCallQiMenAPIFailed)
	}

	if resp.Flag == "failure" {
		return *resp, fmt.Errorf("%w: (%s)%s", request.ErrQiMenAPIBizError, resp.Code, resp.Message)
	}

	return *resp, nil
}
