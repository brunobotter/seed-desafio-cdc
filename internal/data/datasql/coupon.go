package datasql

import (
	"context"
	"fmt"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type couponRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *couponRepository) Save(ctx context.Context, coupon entity.Coupon) (model.CouponModel, error) {
	couponModel := model.ToCouponModel(coupon)

	insertPaymentQuery := `
		INSERT INTO coupon (code, discount_percent, valid_until, created_at)
		VALUES (?, ?, ?, ?)
	`
	result, err := r.conn.ExecContext(
		ctx,
		insertPaymentQuery,
		couponModel.Code,
		couponModel.DiscountPercent,
		couponModel.ValidUntil,
		couponModel.CreatedAt,
	)
	if err != nil {
		return model.CouponModel{}, fmt.Errorf("failed to insert coupon: %w", err)
	}

	paymentID, err := result.LastInsertId()
	if err != nil {
		return model.CouponModel{}, fmt.Errorf("failed to get inserted coupon ID: %w", err)
	}
	couponModel.ID = paymentID

	return couponModel, nil
}
