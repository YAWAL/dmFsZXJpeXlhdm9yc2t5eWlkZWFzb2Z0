package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/YAWAL/dmFsZXJpeXlhdm9yc2t5eWlkZWFzb2Z0/src/handler"
	"github.com/YAWAL/dmFsZXJpeXlhdm9yc2t5eWlkZWFzb2Z0/src/order"
)

func main() {
	orders := order.Generate()

	orderSlice := order.GenerateSlice(orders)
	orderMap := order.GenerateMap(orders)

	dynamicOrders := order.NewDynamicOrders(orderSlice)

	handler.Register(dynamicOrders, orderMap)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tick := time.NewTicker(200 * time.Millisecond)

	go func() {
		for {
			select {
			case <-tick.C:
				newOrder := order.GenerateOne()
				dynamicOrders.AddOrder(newOrder)
				orderMap.Store(newOrder, 0)
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-tick.C:
				deleted := dynamicOrders.DeleteOrder()
				orderMap.Delete(deleted)
			case <-ctx.Done():
				return
			}
		}
	}()

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
