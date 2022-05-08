package dto

import (
	"github.com/bhaktiutama/banking/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t TransactionRequest) IsTransactionTypeDeposit() bool {
	return t.TransactionType == DEPOSIT
}

func (t TransactionRequest) Validate() *errs.AppError {
	// logger.Info(fmt.Sprintf("test:%f", t.Amount))
	if t.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	if !t.IsTransactionTypeDeposit() && !t.IsTransactionTypeWithdrawal() {
		return errs.NewValidationError("Transaction type can only be withdrawal or deposit")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_d"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
