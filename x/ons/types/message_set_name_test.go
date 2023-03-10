package types

import (
	"testing"

	"github.com/ollo-station/ollo/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSetName_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetName
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetName{
				CreatorAddr: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetName{
				CreatorAddr: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
