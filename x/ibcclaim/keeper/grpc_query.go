package keeper

import (
	"github.com/koalanetwork/koala/x/ibcclaim/types"
)

var _ types.QueryServer = Keeper{}
