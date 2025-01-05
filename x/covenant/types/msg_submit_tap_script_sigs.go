package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
)

var _ sdk.Msg = &SubmitTapScriptSigsRequest{}

// NewSubmitTapScriptSigRequest constructor for SubmitTapScriptSigRequest
func NewSubmitTapScriptSigsRequest(sender sdk.AccAddress, sigID uint64, tapScriptSigs *exported.TapScriptSigList) *SubmitTapScriptSigsRequest {
	return &SubmitTapScriptSigsRequest{
		Sender:        sender,
		SigID:         sigID,
		TapScriptSigs: tapScriptSigs,
	}
}

// ValidateBasic implements the sdk.Msg interface.
func (m SubmitTapScriptSigsRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	for _, tapScriptSig := range m.TapScriptSigs.TapScriptSigs {
		if err := tapScriptSig.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	return nil
}

// GetSigners implements the sdk.Msg interface
func (m SubmitTapScriptSigsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
