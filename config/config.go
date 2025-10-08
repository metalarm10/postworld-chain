package config

import (
	"github.com/metalarm10/postworld-chain/types"
	evmtypes "github.com/metalarm10/postworld-chain/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ChainsCoinInfo is a map of the chain id and its corresponding EvmCoinInfo
// that allows initializing the app with different coin info based on the
// chain id
var ChainsCoinInfo = map[uint64]evmtypes.EvmCoinInfo{ // TODO:VLAD - Remove this
	EighteenDecimalsChainID: {
		Denom:         ExampleChainDenom,
		ExtendedDenom: ExampleChainDenom,
		DisplayDenom:  ExampleDisplayDenom,
		Decimals:      evmtypes.EighteenDecimals,
	},
	// SixDecimalsChainID provides a chain ID which is being set up with 6 decimals
	SixDecimalsChainID: {
		Denom:         "usurvive",
		ExtendedDenom: "asurvive",
		DisplayDenom:  "survive",
		Decimals:      evmtypes.SixDecimals,
	},
	// EVMChainID provides a chain ID used for internal testing
	EVMChainID: {
		Denom:         "asurvive",
		ExtendedDenom: "asurvive",
		DisplayDenom:  "survive",
		Decimals:      evmtypes.EighteenDecimals,
	},
	TwelveDecimalsChainID: {
		Denom:         "psurvive2",
		ExtendedDenom: "asurvive2",
		DisplayDenom:  "survive2",
		Decimals:      evmtypes.TwelveDecimals,
	},
	TwoDecimalsChainID: {
		Denom:         "csurvive3",
		ExtendedDenom: "asurvive3",
		DisplayDenom:  "survive3",
		Decimals:      evmtypes.TwoDecimals,
	},
	TestChainID1: {
		Denom:         ExampleChainDenom,
		ExtendedDenom: ExampleChainDenom,
		DisplayDenom:  ExampleChainDenom,
		Decimals:      evmtypes.EighteenDecimals,
	},
	TestChainID2: {
		Denom:         ExampleChainDenom,
		ExtendedDenom: ExampleChainDenom,
		DisplayDenom:  ExampleChainDenom,
		Decimals:      evmtypes.EighteenDecimals,
	},
}

const (
	// Bech32Prefix defines the Bech32 prefix used for accounts on the Postworld Chain.
	Bech32Prefix = "postworld"
	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address.
	Bech32PrefixAccAddr = Bech32Prefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key.
	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address.
	Bech32PrefixValAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key.
	Bech32PrefixValPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address.
	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key.
	Bech32PrefixConsPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = "survive"
	// BaseDenom defines to the default denomination used in Postworld Chain.
	BaseDenom = "asurvive"
	// BaseDenomUnit defines the precision of the base denomination.
	BaseDenomUnit = 18
	// EVMChainID defines the EIP-155 replay-protection chain id for the current ethereum chain config.
	EVMChainID = 9000
)

// SetBech32Prefixes sets the global prefixes to be used when serializing addresses and public keys to Bech32 strings.
func SetBech32Prefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

// SetBip44CoinType sets the global coin type to be used in hierarchical deterministic wallets.
func SetBip44CoinType(config *sdk.Config) {
	config.SetCoinType(types.Bip44CoinType)
	config.SetPurpose(sdk.Purpose)                  // Shared
	config.SetFullFundraiserPath(types.BIP44HDPath) //nolint: staticcheck
}
