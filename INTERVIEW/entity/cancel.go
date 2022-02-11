package entity

type CancelRequest struct {
	TimeStamp        string `json:"timeStamp"`
	TransactionID    string `json:"tXid"`
	MerchantID       string `json:"iMid"`
	PayMethod        string `json:"payMethod"`
	CancelType       string `json:"cancelType"`
	MerchantToken    string `json:"merchantToken"`
	Amount           string `json:"amt"`
	CancelMessage    string `json:"cancelMsg"`
	PreauthToken     string `json:"preauthToken"`
	ReferenceNo      string `json:"referenceNo"`
	Fee              string `json:"fee"`
	Vat              string `json:"vat"`
	NotaxAmt         string `json:"notaxAmt"`
	CancelServerIP   string `json:"cancelServerIp"`
	CancelUserID     string `json:"cancelUserId"`
	CancelUserIP     string `json:"cancelUserIp"`
	CancelUserInfo   string `json:"cancelUserInfo"`
	CancelRetryCount string `json:"cancelRetryCnt"`
	Worker           string `json:"worker"`
}

type CancelResponse struct {
	ResultCode              string `json:"resultCd"`
	ResultMessage           string `json:"resultMsg"`
	TransactionID           string `json:"tXid"`
	ReferenceNo             string `json:"referenceNo"`
	TransactionDate         string `json:"transDt"`
	TransactionTime         string `json:"transTm"`
	Description             string `json:"description"`
	Amount                  string `json:"amt"`
	CancelTransactionId     string `json:"cancelTxId"`
	CancelTransactionNumber string `json:"cancelTrxSn"`
	CancelReferenceNo       string `json:"cancelReferenceNo"`
}
