package request

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

func timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

type body map[string]string

func newBody(config Config, data interface{}) body {
	b := body{
		"msg_type":             config.MsgType,
		"logistic_provider_id": config.LogisticProviderId,
		"partner_code":         config.PartnerCode,
		"from_code":            config.FromCode,
		"msg_id":               strconv.FormatInt(timestamp(time.Now()), 10),
		"to_code":              config.ToCode,
	}

	if bs, err := json.Marshal(data); err == nil {
		d := string(bs)
		b["logistics_interface"] = d
		b["data_digest"] = sign(d, config.AppKey)
	}

	return b
}

func (b body) EncodeToString() string {
	str := ""
	for k, v := range b {
		str += k + "=" + v
	}

	return url.QueryEscape(str)
}

type Config struct {
	BaseURL            string
	AppKey             string
	AppSecret          string
	MsgType            string
	LogisticProviderId string // 货主编码
	ToCode             string
	PartnerCode        string
	FromCode           string
}

type Request struct {
	config Config
}

func New(config Config) Request {
	req := Request{
		config: config,
	}

	return req
}

func (req Request) execute(r *resty.Request, method, url string) (Values, error) {
	resp, err := r.
		SetResult(Values{}).
		Execute(method, url)

	if err != nil {
		return nil, fmt.Errorf("%w error: %s", ErrCallQiMenAPIFailed, err.Error())
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("%w error: %s", ErrCallQiMenAPIFailed, resp.String())
	}

	result, ok := resp.Result().(*Values)
	if !ok {
		return Values{}, fmt.Errorf("%w: type of result is not Values", ErrCallQiMenAPIFailed)
	}

	return *result, err
}

func (req Request) Post(data interface{}) (Values, error) {
	b := newBody(req.config, data)

	r := resty.New().R().
		EnableTrace().
		SetBody(b.EncodeToString())

	return req.execute(r, resty.MethodPost, req.config.BaseURL)
}
