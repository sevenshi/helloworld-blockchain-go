package vo

/*
 @author king 409060350@qq.com
*/

import "helloworld-blockchain-go/dto"

type PayerVo struct {
	PrivateKey             string `json:"privateKey"`
	TransactionHash        string `json:"transactionHash"`
	TransactionOutputIndex uint64 `json:"transactionOutputIndex"`
	Value                  uint64 `json:"value"`
	Address                string `json:"address"`
}
type PayeeVo struct {
	Address string `json:"address"`
	Value   uint64 `json:"value"`
}

type AutomaticBuildTransactionRequest struct {
	NonChangePayees []*PayeeVo `json:"nonChangePayees"`
}
type AutomaticBuildTransactionResponse struct {
	BuildTransactionSuccess bool                `json:"buildTransactionSuccess"`
	Message                 string              `json:"message"`
	TransactionHash         string              `json:"transactionHash"`
	Fee                     uint64              `json:"fee"`
	Payers                  []*PayerVo          `json:"payers"`
	NonChangePayees         []*PayeeVo          `json:"nonChangePayees"`
	ChangePayee             *PayeeVo            `json:"changePayee"`
	Payees                  []*PayeeVo          `json:"payees"`
	Transaction             *dto.TransactionDto `json:"transaction"`
}
