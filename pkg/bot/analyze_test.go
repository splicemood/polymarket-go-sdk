package bot

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/splicemood/polymarket-go-sdk/v2/pkg/clob/clobtypes"
)

func TestTopOfBook(t *testing.T) {
	price, depth, err := topOfBook([]clobtypes.PriceLevel{{Price: "0.53", Size: "220.12"}})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !price.Equal(decimal.RequireFromString("0.53")) {
		t.Fatalf("unexpected price: %s", price)
	}
	if !depth.Equal(decimal.RequireFromString("220.12")) {
		t.Fatalf("unexpected depth: %s", depth)
	}
}
