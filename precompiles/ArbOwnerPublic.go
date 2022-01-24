//
// Copyright 2021, Offchain Labs, Inc. All rights reserved.
//

package precompiles

import (
	"github.com/ethereum/go-ethereum/common"
)

// The calls to this precompile do not require the sender be a chain owner.
// For those that are, see ArbOwner
type ArbOwnerPublic struct {
	Address addr
}

func (con ArbOwnerPublic) GetAllChainOwners(c ctx, evm mech) ([]common.Address, error) {
	return c.state.ChainOwners().AllMembers()
}

func (con ArbOwnerPublic) IsChainOwner(c ctx, evm mech, addr addr) (bool, error) {
	return c.state.ChainOwners().IsMember(addr)
}

func (con ArbOwnerPublic) GetNetworkFeeAccount(c ctx, evm mech) (addr, error) {
	return c.state.NetworkFeeAccount()
}