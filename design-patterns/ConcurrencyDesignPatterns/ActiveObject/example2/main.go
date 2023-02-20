package main

import (
	"fmt"
	"time"
)

/*
	Mein Beispiel verwendet das das Bankkonto.
	Mit Hilfe eines aktiven Objektes verhindere ich hier eine sog. Race-Condition
*/

type Money int

const Euro Money = 100000

type (
	Account struct {
		bal              Money
		pubChan, priChan chan<- func()
	}
)

func NewAccount() *Account {

	pub := make(chan func(), 2)
	pri := make(chan func(), 20) // "private" chan

	go func() {
		for {
			if len(pri) > 0 {
				f := <-pri
				f()
			} else {
				select {
				case f := <-pri:
					f()
				case f := <-pub:
					f()
				}
			}
		}
	}()

	return &Account{
		bal:     100 * Euro,
		pubChan: pub,
		priChan: pri,
	}
}

// Deposit adds money to an account.
func (ac *Account) Deposit(amt Money) {
	ac.pubChan <- func() {
		current := ac.bal
		time.Sleep(1 * time.Millisecond)
		ac.bal = current + amt
	}
}

// Adds transfers money to/from an account.
func (ac *Account) Add(amt Money, callback func(error)) {
	ac.pubChan <- func() {
		if ac.bal+amt < 0 {
			callback(fmt.Errorf("insuff. funds %v for w/d %v",
				ac.bal, amt))
			return
		}
		ac.bal += amt
		callback(nil) // successful transfer
	}
}

// Balance provides funds available.
func (ac *Account) Balance(callback func(Money)) {
	ac.pubChan <- func() {
		callback(ac.bal)
	}
}

// TransferTo moves funds between accounts.
func (ac *Account) TransferTo(to *Account, amt Money, callback func(error)) {
	ac.pubChan <- func() {
		if amt < 0 {
			go to.TransferTo(ac, -amt, callback)
			return
		}
		if amt > ac.bal {
			callback(fmt.Errorf("Insuff. funds %v for tfr %v", ac.bal, amt))
			return
		}
		ac.bal -= amt
		go to.Add(amt, callback)
	}
}

func main() {

	var d Money = 1

	ac := NewAccount()
	ac.Add(1000*d, func(error) {})
	ac.Add(200*d, func(error) {})
	ac.Add(-1e6*d, func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	})
	ac.Balance(func(bal Money) {
		fmt.Printf("Balance: $%v\n", bal/100)
	})
	time.Sleep(100 * time.Millisecond)

}
