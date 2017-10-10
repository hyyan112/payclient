package client

import (
	"encoding/xml"
	"fmt"
	"testing"
	"github.com/fatih/structs"
	"encoding/json"
)

func TestAppendSign(t *testing.T) {
	requestXML := WXPayUnifiedOrderRequest{}
	tmp := struct {
		WXPayUnifiedOrderRequest
		XMLName struct{}    `xml:"xml"`
	}{WXPayUnifiedOrderRequest: requestXML}


	s, _ := xml.Marshal(tmp)
	fmt.Println(string(s))
	fmt.Println(appendSign(tmp))
	d,_:=json.Marshal(structs.Map(tmp))
	fmt.Println(string(d))
}
