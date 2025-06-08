package datasql

import (
	"context"
	"database/sql"
	"errors"
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

func (r *couponRepository) GetByCoupon(ctx context.Context, coupon string) (model.CouponModel, error) {
	var couponModel model.CouponModel

	query := `
        SELECT id, code, discount_percent, valid_until, created_at
        FROM coupon
        WHERE code = ?
        ORDER BY created_at DESC
        LIMIT 1
    `
	row := r.conn.QueryRowContext(ctx, query, coupon)
	err := row.Scan(
		&couponModel.ID,
		&couponModel.Code,
		&couponModel.DiscountPercent,
		&couponModel.ValidUntil,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.CouponModel{}, fmt.Errorf("coupon not found")
		}
		return model.CouponModel{}, fmt.Errorf("failed to fetch coupon: %w", err)
	}
	return couponModel, nil

}
