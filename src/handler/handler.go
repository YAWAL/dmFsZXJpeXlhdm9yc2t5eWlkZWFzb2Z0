package handler

import (
	"github.com/YAWAL/dmFsZXJpeXlhdm9yc2t5eWlkZWFzb2Z0/src/order"
	"net/http"
	"strconv"
	"sync"
)

func Register(orderSlice *order.DynamicOrders, orderMap *sync.Map) {
	http.HandleFunc("/request", findOrder(orderSlice, orderMap))
	http.HandleFunc("/admin/requests", showOrders(orderMap))

}

func findOrder(orderSlice *order.DynamicOrders, orderMap *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item := orderSlice.GetOrder()
		counter, ok := orderMap.Load(item)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Can't get order"))
			return
		}
		intCounter := counter.(int)
		intCounter = intCounter + 1
		orderMap.Store(item, intCounter)
		w.Write([]byte(item))
	}
}

func showOrders(orderMap *sync.Map) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var orders []string
		orderMap.Range(func(key, value interface{}) bool {
			if value.(int) == 0 {
				stringKey := key.(string)
				stringValue := strconv.Itoa(value.(int))
				orders = append(orders, stringKey+" - "+stringValue)
			}
			return true
		})
		var response string

		for _, order := range orders {
			response = response + order + "\n"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
}
