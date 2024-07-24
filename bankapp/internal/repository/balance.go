package repository

import (
	"bankapp/internal/domain"
	"bankapp/internal/utils"
	"fmt"
	"sync"
	"time"
)

type BalanceRepository struct {
	balances map[string]*domain.Balance
	mu       sync.Mutex
}

func NewBalanceRepository() *BalanceRepository {
	return &BalanceRepository{
		balances: make(map[string]*domain.Balance),
	}
}

func (r *BalanceRepository) Create(balance *domain.Balance) *domain.Balance {
	r.balances[balance.ID] = balance
	return balance
}

func (r *BalanceRepository) GetByID(balanceID string) (*domain.Balance, error) {
	if balance, exists := r.balances[balanceID]; exists {
		return balance, nil
	}

	return nil, utils.ErrBalanceNotFound
}

func (r *BalanceRepository) AddAmount(balanceID string, amount float64) (*domain.Balance, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if balance, exists := r.balances[balanceID]; exists {
		time.Sleep(time.Millisecond * 5)

		balance.Amount += amount
		return balance, nil
	}

	return nil, utils.ErrBalanceNotFound
}

func (r *BalanceRepository) WithdrawAmount(balanceID string, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	balance, exists := r.balances[balanceID]
	if !exists {
		return utils.ErrBalanceNotFound
	}

	time.Sleep(time.Millisecond * 5)

	balance.Amount -= amount

	fmt.Printf("withdraw id:%s | username: %s | amount: %.2f\n", balance.ID, balance.Username, balance.Amount)
	return nil
}
