package resource

import (
	"github.com/digitalbitsorg/go/amount"
	"github.com/digitalbitsorg/go/price"
	"github.com/digitalbitsorg/go/services/frontier/internal/db2/history"
	"golang.org/x/net/context"
)

// Populate fills out the details of a trade using a row from the history_trades
// table.
func (res *TradeAggregation) Populate(
	ctx context.Context,
	row history.TradeAggregation,
) (err error) {
	res.Timestamp = row.Timestamp
	res.TradeCount = row.TradeCount
	res.BaseVolume = amount.StringFromInt64(row.BaseVolume)
	res.CounterVolume = amount.StringFromInt64(row.CounterVolume)
	res.Average = price.StringFromFloat64(row.Average)
	res.High = row.High.String()
	res.HighR = row.High
	res.Low = row.Low.String()
	res.LowR = row.Low
	res.Open = row.Open.String()
	res.OpenR = row.Open
	res.Close = row.Close.String()
	res.CloseR = row.Close
	return
}

// PagingToken implementation for hal.Pageable. Not actually used
func (res TradeAggregation) PagingToken() string {
	return string(res.Timestamp)
}
