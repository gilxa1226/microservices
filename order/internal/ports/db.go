package ports

import "github.com/gilxa1226/microservices/order/internal/application/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
