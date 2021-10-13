package triangular

import "github.com/nxdht/decimal"



var (
	decimal1 = decimal.NewFromInt(1)
	decimal0 = decimal.NewFromInt(0)
)

type Triangular struct {
	// Pairs: A/B, A/C, B/C
	SymbolA, SymbolB, SymbolC string
	FeeAB, FeeAC, FeeBC decimal.Decimal
	SlippageAB, SlippageAC, SlippageBC decimal.Decimal

	BidAB, BidAC, BidBC decimal.Decimal
	AskAB, AskAC, AskBC decimal.Decimal

	Result Result
}

type Result struct {
	Success bool
	Forward bool

	BuyAB, BuyAC, BuyBC bool
	PriceAB, PriceAC, PriceBC decimal.Decimal
	QuantityAB, QuantityAC, QuantityBC decimal.Decimal
}

func (t *Triangular) CheckForwardResult()  {
	// buy A x 1 need B:
	buyAneedB := t.AskAB.Mul(decimal1.Add(t.SlippageAB))
	// sell A x 1 get C:
	sellAgetC := t.BidAC.Mul(decimal1.Sub(t.SlippageAC))
	// buy B x 1 need C:
	buyBneedC := t.AskBC.Mul(decimal1.Add(t.SlippageBC))

	// buy A x 1 need B actually:
	buyAneedBActually := buyAneedB.Div(decimal1.Sub(t.FeeAB))
	// sell A x 1 get C actually:
	sellAgetCActually := sellAgetC.Mul(decimal1.Sub(t.FeeAC))
	// buy B x buyAneedBActually need buy B actually: buyAneedBActually.Div(decimal1.Sub(t.FeeBC))
	// buy B x buyAneedBActually.Div(decimal1.Sub(t.FeeBC)) need C actually:
	buyBneedCActually := buyBneedC.Mul(buyAneedBActually).Div(decimal1.Sub(t.FeeBC))

	if sellAgetCActually.GreaterThan(buyBneedCActually) {
		t.Result.Success = true
		t.Result.Forward = true

		t.Result.BuyAB = true
		t.Result.PriceAB = buyAneedB
		t.Result.QuantityAB = decimal1
		t.Result.BuyAC = false
		t.Result.PriceAC = sellAgetC
		t.Result.QuantityAC = decimal1
		t.Result.BuyBC = true
		t.Result.PriceBC = buyBneedC
		t.Result.QuantityBC = buyAneedBActually.Div(decimal1.Sub(t.FeeBC))
	} else {
		t.Result.Success = false
	}
}

func (t *Triangular) CheckReverseResult() {
	//sellAgetB := t.BidAB.Mul(decimal1.Sub(t.SlippageAB))
	//buyAneedC := t.AskAC.Mul(decimal1.Add(t.SlippageAC))
	//sellBgetC := t.BidBC.Mul(decimal1.Sub(t.SlippageBC))
	//
	//sellAgetB = sellAgetB.Mul(decimal1.Sub(t.FeeAB))
	//buyAneedC = buyAneedC.Div(decimal1.Sub(t.FeeAC))
}
