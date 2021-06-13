package BlockTool

import (
	"helloworldcoin-go/core/Model/TransactionType"
	"helloworldcoin-go/core/tool/BlockDtoTool"
	"helloworldcoin-go/core/tool/Model2DtoTool"
	"helloworldcoin-go/core/tool/TransactionTool"

	"helloworldcoin-go/core/Model"
)

func CalculateBlockHash(block Model.Block) string {
	blockDto := Model2DtoTool.Block2BlockDto(&block)
	return BlockDtoTool.CalculateBlockHash(blockDto)
}

func CalculateBlockMerkleTreeRoot(block Model.Block) string {
	blockDto := Model2DtoTool.Block2BlockDto(&block)
	return BlockDtoTool.CalculateBlockMerkleTreeRoot(blockDto)
}

func GetTransactionCount(block *Model.Block) uint64 {
	transactions := block.Transactions
	return uint64(len(transactions))
}
func GetTransactionOutputCount(block *Model.Block) uint64 {
	transactionOutputCount := uint64(0)
	transactions := block.Transactions
	if transactions != nil {
		for _, transaction := range transactions {
			transactionOutputCount = transactionOutputCount + TransactionTool.GetTransactionOutputCount(&transaction)
		}
	}
	return transactionOutputCount
}
func GetBlockFee(block *Model.Block) uint64 {
	blockFee := uint64(0)
	transactions := block.Transactions
	if transactions != nil {
		for _, transaction := range transactions {
			if transaction.TransactionType == TransactionType.GENESIS_TRANSACTION {
				continue
			} else if transaction.TransactionType == TransactionType.STANDARD_TRANSACTION {
				fee := TransactionTool.GetTransactionFee(&transaction)
				blockFee += fee
			} else {
			}
		}
	}
	return blockFee
}
func GetWritedIncentiveValue(block *Model.Block) uint64 {
	return block.Transactions[0].Outputs[0].Value
}
