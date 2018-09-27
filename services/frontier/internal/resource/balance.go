package resource

import (
	"github.com/digitalbitsorg/go/amount"
	"github.com/digitalbitsorg/go/services/frontier/internal/assets"
	"github.com/digitalbitsorg/go/services/frontier/internal/db2/core"
	"github.com/digitalbitsorg/go/xdr"
	"golang.org/x/net/context"
)

func (this *Balance) Populate(ctx context.Context, row core.Trustline) (err error) {
	this.Type, err = assets.String(row.Assettype)
	if err != nil {
		return
	}

	this.Balance = amount.String(row.Balance)
	this.Limit = amount.String(row.Tlimit)
	this.Issuer = row.Issuer
	this.Code = row.Assetcode
	return
}

func (this *Balance) PopulateNative(stroops xdr.Int64) (err error) {
	this.Type, err = assets.String(xdr.AssetTypeAssetTypeNative)
	if err != nil {
		return
	}

	this.Balance = amount.String(stroops)
	this.Limit = ""
	this.Issuer = ""
	this.Code = ""
	return
}
