package resource

import (
	"github.com/digitalbitsorg/go/services/frontier/internal/db2/core"
	"golang.org/x/net/context"
)

// Populate fills out the fields of the signer, using one of an account's
// secondary signers.
func (this *Signer) Populate(ctx context.Context, row core.Signer) {
	this.PublicKey = row.Publickey
	this.Weight = row.Weight
	this.Key = row.Publickey
	this.Type = MustKeyTypeFromAddress(this.PublicKey)
}

// PopulateMaster fills out the fields of the signer, using a digitalbits account to
// provide the data.
func (this *Signer) PopulateMaster(row core.Account) {
	this.PublicKey = row.Accountid
	this.Weight = int32(row.Thresholds[0])
	this.Key = row.Accountid
	this.Type = MustKeyTypeFromAddress(this.PublicKey)
}
