// ./internal/db.go
package internal

import (
	"fmt"

	"github.com/koyeb/example-grpc-api/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

// NewDB creates a new array to mimic the behaviour of a in-memory database
func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

// AddOrder adds a new order to the DB collection. Returns an error on duplicate ids
func (d *DB) AddOrder(order *orders.Order) error {
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("duplicate order id: %d", order.GetOrderId())
		}
	}
	d.collection = append(d.collection, order)
	return nil
}

// GetOrderByID returns an order by the order_id
func (d *DB) GetOrderByID(orderID uint64) *orders.Order {
	for _, o := range d.collection {
		if o.OrderId == orderID {
			return o
		}
	}
	return nil
}

// GetOrdersByIDs returns all orders pertaining to the given order ids
func (d *DB) GetOrdersByIDs(orderIDs []uint64) []*orders.Order {
	filtered := make([]*orders.Order, 0)

	for _, idx := range orderIDs {
		for _, order := range d.collection {
			if order.OrderId == idx {
				filtered = append(filtered, order)
				break
			}
		}
	}

	return filtered
}

// UpdateOrder updates an order in place
func (d *DB) UpdateOrder(order *orders.Order) {
	for i, o := range d.collection {
		if o.OrderId == order.OrderId {
			d.collection[i] = order
			return
		}
	}
}

// RemoveOrder removes an order from the orders collection
func (d *DB) RemoveOrder(orderID uint64) {
	filtered := make([]*orders.Order, 0, len(d.collection)-1)
	for i := range d.collection {
		if d.collection[i].OrderId != orderID {
			filtered = append(filtered, d.collection[i])
		}
	}
	d.collection = filtered
}
