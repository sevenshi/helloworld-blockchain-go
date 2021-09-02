package controller

import (
	"helloworld-blockchain-go/application/service"
	"helloworld-blockchain-go/application/vo/block"
	"helloworld-blockchain-go/application/vo/framwork"
	"helloworld-blockchain-go/application/vo/node"
	"helloworld-blockchain-go/application/vo/transaction"
	"helloworld-blockchain-go/core/Model"
	"helloworld-blockchain-go/core/tool/BlockTool"
	"helloworld-blockchain-go/core/tool/TransactionDtoTool"
	"helloworld-blockchain-go/netcore"
	"helloworld-blockchain-go/setting/GenesisBlockSetting"
	"helloworld-blockchain-go/util/JsonUtil"
	"helloworld-blockchain-go/util/StringUtil"
	"helloworld-blockchain-go/util/TimeUtil"
	"io"
	"io/ioutil"
	"net/http"
)

type BlockchainBrowserApplicationController struct {
	blockchainNetCore                   *netcore.BlockchainNetCore
	blockchainBrowserApplicationService *service.BlockchainBrowserApplicationService
}

func NewBlockchainBrowserApplicationController(blockchainNetCore *netcore.BlockchainNetCore) *BlockchainBrowserApplicationController {
	var b BlockchainBrowserApplicationController
	b.blockchainNetCore = blockchainNetCore
	return &b
}

func (b *BlockchainBrowserApplicationController) QueryTransactionByTransactionHash(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), transaction.QueryTransactionByTransactionHashRequest{}).(*transaction.QueryTransactionByTransactionHashRequest)

	transactionVo := b.blockchainBrowserApplicationService.QueryTransactionByTransactionHash(request.TransactionHash)
	if transactionVo == nil {
		//return Response.createFailResponse("根据交易哈希未能查询到交易")
	}

	var response transaction.QueryTransactionByTransactionHashResponse
	response.Transaction = transactionVo

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}
func (b *BlockchainBrowserApplicationController) QueryTransactionsByBlockHashTransactionHeight(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), transaction.QueryTransactionsByBlockHashTransactionHeightRequest{}).(*transaction.QueryTransactionsByBlockHashTransactionHeightRequest)

	pageCondition := request.PageCondition
	if StringUtil.IsNullOrEmpty(request.BlockHash) {
		//return Response.createFailResponse("区块哈希不能是空");
	}
	transactionVos := b.blockchainBrowserApplicationService.QueryTransactionListByBlockHashTransactionHeight(request.BlockHash, pageCondition.From, pageCondition.Size)
	var response transaction.QueryTransactionsByBlockHashTransactionHeightResponse
	response.Transactions = transactionVos

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}
func (b *BlockchainBrowserApplicationController) QueryTransactionOutputByAddress(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), transaction.QueryTransactionOutputByAddressRequest{}).(*transaction.QueryTransactionOutputByAddressRequest)

	transactionOutputDetailVo := b.blockchainBrowserApplicationService.QueryTransactionOutputByAddress(request.Address)
	var response transaction.QueryTransactionOutputByAddressResponse
	response.TransactionOutputDetail = transactionOutputDetailVo

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}
func (b *BlockchainBrowserApplicationController) QueryTransactionOutputByTransactionOutputId(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), transaction.QueryTransactionOutputByTransactionOutputIdRequest{}).(*transaction.QueryTransactionOutputByTransactionOutputIdRequest)

	transactionOutputDetailVo := b.blockchainBrowserApplicationService.QueryTransactionOutputByTransactionOutputId(request.TransactionHash, request.TransactionOutputIndex)
	var response transaction.QueryTransactionOutputByTransactionOutputIdResponse
	response.TransactionOutputDetail = transactionOutputDetailVo

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}
func (b *BlockchainBrowserApplicationController) QueryBlockchainHeight(w http.ResponseWriter, req *http.Request) {

	blockchainHeight := b.blockchainNetCore.GetBlockchainCore().QueryBlockchainHeight()
	var response node.QueryBlockchainHeightResponse
	response.BlockchainHeight = blockchainHeight
	s := framwork.CreateSuccessResponse("", response)

	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}

