package triangular

import (
	"github.com/nxdht/decimal"
	"testing"
)

func TestTriangular_CheckForwardResult(t *testing.T) {
	triangular := Triangular{
		SymbolA:    "BETA",
		SymbolB:    "BNB",
		SymbolC:    "USDT",
		FeeAB:      decimal0,
		FeeAC:      decimal0,
		FeeBC:      decimal0,
		SlippageAB: decimal0,
		SlippageAC: decimal0,
		SlippageBC: decimal0,
		BidAB:      decimal.NewFromFloat(0.0053104),
		BidAC:      decimal.NewFromFloat(2.48002),
		BidBC:      decimal.NewFromFloat(461.5),
		AskAB:      decimal.NewFromFloat(0.0053442),
		AskAC:      decimal.NewFromFloat(2.48199),
		AskBC:      decimal.NewFromFloat(461.7),
		Result:     Result{},
	}

	triangular.CheckForwardResult()

	t.Log(triangular.Result)
}
