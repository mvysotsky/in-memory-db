package main

import "testing"

func TestTransaction(t *testing.T) {
	db := NewMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	db.Commit()
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
}

func TestRollBack(t *testing.T) {
	db := NewMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
	db.RollBack()
	if val, _ := db.Get("key1"); val != "value1" {
		t.Errorf("Get() after rollback = %v, want %v", val, "value1")
	}
}

func TestNestedTransactions(t *testing.T) {
	db := NewMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
	db.StartTransaction()
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside nested transaction = %v, want %v", val, "value2")
	}
	db.Delete("key1")
	db.Commit()
	if val, exists := db.Get("key1"); exists != false {
		t.Errorf("Get() after nested transaction commit = %v, want exists == %v", val, false)
	}
	db.Commit()
	if val, exists := db.Get("key1"); exists != false {
		t.Errorf("Get() after nested transaction commit = %v, want exists == %v", val, false)
	}
}

func TestNestedTransactionsRollback(t *testing.T) {
	db := NewMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
	db.StartTransaction()
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside nested transaction = %v, want %v", val, "value2")
	}
	db.Delete("key1")
	db.RollBack()
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
	db.Commit()
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
	db.Commit()
	if val, _ := db.Get("key1"); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
}

func TestMemoryDB(t *testing.T) {
	db := NewMemoryDB()
	key := "testKey"
	value := "testValue"

	// Test Set and Get
	db.Set(key, value)
	if val, exists := db.Get(key); !exists || val != value {
		t.Errorf("Get() = %v, want %v", val, value)
	}

	// Test Delete
	db.Delete(key)
	if _, exists := db.Get(key); exists {
		t.Errorf("Delete() failed, key %v still exists", key)
	}

	// Test Transactions
	db.Set(key, "value1")
	db.StartTransaction()
	db.Set(key, "value2")
	if val, _ := db.Get(key); val != "value2" {
		t.Errorf("Get() inside transaction = %v, want %v", val, "value2")
	}
	db.RollBack()
	if val, _ := db.Get(key); val != "value1" {
		t.Errorf("Get() after rollback = %v, want %v", val, "value1")
	}

	db.StartTransaction()
	db.Set(key, "value3")
	db.Commit()
	if val, _ := db.Get(key); val != "value3" {
		t.Errorf("Get() after commit = %v, want %v", val, "value3")
	}
}
