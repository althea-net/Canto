package keeper

// DONTCOVER

import (
	"github.com/Canto-Network/Canto/v2/x/csr/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/csr module sentinel errors
var (
	ErrPrevRegisteredSmartContract = sdkerrors.Register(types.ModuleName, 2000, "csr::CSR")
	ErrFeeCollectorDistribution    = sdkerrors.Register(types.ModuleName, 2001, "csr::EVMHOOK")
	ErrCSRNFTNotDeployed           = sdkerrors.Register(types.ModuleName, 2002, "csr::Keeper")
	ErrAddressDerivation           = sdkerrors.Register(types.ModuleName, 2003, "csr::Keeper")
	ErrContractDeployments         = sdkerrors.Register(types.ModuleName, 2004, "csr::Keeper")
	ErrMethodCall                  = sdkerrors.Register(types.ModuleName, 2005, "csr::Keeper")
	ErrUnpackData                  = sdkerrors.Register(types.ModuleName, 2006, "csr:Keeper")
)
