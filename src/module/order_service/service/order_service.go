package order_service

import (
	"context"
	"fmt"

	order_dto "github.com/Mirsadikovv/idoctor_platform/src/module/order_service/dto"
	order_model "github.com/Mirsadikovv/idoctor_platform/src/module/order_service/model"
	problem_model "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/model"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
)

type OrderService interface {
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*order_dto.OrderPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]order_dto.Order, error)
	FindOne(ctx context.Context, filter pg.Filter) (*order_dto.Order, error)
	Create(ctx context.Context, orderDto *order_dto.OrderCreate) (int64, error)
	Update(ctx context.Context, id int64, orderDto *order_dto.OrderUpdate) error
	DeleteOrRestore(ctx context.Context, id int64) error
}

type orderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{
		db: db,
	}
}

func (s *orderService) Find(ctx context.Context, filter pg.Filter) ([]order_dto.Order, error) {
	var orders []order_dto.Order

	tx := s.db.WithContext(ctx)
	if filter != nil {
		tx = filter(tx)
	}

	if err := tx.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderService) FindOne(ctx context.Context, filter pg.Filter) (*order_dto.Order, error) {
	return pg.FindOneWithScan[order_model.Order, order_dto.Order](s.db.WithContext(ctx), filter)
}

func (s *orderService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*order_dto.OrderPage, error) {
	return pg.PageWithScan[order_model.Order, order_dto.Order](s.db.WithContext(ctx), paginate, filter)
}

func (s *orderService) Create(ctx context.Context, orderDto *order_dto.OrderCreate) (int64, error) {
	// Validate status
	if !order_model.IsValidStatus(orderDto.Status) {
		return 0, fmt.Errorf("invalid order status: %s", orderDto.Status)
	}

	// Validate payment type
	if !order_model.IsValidPaymentType(orderDto.PaymentType) {
		return 0, fmt.Errorf("invalid payment type: %s", orderDto.PaymentType)
	}

	// Validate payment status
	if !order_model.IsValidPaymentStatus(orderDto.PaymentStatus) {
		return 0, fmt.Errorf("invalid payment status: %s", orderDto.PaymentStatus)
	}

	orderModel := &order_model.Order{
		ClientId:      orderDto.ClientId,
		ClientPhone:   orderDto.ClientPhone,
		ClientName:    orderDto.ClientName,
		PhonePassword: orderDto.PhonePassword,
		MasterId:      orderDto.MasterId,
		Price:         orderDto.Price,
		Status:        orderDto.Status,
		PaymentType:   orderDto.PaymentType,
		PaymentStatus: orderDto.PaymentStatus,
	}

	// Start transaction
	return orderModel.Id, s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create order
		if err := tx.Create(orderModel).Error; err != nil {
			return err
		}

		// Associate problems
		if len(orderDto.ProblemIds) > 0 {
			var problems []problem_model.Problem
			if err := tx.Where("id IN ?", orderDto.ProblemIds).Find(&problems).Error; err != nil {
				return err
			}
			if err := tx.Model(orderModel).Association("Problems").Replace(problems); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *orderService) Update(ctx context.Context, id int64, orderDto *order_dto.OrderUpdate) error {
	// Validate status
	if !order_model.IsValidStatus(orderDto.Status) {
		return fmt.Errorf("invalid order status: %s", orderDto.Status)
	}

	// Validate payment type
	if !order_model.IsValidPaymentType(orderDto.PaymentType) {
		return fmt.Errorf("invalid payment type: %s", orderDto.PaymentType)
	}

	// Validate payment status
	if !order_model.IsValidPaymentStatus(orderDto.PaymentStatus) {
		return fmt.Errorf("invalid payment status: %s", orderDto.PaymentStatus)
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Find existing order
		var order order_model.Order
		if err := tx.Where("id = ?", id).First(&order).Error; err != nil {
			return err
		}

		// Update basic fields
		order.ClientId = orderDto.ClientId
		order.ClientPhone = orderDto.ClientPhone
		order.ClientName = orderDto.ClientName
		order.PhonePassword = orderDto.PhonePassword
		order.MasterId = orderDto.MasterId
		order.Price = orderDto.Price
		order.Status = orderDto.Status
		order.PaymentType = orderDto.PaymentType
		order.PaymentStatus = orderDto.PaymentStatus

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		// Update problems association
		if orderDto.ProblemIds != nil {
			var problems []problem_model.Problem
			if len(orderDto.ProblemIds) > 0 {
				if err := tx.Where("id IN ?", orderDto.ProblemIds).Find(&problems).Error; err != nil {
					return err
				}
			}
			if err := tx.Model(&order).Association("Problems").Replace(problems); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *orderService) DeleteOrRestore(ctx context.Context, id int64) error {
	var order order_model.Order

	// Check if order exists (including soft deleted)
	if err := s.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&order).Error; err != nil {
		return err
	}

	// If deleted, restore; otherwise, soft delete
	if order.DeletedAt != nil && (*order.DeletedAt).Valid {
		// Restore
		return s.db.WithContext(ctx).Model(&order_model.Order{}).Where("id = ?", id).Update("deleted_at", nil).Error
	} else {
		// Soft delete
		return s.db.WithContext(ctx).Where("id = ?", id).Delete(&order_model.Order{}).Error
	}
}

