package main

import (
	"fmt"
)

/*

	Beispiel:

	Der Quotient steht im Zusammenhang mit der Division ganzer Zahlen,
	die Anzahl ganzer Zahlen im Zähler und im Nenner.

	Mit anderen Worten, er ist identisch mit der Dezimalanweisung: FLOOR(n/d)

	Mit Modulo erhält man den Rest einer solchen Division. Der Modulo eines Zählers
	und eines Nenners liegt immer zwischen 0 und d-1 (wobei d der Nenner ist)

*/

func integerDivision(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func main() {
	hours, remainder := integerDivision(5566, 3600)
	minutes, seconds := integerDivision(remainder, 60)
	fmt.Printf("%d:%d:%d", hours, minutes, seconds)
}
