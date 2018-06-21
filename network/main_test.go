package network

import (
	"testing"

	"github.com/digitalbitsorg/go/xdr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashTransaction(t *testing.T) {
	var txe xdr.TransactionEnvelope

	err := xdr.SafeUnmarshalBase64("AAAAANQ3vRdz7vwQ7O3sz8O9esdqAJjBO2LhK+MhfOT5i0+iAAAAZAAAAskAAAAgAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAx/kTxNoYXcAhqH2VUuSaw4dBiYAAVOUzrG0R60fV8K8AAAAXSHboAAAAAAAAAAAB+YtPogAAAEDoW1YfsTwPz0vrAF7bkvL2+sA47x1REqSI2OwvEfM0J7mWGNmCeS/1vh/bGWuAV+Wh482FRy3Oxo2MqI8ITVwI", &txe)

	require.NoError(t, err)

	expected := [32]byte{
		0x7c, 0xd4, 0x75, 0x39, 0xfe, 0xab, 0xd7, 0x1a, 
		0x8f, 0x1b, 0x1d, 0x26, 0x88, 0x8, 	0x7c, 0x67, 
		0xad, 0x64, 0xf8, 0xf0, 0x90, 0xcb, 0x51, 0x80, 
		0x17, 0xd5, 0x37, 0x2, 	0x6f, 0x74, 0x5,  0xd,
	}

	actual, err := HashTransaction(&txe.Tx, TestNetworkPassphrase)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

	// sadpath: empty passphrase
	_, err = HashTransaction(&txe.Tx, "")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "empty network passphrase")
	}
}
