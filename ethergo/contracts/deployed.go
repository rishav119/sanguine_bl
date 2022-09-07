package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// DeployedContract is the contract interface.
type DeployedContract interface {
	// Address is the address where the contract has been deployed
	Address() common.Address
	// ContractHandle is the actual handle returned by deploying the contract
	// this must be castt o be useful
	ContractHandle() interface{}
	// Owner of the contract
	Owner() common.Address
	// DeployTx is the transaction where the contract was created
	DeployTx() *types.Transaction
	// ChainID is the chain id
	ChainID() *big.Int
	// OwnerPtr is a pointer to the owner
	OwnerPtr() *common.Address
}