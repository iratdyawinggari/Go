package entity

import "time"

type StatusRequest struct {
	TimeStamp     string `json:"timeStamp"`
	TransactionID string `json:"tXid"`
	MerchantID    string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo"`
	Amount        string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}

type StatusResponse struct {
	ResultCode             string    `json:"resultCd"`
	ResultMessage          string    `json:"resultMsg"`
	TransactionID          string    `json:"tXid"`
	MerchantID             string    `json:"iMid"`
	ReferenceNo            string    `json:"referenceNo"`
	PayMethod              string    `json:"payMethod"`
	Amount                 string    `json:"amt"`
	CancelAmount           int       `json:"cancelAmt"`
	RequestDate            string    `json:"reqDt"`
	RequestTime            string    `json:"reqTm"`
	TransactionDate        time.Time `json:"transDt"`
	TransactionTime        time.Time `json:"transTm"`
	TransactionDepositDate time.Time `json:"depositDt"`
	TransactionDepositTime time.Time `json:"depositTm"`
	Currency               string    `json:"currency"`
	GoodsName              string    `json:"goodsNm"`
	BillingName            string    `json:"billingNm"`
	Status                 string    `json:"status"`

	BankCd       string `json:"bankCd"`
	VacctNo      string `json:"vacctNo"`
	VacctValidDt string `json:"vacctValidDt"`
	VacctValidTm string `json:"vacctValidTm"`

	MitraCode   string `json:"mitraCd"`
	PayNo       string `json:"payNo"`
	PayValidDt  string `json:"payValidDt"`
	PayValidTm  string `json:"payValidTm"`
	ReceiptCode string `json:"receiptCode"`
	MRefNo      string `json:"mRefNo"`
}
