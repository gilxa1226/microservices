package ports

import "github.com/gilxa1226/microservices/order/internal/application/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
