package order

import (
	"math/rand"
	"sync"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func Generate() map[string]int {
	orders := make(map[string]int)
	for {
		seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		order := make([]byte, 2)
		for i := range order {
			order[i] = alphabet[seededRand.Intn(len(alphabet))]
		}
		orders[string(order)] = 0
		if len(orders) == 50 {
			return orders
		}
	}
}

func GenerateMap(items map[string]int) *sync.Map {
	var orders sync.Map
	for k, v := range items {
		orders.Store(k, v)
	}
	return &orders
}

func GenerateSlice(items map[string]int) []string {
	orders := make([]string, 0)
	for item := range items {
		orders = append(orders, item)
	}
	return orders
}

func GenerateOne() string {
	order := make([]byte, 2)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range order {
		order[i] = alphabet[seededRand.Intn(len(alphabet))]
	}
	return string(order)
}

type DynamicOrders struct {
	orders []string
	mu     sync.RWMutex
}

func NewDynamicOrders(orders []string) *DynamicOrders {
	return &DynamicOrders{
		orders: orders,
		mu:     sync.RWMutex{},
	}
}

func (d *DynamicOrders) AddOrder(order string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.orders = append(d.orders, order)
}

func (d *DynamicOrders) GetOrder() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano() + time.Now().UnixNano() + time.Now().UnixNano()))
	return d.orders[seededRand.Intn(len(d.orders))]
}

func (d *DynamicOrders) DeleteOrder() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	k := seededRand.Intn(len(d.orders))
	deleted := d.orders[k]
	d.orders = append(d.orders[:k], d.orders[k:]...)
	return deleted
}
