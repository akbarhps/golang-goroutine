package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0

	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}

	fmt.Println("Counter: ", counter)
}

func TestRaceConditionWithMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", counter)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.RWMutex.Lock()
	defer b.RWMutex.Unlock()

	b.Balance += amount
}

func (b *BankAccount) GetBalance() int {
	b.RWMutex.RLock()
	defer b.RWMutex.RUnlock()

	return b.Balance
}

func TestBankAccount(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.Deposit(1)
				fmt.Println("Balance: ", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance: ", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func Transfer(from, to *UserBalance, amount int) {
	from.Lock()
	defer from.Unlock()

	fmt.Printf("%s Locked\n", from.Name)
	from.Balance -= amount

	time.Sleep(time.Second)

	to.Lock()
	defer to.Unlock()

	fmt.Printf("%s Locked\n", to.Name)
	to.Balance += amount

	time.Sleep(time.Second)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "User1",
		Balance: 10000,
	}
	user2 := UserBalance{
		Name:    "User2",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 100)
	go Transfer(&user2, &user1, 200)

	time.Sleep(3 * time.Second)

	fmt.Printf("User %s Balance %d\n", user1.Name, user1.Balance) // expect 10100
	fmt.Printf("User %s Balance %d\n", user2.Name, user2.Balance) // expect 9900
}
