package keeper

import "github.com/lum-network/chain/x/millions/types"

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
