package keeper

import (
	"github.com/koalanetwork/koala/x/ibcdex/types"
)

var _ types.QueryServer = Keeper{}
