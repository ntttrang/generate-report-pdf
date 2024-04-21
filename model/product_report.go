package model

type (
	ProductInfo struct {
		Id          string
		Name        string
		Description string
	}
	Inventory struct {
		Location  string
		Quantity  int64
		Available int64
	}
	Sale struct {
		Date         string
		QuantitySold int64
		TotalRevenue float64
	}
)
