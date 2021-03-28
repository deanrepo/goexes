package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type QooAppPTDReceipts struct {
	Total int                 `json:"total"`
	Data  []*QooAppPTDReceipt `json:"data"`
}
type QooAppPTDReceipt struct {
	OrderId  string     `json:"orderId"`
	AppId    string     `json:"appId"`
	Status   string     `json:"status"`
	PayTime  *time.Time `json:"payTime"`
	ClientId string     `json:"clientId"`
	Amount   string     `json:"amount"`
	Currency string     `json:"currency"`
	Country  string     `json:"country"`
}

func main() {
	testJ := `{
		"total": 2,
		"data": [
			{
				"orderId": "7eedb79356311a5ad3a642b8ad85bc77",
				"appId": "SAND01aae732253811eaacd306c11e2c3ade",
				"status": "SUCCESS",
				"payTime": "2019-12-30T16:50:50Z",
				"clientId": "sX9Db_cLEbSAi1Y8UX_CFQ",
				"amount": "200",
				"currency": "TWD",
				"country": ""
			},
			{
				"orderId": "8eedb79356311a5ad3a642b8ad85bc78",
				"appId": "SAND01aae732253811eaacd306c11e2c3ade",
				"status": "SUCCESS",
				"payTime": "2019-12-30T16:55:50Z",
				"clientId": "sX9Db_cLEbSAi1Y8UX_CFQ",
				"amount": "200",
				"currency": "TWD",
				"country": ""
			}
		]
	}`

	var t QooAppPTDReceipts
	err := json.Unmarshal([]byte(testJ), &t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
	fmt.Println(t.Data[0])
	fmt.Println(t.Data[1])

}
