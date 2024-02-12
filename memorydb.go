package main

type MemoryDBStorage map[string]string

// MemoryDB is an in-memory key-value store with transaction support.
type MemoryDB struct {
	storage         MemoryDBStorage
	transactionMode bool
	transactions    TransactionStack
}

// NewMemoryDB creates a new MemoryDB instance.
func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		storage: make(map[string]string),
	}
}

// Get retrieves the value for a key.
func (db *MemoryDB) Get(key string) (string, bool) {
	if db.transactionMode {
		transaction := db.transactions.Peek()
		value, exists := transaction.storage[key]
		return value, exists
	}

	value, exists := db.storage[key]
	return value, exists
}

// Set sets the value for a key.
func (db *MemoryDB) Set(key string, value string) {
	if db.transactionMode {
		transaction := db.transactions.Peek()
		transaction.storage[key] = value
		return
	}
	db.storage[key] = value
}

// Delete removes a key from the store.
func (db *MemoryDB) Delete(key string) bool {
	var storage MemoryDBStorage

	if db.transactionMode {
		transaction := db.transactions.Peek()
		storage = transaction.storage
	} else {
		storage = db.storage
	}

	if _, exists := storage[key]; exists {
		delete(storage, key)
		return true
	}

	return false
}

// StartTransaction begins a new transaction.
func (db *MemoryDB) StartTransaction() {
	db.transactionMode = true
	newTransaction := &MemoryDB{
		storage:         make(map[string]string),
		transactionMode: false,
	}

	latestTransaction := db.transactions.Peek()
	if latestTransaction != nil {
		for k, v := range latestTransaction.storage {
			newTransaction.storage[k] = v
		}
	} else {
		for k, v := range db.storage {
			newTransaction.storage[k] = v
		}
	}

	db.transactions.Push(newTransaction)
}

// Commit finalizes the current transaction.
func (db *MemoryDB) Commit() bool {
	if !db.transactionMode {
		return false
	}

	transaction := db.transactions.Pop()
	if transaction == nil {
		return false
	}

	if parentTransaction := db.transactions.Peek(); parentTransaction != nil {
		parentTransaction.storage = transaction.storage
		return true
	}

	db.storage = transaction.storage
	db.transactionMode = len(db.transactions.transactions) > 0
	return true
}

// RollBack cancels the current transaction.
func (db *MemoryDB) RollBack() bool {
	if !db.transactionMode {
		return false
	}
	db.transactions.Pop()
	db.transactionMode = len(db.transactions.transactions) > 0
	return true
}
