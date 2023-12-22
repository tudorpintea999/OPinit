package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
	groupcodec "github.com/cosmos/cosmos-sdk/x/group/codec"
)

// RegisterLegacyAminoCodec registers the move types and interface
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgRecordBatch{}, "ophost/MsgRecordBatch")
	legacy.RegisterAminoMsg(cdc, &MsgCreateBridge{}, "ophost/MsgCreateBridge")
	legacy.RegisterAminoMsg(cdc, &MsgProposeOutput{}, "ophost/MsgProposeOutput")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteOutput{}, "ophost/MsgDeleteOutput")
	legacy.RegisterAminoMsg(cdc, &MsgInitiateTokenDeposit{}, "ophost/MsgInitiateTokenDeposit")
	legacy.RegisterAminoMsg(cdc, &MsgFinalizeTokenWithdrawal{}, "ophost/MsgFinalizeTokenWithdrawal")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateProposer{}, "ophost/MsgUpdateProposer")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateChallenger{}, "ophost/MsgUpdateChallenger")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "ophost/MsgUpdateParams")

	cdc.RegisterConcrete(Params{}, "ophost/Params", nil)
	cdc.RegisterConcrete(&BridgeAccount{}, "ophost/BridgeAccount", nil)
}

// RegisterInterfaces registers the x/market interfaces types with the interface registry
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRecordBatch{},
		&MsgCreateBridge{},
		&MsgProposeOutput{},
		&MsgDeleteOutput{},
		&MsgInitiateTokenDeposit{},
		&MsgFinalizeTokenWithdrawal{},
		&MsgUpdateProposer{},
		&MsgUpdateChallenger{},
		&MsgUpdateParams{},
	)

	// auth account registration
	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&BridgeAccount{},
	)
	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&BridgeAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz  and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}