package models

import (
	"errors"
)

type TransactionType string

const (
	Credit TransactionType = "c"
	Debit  TransactionType = "d"
)

type Customer struct {
	Id           uint8 `json:"id"`
	Balance      int32 `json:"balance"`
	Limit        int32 `json:"limit"`
	Transactions []Transaction
}

func (c *Customer) Transact(transaction Transaction) error {
	switch transaction.TransactionType {
	case Credit:
		{
			c.Balance += transaction.Value
			c.Transactions = append(c.Transactions, transaction)
		}
	case Debit:
		{
			if c.Balance+c.Limit >= transaction.Value {
				c.Balance -= transaction.Value
				c.Transactions = append(c.Transactions, transaction)
			} else {
				return errors.New("não tem limite suficiente")
			}
		}
	default:
		return errors.New("unknown TransactionType")
	}

	return nil
}

type Transaction struct {
	Value           int32           `json:"valor"`
	TransactionType TransactionType `json:"tipo"`
	Description     string          `json:"descricao"`
}
