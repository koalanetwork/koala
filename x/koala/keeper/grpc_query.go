package keeper

import (
	"github.com/koalanetwork/koala/x/koala/types"
)

var _ types.QueryServer = Keeper{}
