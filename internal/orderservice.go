// ./internal/orderservice.go

package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/koyeb/example-grpc-api/protogen/golang/orders"
)

// OrderService should implement the OrdersServer interface generated from grpc.
//
// UnimplementedOrdersServer must be embedded to have forwarded compatible implementations.
type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

// NewOrderService creates a new OrderService
func NewOrderService(db *DB) OrderService {
	return OrderService{db: db}
}

// AddOrder implements the AddOrder method of the grpc OrdersServer interface to add a new order
func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an add order request")

	err := o.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, err
}

// GetOrder implements the GetOrder method of the grpc OrdersServer interface to fetch an order for a given orderID
func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Received get order request")

	order := o.db.GetOrderByID(req.GetOrderId())
	if order == nil {
		return nil, fmt.Errorf("order not found for orderID: %d", req.GetOrderId())
	}

	return &orders.PayloadWithSingleOrder{Order: order}, nil
}

// UpdateOrder implements the UpdateOrder method of the grpc OrdersServer interface to update an order
func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an update order request")

	o.db.UpdateOrder(req.GetOrder())

	return &orders.Empty{}, nil
}

// RemoveOrder implements the RemoveOrder method of the grpc OrdersServer interface to remove an order
func (o *OrderService) RemoveOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.Empty, error) {
	log.Printf("Received a remove order request")

	o.db.RemoveOrder(req.GetOrderId())

	return &orders.Empty{}, nil
}
