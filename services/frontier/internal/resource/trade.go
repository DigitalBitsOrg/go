package resource

import (
	"fmt"

	"github.com/digitalbitsorg/go/amount"
	"github.com/digitalbitsorg/go/services/frontier/internal/db2/history"
	"github.com/digitalbitsorg/go/services/frontier/internal/httpx"
	"github.com/digitalbitsorg/go/services/frontier/internal/render/hal"
	"golang.org/x/net/context"
	"github.com/digitalbitsorg/go/xdr"
)

// Populate fills out the details of a trade using a row from the history_trades
// table.
func (res *Trade) Populate(
	ctx context.Context,
	row history.Trade,
) (err error) {
	res.ID = row.PagingToken()
	res.PT = row.PagingToken()
	res.OfferID = fmt.Sprintf("%d", row.OfferID)
	res.BaseAccount = row.BaseAccount
	res.BaseAssetType = row.BaseAssetType
	res.BaseAssetCode = row.BaseAssetCode
	res.BaseAssetIssuer = row.BaseAssetIssuer
	res.BaseAmount = amount.String(row.BaseAmount)
	res.CounterAccount = row.CounterAccount
	res.CounterAssetType = row.CounterAssetType
	res.CounterAssetCode = row.CounterAssetCode
	res.CounterAssetIssuer = row.CounterAssetIssuer
	res.CounterAmount = amount.String(row.CounterAmount)
	res.LedgerCloseTime = row.LedgerCloseTime
	res.BaseIsSeller = row.BaseIsSeller
	res.Price = xdr.Price{row.PriceN, row.PriceD}
	res.populateLinks(ctx, row.HistoryOperationID)
	return
}

// PagingToken implementation for hal.Pageable
func (res Trade) PagingToken() string {
	return res.PT
}

func (res *Trade) populateLinks(
	ctx context.Context,
	opid int64,
) {
	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	res.Links.Base = lb.Link("/accounts", res.BaseAccount)
	res.Links.Counter = lb.Link("/accounts", res.CounterAccount)
	res.Links.Operation = lb.Link(
		"/operations",
		fmt.Sprintf("%d", opid),
	)
}
