package ports

import (
	"context"
	"github.com/gilxa1226/microservices/payment/internal/application/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Payment, error)
	Save(ctx context.Context, payment *domain.Payment) error
}
