package main

// TransactionStack is a stack of transactions.
type TransactionStack struct {
	transactions []*MemoryDB
}

// Push adds a new transaction to the stack.
func (ts *TransactionStack) Push(transaction *MemoryDB) {
	ts.transactions = append(ts.transactions, transaction)
}

// Pop removes the last transaction from the stack.
func (ts *TransactionStack) Pop() *MemoryDB {
	if len(ts.transactions) == 0 {
		return nil
	}
	lastIndex := len(ts.transactions) - 1
	transaction := ts.transactions[lastIndex]
	ts.transactions = ts.transactions[:lastIndex]
	return transaction
}

// Peek returns the last transaction without removing it.
func (ts *TransactionStack) Peek() *MemoryDB {
	if len(ts.transactions) == 0 {
		return nil
	}
	return ts.transactions[len(ts.transactions)-1]
}
