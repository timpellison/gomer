package gomer

import (
	"errors"
)

type Gomer struct {
	Revenue          float64
	MarketingExpense float64
}

func (g *Gomer) Calculate() (float64, error) {
	if g.MarketingExpense == 0 {
		return 0, errors.New("cannot calculate MER when marketing expense is 0")
	}
	return g.Revenue / g.MarketingExpense, nil
}

func MustNewGomer(revenue float64, marketingExpense float64) *Gomer {
	v := &Gomer{Revenue: revenue, MarketingExpense: marketingExpense}
	return v
}
