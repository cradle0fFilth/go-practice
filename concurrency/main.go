package main

import (
	"sync"
	"sync/atomic"
)

type Account struct {
	Balance int64
	InTx    bool
}

func transfer1(amount int64, accountFrom, accountTo *Account) bool {
	if accountFrom.Balance < amount {
		return false
	}

	accountTo.Balance += amount
	accountFrom.Balance -= amount
	return true
}

func transfer2(amount int64, accountFrom, accountTo *Account) bool {
	bal := atomic.LoadInt64(&accountFrom.Balance)
	if bal < amount {
		return false
	}
	atomic.AddInt64(&accountTo.Balance, amount)
	atomic.AddInt64(&accountFrom.Balance, -amount)

	return true
}

var txMutex sync.Mutex

func transfer3(amount int64, accountFrom, accountTo *Account) bool {
	txMutex.Lock()

	//在函数结束的时候执行此语句
	defer txMutex.Unlock()

	bal := atomic.LoadInt64(&accountFrom.Balance)
	if bal < amount {
		return false
	}
	atomic.AddInt64(&accountTo.Balance, amount)
	atomic.AddInt64(&accountFrom.Balance, -amount)

	return true
}

var Mutex sync.Mutex

func transfer4(amount int64, accountFrom, accountTo *Account) bool {
	accountFrom.InTx = true
	accountTo.InTx = true

	defer func() {
		accountTo.InTx = false
		accountFrom.InTx = false
	}()

	Mutex.Lock()
	defer txMutex.Unlock()
	bal := atomic.LoadInt64(&accountFrom.Balance)
	if bal < amount {
		return false
	}

	atomic.AddInt64(&accountTo.Balance, amount)
	atomic.AddInt64(&accountFrom.Balance, -amount)
	return true
}
