package rpc

import (
	"github.com/HydroProtocol/ethereum-watcher/blockchain"
	"time"
)

type EthBlockChainMultiRPCWithRetry struct {
	rpcList       []*EthBlockChainRPC
	maxRetryTimes int
}

func NewEthMultiRPCWithRetry(api []string, maxRetryCount int) *EthBlockChainMultiRPCWithRetry {
	rpcList := make([]*EthBlockChainRPC, 0)
	for _, a := range api {
		rpcList = append(rpcList, NewEthRPC(a))
	}

	return &EthBlockChainMultiRPCWithRetry{rpcList, maxRetryCount}
}

func (rpc EthBlockChainMultiRPCWithRetry) GetBlockByNum(num uint64) (rst blockchain.Block, err error) {
	for _, r := range rpc.rpcList {
		for i := 0; i <= rpc.maxRetryTimes; i++ {
			rst, err = r.GetBlockByNum(num)
			if err == nil {
				break
			} else {
				time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			}
		}
	}

	return
}

func (rpc EthBlockChainMultiRPCWithRetry) GetLiteBlockByNum(num uint64) (rst blockchain.Block, err error) {
	for _, r := range rpc.rpcList {
		for i := 0; i <= rpc.maxRetryTimes; i++ {
			rst, err = r.GetLiteBlockByNum(num)
			if err == nil {
				break
			} else {
				time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			}
		}
	}

	return
}

func (rpc EthBlockChainMultiRPCWithRetry) GetTransactionReceipt(txHash string) (rst blockchain.TransactionReceipt, err error) {
	for _, r := range rpc.rpcList {
		for i := 0; i <= rpc.maxRetryTimes; i++ {
			rst, err = r.GetTransactionReceipt(txHash)
			if err == nil {
				break
			} else {
				time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			}
		}

	}

	return
}

func (rpc EthBlockChainMultiRPCWithRetry) GetCurrentBlockNum() (rst uint64, err error) {
	for _, r := range rpc.rpcList {
		for i := 0; i <= rpc.maxRetryTimes; i++ {
			rst, err = r.GetCurrentBlockNum()
			if err == nil {
				break
			} else {
				time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			}
		}
	}

	return
}
func (rpc EthBlockChainMultiRPCWithRetry) GetLogs(
	fromBlockNum, toBlockNum uint64,
	address string,
	topics []string,
) (rst []blockchain.IReceiptLog, err error) {
	for _, r := range rpc.rpcList {
		for i := 0; i <= rpc.maxRetryTimes; i++ {
			rst, err = r.GetLogs(fromBlockNum, toBlockNum, address, topics)
			if err == nil {
				break
			} else {
				time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			}
		}
	}
	return
}
