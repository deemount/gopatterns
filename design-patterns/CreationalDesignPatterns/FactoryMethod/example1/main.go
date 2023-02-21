package financial

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
 * Factory Method
 * A factory is an object that is used to create other objects.
 * In a factory method pattern, a helper method (or function) is defined,
 * to enable object creation without knowing the implementation class details.
 */

// One of eight common sales types (inside,outside,b2b,b2c,businessdevelopment,agency,consultative,ecommerce)
func NewSales(types, sellingDate string) Financial {
	switch types {
	case "inside":
		return InsideSalesImpl{sellingDate}
	case "outside":
		return OutsideSalesImpl{sellingDate}
	default:
		return nil
	}
}

func main() {
	insideSales := NewSales("inside", "20220601")
	log.Printf("%v", insideSales)
}
