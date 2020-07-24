package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNameDosNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")
)
