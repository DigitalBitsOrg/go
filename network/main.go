// Package network contains functions that deal with digitalbits network passphrases
// and IDs.
package network

import (
	"bytes"
	"fmt"

	"strings"

	"github.com/digitalbitsorg/go/hash"
	"github.com/digitalbitsorg/go/support/errors"
	"github.com/digitalbitsorg/go/xdr"
)

const (
	// PublicNetworkPassphrase is the pass phrase used for every transaction intended for the public digitalbits network
	PublicNetworkPassphrase = "Live DigitalBits Network ; March 2018"
	// TestNetworkPassphrase is the pass phrase used for every transaction intended for the SDF-run test network
	TestNetworkPassphrase = "Test DigitalBits Network ; December 2017"
)

// ID returns the network ID derived from the provided passphrase.  This value
// also happens to be the raw (i.e. not strkey encoded) secret key for the root
// account of the network.
func ID(passphrase string) [32]byte {
	return hash.Hash([]byte(passphrase))
}

// HashTransaction derives the network specific hash for the provided
// transaction using the network identified by the supplied passphrase.  The
// resulting hash is the value that can be signed by digitalbits secret key to
// authorize the transaction identified by the hash to digitalbits validators.
func HashTransaction(tx *xdr.Transaction, passphrase string) ([32]byte, error) {
	var txBytes bytes.Buffer

	if strings.TrimSpace(passphrase) == "" {
		return [32]byte{}, errors.New("empty network passphrase")
	}

	_, err := fmt.Fprintf(&txBytes, "%s", ID(passphrase))
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "fprint network id failed")
	}

	_, err = xdr.Marshal(&txBytes, xdr.EnvelopeTypeEnvelopeTypeTx)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "marshal type failed")
	}

	_, err = xdr.Marshal(&txBytes, tx)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "marshal tx failed")
	}

	return hash.Hash(txBytes.Bytes()), nil
}
