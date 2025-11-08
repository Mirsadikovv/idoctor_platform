package order_service

import (
	"context"
	"time"

	order_dto "github.com/Mirsadikovv/idoctor_platform/src/module/order_service/dto"
	order_model "github.com/Mirsadikovv/idoctor_platform/src/module/order_service/model"
	part_model "github.com/Mirsadikovv/idoctor_platform/src/module/part_service/model"
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
	var orders []order_model.Order

	tx := s.db.WithContext(ctx)
	if filter != nil {
		tx = filter(tx)
	}

	if err := tx.Preload("Parts").Preload("Problems").Find(&orders).Error; err != nil {
		return nil, err
	}

	// Convert to DTOs
	result := make([]order_dto.Order, len(orders))
	for i, order := range orders {
		result[i] = order_dto.Order{
			ID:            order.ID,
			ClientID:      order.ClientID,
			MasterID:      order.MasterID,
			Price:         order.Price,
			Status:        order.Status,
			PaymentType:   order.PaymentType,
			PaymentStatus: order.PaymentStatus,
			CreatedAt:     order.CreatedAt,
			DeletedAt:     convertDeletedAt(order.DeletedAt),
		}

		// Extract part IDs
		if len(order.Parts) > 0 {
			result[i].PartIDs = make([]int64, len(order.Parts))
			for j, part := range order.Parts {
				result[i].PartIDs[j] = part.ID
			}
		}

		// Extract problem IDs
		if len(order.Problems) > 0 {
			result[i].ProblemIDs = make([]int64, len(order.Problems))
			for j, problem := range order.Problems {
				result[i].ProblemIDs[j] = problem.ID
			}
		}
	}

	return result, nil
}

func (s *orderService) FindOne(ctx context.Context, filter pg.Filter) (*order_dto.Order, error) {
	var order order_model.Order

	tx := s.db.WithContext(ctx)
	if filter != nil {
		tx = filter(tx)
	}

	if err := tx.Preload("Parts").Preload("Problems").First(&order).Error; err != nil {
		return nil, err
	}

	result := &order_dto.Order{
		ID:            order.ID,
		ClientID:      order.ClientID,
		MasterID:      order.MasterID,
		Price:         order.Price,
		Status:        order.Status,
		PaymentType:   order.PaymentType,
		PaymentStatus: order.PaymentStatus,
		CreatedAt:     order.CreatedAt,
		DeletedAt:     convertDeletedAt(order.DeletedAt),
	}

	// Extract part IDs
	if len(order.Parts) > 0 {
		result.PartIDs = make([]int64, len(order.Parts))
		for i, part := range order.Parts {
			result.PartIDs[i] = part.ID
		}
	}

	// Extract problem IDs
	if len(order.Problems) > 0 {
		result.ProblemIDs = make([]int64, len(order.Problems))
		for i, problem := range order.Problems {
			result.ProblemIDs[i] = problem.ID
		}
	}

	return result, nil
}

func (s *orderService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*order_dto.OrderPage, error) {
	return pg.PageWithScan[order_model.Order, order_dto.Order](s.db.WithContext(ctx), paginate, filter)
}

func (s *orderService) Create(ctx context.Context, orderDto *order_dto.OrderCreate) (int64, error) {
	orderModel := &order_model.Order{
		ClientID:      orderDto.ClientID,
		MasterID:      orderDto.MasterID,
		Price:         orderDto.Price,
		Status:        orderDto.Status,
		PaymentType:   orderDto.PaymentType,
		PaymentStatus: orderDto.PaymentStatus,
	}

	// Start transaction
	return orderModel.ID, s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create order
		if err := tx.Create(orderModel).Error; err != nil {
			return err
		}

		// Associate parts
		if len(orderDto.PartIDs) > 0 {
			var parts []part_model.Part
			if err := tx.Where("id IN ?", orderDto.PartIDs).Find(&parts).Error; err != nil {
				return err
			}
			if err := tx.Model(orderModel).Association("Parts").Replace(parts); err != nil {
				return err
			}
		}

		// Associate problems
		if len(orderDto.ProblemIDs) > 0 {
			var problems []problem_model.Problem
			if err := tx.Where("id IN ?", orderDto.ProblemIDs).Find(&problems).Error; err != nil {
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
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Find existing order
		var order order_model.Order
		if err := tx.Where("id = ?", id).First(&order).Error; err != nil {
			return err
		}

		// Update basic fields
		order.ClientID = orderDto.ClientID
		order.MasterID = orderDto.MasterID
		order.Price = orderDto.Price
		order.Status = orderDto.Status
		order.PaymentType = orderDto.PaymentType
		order.PaymentStatus = orderDto.PaymentStatus

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		// Update parts association
		if orderDto.PartIDs != nil {
			var parts []part_model.Part
			if len(orderDto.PartIDs) > 0 {
				if err := tx.Where("id IN ?", orderDto.PartIDs).Find(&parts).Error; err != nil {
					return err
				}
			}
			if err := tx.Model(&order).Association("Parts").Replace(parts); err != nil {
				return err
			}
		}

		// Update problems association
		if orderDto.ProblemIDs != nil {
			var problems []problem_model.Problem
			if len(orderDto.ProblemIDs) > 0 {
				if err := tx.Where("id IN ?", orderDto.ProblemIDs).Find(&problems).Error; err != nil {
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

func convertDeletedAt(deletedAt *gorm.DeletedAt) *time.Time {
	if deletedAt != nil && deletedAt.Valid {
		return &deletedAt.Time
	}
	return nil
}
