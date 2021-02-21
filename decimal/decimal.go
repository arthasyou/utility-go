package decimal

import "github.com/shopspring/decimal"

// Add string
func Add(a string, b string) string {
	c, _ := decimal.NewFromString(a)
	d, _ := decimal.NewFromString(b)
	r := c.Add(d)
	return r.String()
}

// Sub string
func Sub(a string, b string) string {
	c, _ := decimal.NewFromString(a)
	d, _ := decimal.NewFromString(b)
	r := c.Sub(d)
	return r.String()
}

// Mul string
func Mul(a string, b string) string {
	c, _ := decimal.NewFromString(a)
	d, _ := decimal.NewFromString(b)
	r := c.Mul(d)
	return r.String()
}

// Div string
func Div(a string, b string) string {
	c, _ := decimal.NewFromString(a)
	d, _ := decimal.NewFromString(b)
	r := c.Div(d)
	return r.String()
}

// Abs string
func Abs(a string) string {
	c, _ := decimal.NewFromString(a)
	r := c.Abs()
	return r.String()
}