func (b *BlockchainBrowserApplicationController) QueryUnconfirmedTransactionByTransactionHash(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), transaction.QueryUnconfirmedTransactionByTransactionHashRequest{}).(*transaction.QueryUnconfirmedTransactionByTransactionHashRequest)

	unconfirmedTransactionVo := b.blockchainBrowserApplicationService.QueryUnconfirmedTransactionByTransactionHash(request.TransactionHash)
	if unconfirmedTransactionVo == nil {
		//return Response.createFailResponse("交易哈希["+request.getTransactionHash()+"]不是未确认交易。");
	}
	var response transaction.QueryUnconfirmedTransactionByTransactionHashResponse
	response.Transaction = unconfirmedTransactionVo

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}

func (b *BlockchainBrowserApplicationController) QueryUnconfirmedTransactions(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), transaction.QueryUnconfirmedTransactionsRequest{}).(*transaction.QueryUnconfirmedTransactionsRequest)

	pageCondition := request.PageCondition
	transactionDtos := b.blockchainNetCore.GetBlockchainCore().QueryUnconfirmedTransactions(pageCondition.From, pageCondition.Size)
	if transactionDtos == nil {
		//return Response.createSuccessResponse("未查询到未确认的交易");
	}

	var unconfirmedTransactionVos []*transaction.UnconfirmedTransactionVo
	for _, transactionDto := range transactionDtos {
		unconfirmedTransactionVo := b.blockchainBrowserApplicationService.QueryUnconfirmedTransactionByTransactionHash(TransactionDtoTool.CalculateTransactionHash(transactionDto))
		if unconfirmedTransactionVo != nil {
			unconfirmedTransactionVos = append(unconfirmedTransactionVos, unconfirmedTransactionVo)
		}
	}
	var response transaction.QueryUnconfirmedTransactionsResponse
	response.UnconfirmedTransactions = unconfirmedTransactionVos

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}

func (b *BlockchainBrowserApplicationController) QueryBlockByBlockHeight(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), block.QueryBlockByBlockHeightRequest{}).(*block.QueryBlockByBlockHeightRequest)

	blockVo := b.blockchainBrowserApplicationService.QueryBlockViewByBlockHeight(request.BlockHeight)
	if blockVo == nil {
		//return Response.createFailResponse("区块链中不存在区块高度["+request.getBlockHeight()+"]，请检查输入高度。");
	}
	var response block.QueryBlockByBlockHeightResponse
	response.Block = blockVo

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}

func (b *BlockchainBrowserApplicationController) QueryBlockByBlockHash(w http.ResponseWriter, req *http.Request) {
	result, _ := ioutil.ReadAll(req.Body)
	request := JsonUtil.ToObject(string(result), block.QueryBlockByBlockHashRequest{}).(*block.QueryBlockByBlockHashRequest)

	block1 := b.blockchainNetCore.GetBlockchainCore().QueryBlockByBlockHash(request.BlockHash)
	if block1 == nil {
		//return Response.createFailResponse("区块链中不存在区块哈希["+request.getBlockHash()+"]，请检查输入哈希。");
	}
	blockVo := b.blockchainBrowserApplicationService.QueryBlockViewByBlockHeight(block1.Height)
	var response block.QueryBlockByBlockHashResponse
	response.Block = blockVo

	s := framwork.CreateSuccessResponse("", response)
	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}

func (b *BlockchainBrowserApplicationController) QueryTop10Blocks(w http.ResponseWriter, req *http.Request) {

	var blocks []*Model.Block
	blockHeight := b.blockchainNetCore.GetBlockchainCore().QueryBlockchainHeight()
	for {
		if blockHeight <= GenesisBlockSetting.HEIGHT {
			break
		}
		block := b.blockchainNetCore.GetBlockchainCore().QueryBlockByBlockHeight(blockHeight)
		blocks = append(blocks, block)
		if len(blocks) >= 10 {
			break
		}
		blockHeight--
	}

	var blockVos []block.BlockVo2
	for _, block1 := range blocks {
		var blockVo block.BlockVo2
		blockVo.Height = block1.Height
		blockVo.BlockSize = "100字符" //TODO SizeTool.CalculateBlockSize(block1) + "字符" TODO
		blockVo.TransactionCount = BlockTool.GetTransactionCount(block1)
		blockVo.MinerIncentiveValue = BlockTool.GetWritedIncentiveValue(block1)
		blockVo.Time = TimeUtil.FormatMillisecondTimestamp(block1.Timestamp)
		blockVo.Hash = block1.Hash
		blockVos = append(blockVos, blockVo)
	}

	var response block.QueryTop10BlocksResponse
	response.Blocks = blockVos
	s := framwork.CreateSuccessResponse("", response)

	w.Header().Set("content-type", "text/json")
	io.WriteString(w, s)
}
