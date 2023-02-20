package main

import "log"

/*
 * Single Responsibility Principle (S)
 * "One class should have one, and only one, responsibility."
 */

// There are three common accounting transactions types are Cash, Non-Cash & Credit.
// I use one common (Sales) of four financial transaction types (which are Sales,Purchases,Receipts,Payments)
// to build a class
type Financial interface {
	GetSellingDate() string
	CalculateSalesFee() float64
}

/*
 * Open/Closed Principle (O)
 * "You should be able to extend a class's behavior without modifying it."
 */

type Sales struct {
	financials []Financial
}

func (s *Sales) CalculateSalesFee() float64 {
	total := 0.0
	for _, f := range s.financials {
		total += f.CalculateSalesFee()
	}
	return total
}

func (s *Sales) AddSales(f Financial) {
	s.financials = append(s.financials, f)
}

/*
 * Interface Segregation Principle (I)
 * "Many client-specific interfaces are better than one general-purpose interface."
 */
type InsideSales interface {
	Financial
}

type OutsideSales interface {
	Financial
}

type InsideSalesImpl struct {
	sellingDate string
}

func (i InsideSalesImpl) GetSellingDate() string {
	return i.sellingDate
}

func (i InsideSalesImpl) CalculateSalesFee() float64 {
	return 1.0
}

type OutsideSalesImpl struct {
	sellingDate string
}

func (o OutsideSalesImpl) GetSellingDate() string {
	return o.sellingDate
}

func (o OutsideSalesImpl) CalculateSalesFee() float64 {
	return 2.0
}

/*
 * Builder
 * The builder pattern allows you to create different flavors of an object
 * while enforcing the constraints mentioned previously
 */

type FinancialBuilder interface {
	Types(string) FinancialBuilder
	SellingDate(string) FinancialBuilder
	Build() Financial
}

type financialBuilder struct {
	types string
	sdate string
}

func (fb *financialBuilder) Types(t string) FinancialBuilder {
	fb.types = t
	return fb
}

func (fb *financialBuilder) SellingDate(date string) FinancialBuilder {
	fb.sdate = date
	return fb
}

// two of eight common sales types (inside,outside,b2b,b2c,businessdevelopment,agency,consultative,ecommerce)
func (fb *financialBuilder) Build() Financial {
	var builtFinancial Financial

	switch fb.types {
	case "inside":
		return InsideSalesImpl{fb.sdate}
	case "outside":
		return OutsideSalesImpl{fb.sdate}
	}

	return builtFinancial
}

func NewFinancialBuilder() FinancialBuilder {
	return &financialBuilder{}
}

func main() {
	fi := NewFinancialBuilder()
	fi.Types("inside")
	fi.SellingDate("20230601")
	fi.Build()
	log.Printf("%+v", fi)

	fo := NewFinancialBuilder()
	fo.Types("outside")
	fo.SellingDate("20230602")
	fo.Build()
	log.Printf("%+v", fo)
}
