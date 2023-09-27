package main

import (
	"fmt"
	"sync"
)

// Several goroutines have access to 1 Family Account Balance and tries to spend
// money from there at the same time. It can cause race condition => as result:
// we spent more money than we originally had.
func main() {
	wg := sync.WaitGroup{}
	familyAccount := &Account{10000}
	var total int
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			total += familyAccount.spendMoney()
		}()
	}

	wg.Wait()
	fmt.Printf("Total spent: %d \nBalance: %d", total, familyAccount.balance)
}

type Account struct {
	balance int
}

func (a *Account) spend(value int) bool {
	if value <= a.balance {
		a.balance -= value
		return true
	}
	return false
}

func (a *Account) spendMoney() int {
	var totalSpent int
	for i := 0; i < 10000; i++ {
		if a.spend(1) {
			totalSpent += 1
		}
	}
	return totalSpent
}
