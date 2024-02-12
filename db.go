package main

// Database interface defines the methods for key-value storage.
type Database interface {
	Get(key string) (string, bool)
	Set(key string, value string)
	Delete(key string) bool
}

// Transaction interface defines the methods for transaction control.
type Transaction interface {
	StartTransaction()
	Commit() bool
	RollBack() bool
}
