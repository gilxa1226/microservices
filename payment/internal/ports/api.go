package ports

import (
	"context"
	"github.com/gilxa1226/microservices/payment/internal/application/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
