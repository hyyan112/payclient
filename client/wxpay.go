package client

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"payclient/utils"
	"strings"

	"github.com/fatih/structs"
)

type (
	WXPay struct {
		AppId          string
		MchId          string
		Key            string
		ClientCertPath string
		ClientKeyPath  string
		CaCertsPath    string
	}

	CommonWXPayRequest struct {
		AppId    string `xml:"appid",json:"appid"`
		MchId    string `xml:"mch_id"`
		NonceStr string `xml:"nonce_str"`
		Sign     string `xml:"sign"`
		SignType string `xml:"sign_type"`
	}

	CommonWXPayResponse struct {
	}

	WXPayUnifiedOrderRequest struct {
		CommonWXPayRequest
		Body           string `xml:"body"`
		OutTradeNo     string `xml:"out_trade_no"`
		TotalFee       int    `xml:"total_fee"`
		SpbillCreateIp string `xml:"spbill_create_ip"`
		NotifyUrl      string `xml:"notify_url"`
		TradeType      string `xml:"trade_type"`
	}

	WXPayUnifiedOrderResponse struct {
	}
)

func sendRequest(req *http.Request, resp *interface{}) error {
	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(resBody, resp)
	if err != nil {
		return err
	}
	return nil
}

func appendSign(requestXML interface{}) string {
	dataMap := structs.Map(requestXML)
	data := []string{}
	for k, v := range dataMap {
		data = append(data, fmt.Sprintf("%v=%v", k, v))
	}
	return strings.Join(data,"&")
}

func (wxpay *WXPay) UnifiedOrder(requestXML *WXPayUnifiedOrderRequest, resp *interface{}) error {
	requestXML.AppId = wxpay.AppId
	requestXML.MchId = wxpay.MchId
	requestXML.NonceStr = utils.RandString(32)

	body, err := xml.Marshal(requestXML)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("https://api.mch.weixin.qq.com/pay/unifiedorder", "POST", strings.NewReader(string(body)))
	if err != nil {
		return nil
	}
	return sendRequest(request, resp)
}
