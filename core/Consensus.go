package core

/*
 @author king 409060350@qq.com
*/

import (
	"fmt"
	"helloworld-blockchain-go/core/model"
	"helloworld-blockchain-go/core/tool/BlockTool"
	"helloworld-blockchain-go/crypto/ByteUtil"
	"helloworld-blockchain-go/setting/GenesisBlockSetting"
	"helloworld-blockchain-go/setting/IncentiveSetting"
	"helloworld-blockchain-go/util/StringUtil"
	"math/big"
)

type Consensus struct {
}

func NewConsensus() *Consensus {
	var consensus Consensus
	return &consensus
}

func (c *Consensus) CheckConsensus(blockchainDatabase *BlockchainDatabase, block *model.Block) bool {
	difficulty := block.Difficulty
	if StringUtil.IsNullOrEmpty(difficulty) {
		difficulty = c.CalculateDifficult(blockchainDatabase, block)
	}

	hash := block.Hash
	if StringUtil.IsNullOrEmpty(hash) {
		hash = BlockTool.CalculateBlockHash(block)
	}

	bigIntDifficulty := new(big.Int).SetBytes(ByteUtil.HexStringToBytes(difficulty))
	bigIntHash := new(big.Int).SetBytes(ByteUtil.HexStringToBytes(hash))
	return bigIntDifficulty.Cmp(bigIntHash) > 0
}

func (c *Consensus) CalculateDifficult(blockchainDatabase *BlockchainDatabase, targetBlock *model.Block) string {

	targetDifficult := ""
	targetBlockHeight := targetBlock.Height
	if targetBlockHeight <= IncentiveSetting.INTERVAL_BLOCK_COUNT*uint64(2) {
		targetDifficult = GenesisBlockSetting.DIFFICULTY
		return targetDifficult
	}

	targetBlockPreviousBlock := blockchainDatabase.QueryBlockByBlockHeight(targetBlockHeight - uint64(1))
	if targetBlockPreviousBlock.Height%IncentiveSetting.INTERVAL_BLOCK_COUNT != 0 {
		targetDifficult = targetBlockPreviousBlock.Difficulty
		return targetDifficult
	}

	previousIntervalLastBlock := targetBlockPreviousBlock
	previousPreviousIntervalLastBlock := blockchainDatabase.QueryBlockByBlockHeight(previousIntervalLastBlock.Height - IncentiveSetting.INTERVAL_BLOCK_COUNT)
	previousIntervalActualTimespan := previousIntervalLastBlock.Timestamp - previousPreviousIntervalLastBlock.Timestamp

	fmt.Println(previousIntervalActualTimespan)

	bigIntPreviousIntervalDifficulty := new(big.Int).SetBytes(ByteUtil.HexStringToBytes(previousIntervalLastBlock.Difficulty))
	bigIntPreviousIntervalActualTimespan := new(big.Int).SetUint64(previousIntervalActualTimespan)
	bigIntIntervalTime := new(big.Int).SetUint64(IncentiveSetting.INTERVAL_TIME)

	bigIntegerMul := new(big.Int).Mul(bigIntPreviousIntervalDifficulty, bigIntPreviousIntervalActualTimespan)
	bigIntegerTargetDifficult := new(big.Int).Div(bigIntegerMul, bigIntIntervalTime)
	return ByteUtil.BytesToHexString(bigIntegerTargetDifficult.Bytes())
}
