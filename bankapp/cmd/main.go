package main

import (
	"bankapp/internal/domain"
	"bankapp/internal/repository"
	"fmt"
	"runtime"
	"sync"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("CPU cores:", runtime.NumCPU())
	runtime.GOMAXPROCS(2)

	repo := repository.NewBalanceRepository()

	balance := repo.Create(&domain.Balance{ID: uuid.NewString(), Amount: 5500000, Username: "john"})
	fmt.Printf("(initial balance of '%s' added with ID '%s' & amount %.2f)\n",
		balance.Username, balance.ID, balance.Amount)

	operations := 500
	var wg sync.WaitGroup

	wg.Add(operations)

	fmt.Println("(simulating concurrent withdrawals)")
	for i := 0; i < operations; i++ {
		go func(index int) {
			defer wg.Done()

			err := repo.WithdrawAmount(balance.ID, 10000)
			if err != nil {
				fmt.Println("withdraw error:", err)
			}

			b, _ := repo.GetByID(balance.ID)
			showBalance(b, index)
		}(i)
	}

	wg.Wait()

	fmt.Println("(expected balance: 500000.00)")
	fmt.Printf("(final balance: %.2f)\n", balance.Amount)
}

func showBalance(balance *domain.Balance, index int) {
	fmt.Printf("withdraw-%d id:%s | username: %s | amount: %.2f\n",
		(index + 1), balance.ID, balance.Username, balance.Amount)
}
