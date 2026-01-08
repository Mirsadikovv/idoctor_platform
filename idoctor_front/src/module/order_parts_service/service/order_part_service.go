package order_part_service

import (
	"context"
	"encoding/json"

	order_part_dto "github.com/Mirsadikovv/idoctor_platform/src/module/order_parts_service/dto"
	order_part_model "github.com/Mirsadikovv/idoctor_platform/src/module/order_parts_service/model"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"github.com/Mirsadikovv/shared/response"
	"gorm.io/gorm"
)

type OrderPartService interface {
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*order_part_dto.OrderPartPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]order_part_dto.OrderPartResponse, error)
	FindOne(ctx context.Context, filter pg.Filter) (*order_part_dto.OrderPartResponse, error)
	Create(ctx context.Context, orderPartDto *order_part_dto.OrderPartCreate) (int64, error)
	Update(ctx context.Context, id int64, orderPartDto *order_part_dto.OrderPartUpdate) error
	DeleteOrRestore(ctx context.Context, id int64) error
}

type orderPartService struct {
	db *gorm.DB
}

func NewOrderPartService(db *gorm.DB) OrderPartService {
	return &orderPartService{
		db: db,
	}
}

func (s *orderPartService) Find(ctx context.Context, filter pg.Filter) ([]order_part_dto.OrderPartResponse, error) {
	var orderParts []order_part_model.OrderPart

	query := s.db.WithContext(ctx)

	if filter != nil {
		query = filter(query)
	}

	if err := query.Preload("Order").Preload("Part").Preload("Supplier").Find(&orderParts).Error; err != nil {
		return nil, err
	}

	responses := make([]order_part_dto.OrderPartResponse, len(orderParts))
	for i, op := range orderParts {
		responses[i] = order_part_dto.OrderPartResponse{
			Id:          op.Id,
			OrderId:     op.OrderId,
			PartId:      op.PartId,
			SupplierId:  op.SupplierId,
			IncomePrice: op.IncomePrice,
			Price:       op.Price,
			CreatedAt:   op.CreatedAt,
		}

		if op.Order != nil {
			if data, err := json.Marshal(op.Order); err == nil {
				responses[i].Order = data
			}
		}

		if op.Part != nil {
			if data, err := json.Marshal(op.Part); err == nil {
				responses[i].Part = data
			}
		}

		if op.Supplier != nil {
			if data, err := json.Marshal(op.Supplier); err == nil {
				responses[i].Supplier = data
			}
		}
	}

	return responses, nil
}

func (s *orderPartService) FindOne(ctx context.Context, filter pg.Filter) (*order_part_dto.OrderPartResponse, error) {
	var orderPart order_part_model.OrderPart

	query := s.db.WithContext(ctx)

	if filter != nil {
		query = filter(query)
	}

	if err := query.Preload("Order").Preload("Part").Preload("Supplier").First(&orderPart).Error; err != nil {
		return nil, err
	}

	response := &order_part_dto.OrderPartResponse{
		Id:          orderPart.Id,
		OrderId:     orderPart.OrderId,
		PartId:      orderPart.PartId,
		SupplierId:  orderPart.SupplierId,
		IncomePrice: orderPart.IncomePrice,
		Price:       orderPart.Price,
		CreatedAt:   orderPart.CreatedAt,
	}

	if orderPart.Order != nil {
		if data, err := json.Marshal(orderPart.Order); err == nil {
			response.Order = data
		}
	}

	if orderPart.Part != nil {
		if data, err := json.Marshal(orderPart.Part); err == nil {
			response.Part = data
		}
	}

	if orderPart.Supplier != nil {
		if data, err := json.Marshal(orderPart.Supplier); err == nil {
			response.Supplier = data
		}
	}

	return response, nil
}

func (s *orderPartService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*order_part_dto.OrderPartPage, error) {
	var orderParts []order_part_model.OrderPart
	var total int64

	query := s.db.WithContext(ctx).Model(&order_part_model.OrderPart{})

	if filter != nil {
		query = filter(query)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (paginate.Page() - 1) * paginate.PerPage()
	if err := query.Offset(offset).Limit(paginate.PerPage()).
		Preload("Order").Preload("Part").Preload("Supplier").
		Find(&orderParts).Error; err != nil {
		return nil, err
	}

	responses := make([]order_part_dto.OrderPartResponse, len(orderParts))
	for i, op := range orderParts {
		responses[i] = order_part_dto.OrderPartResponse{
			Id:          op.Id,
			OrderId:     op.OrderId,
			PartId:      op.PartId,
			SupplierId:  op.SupplierId,
			IncomePrice: op.IncomePrice,
			Price:       op.Price,
			CreatedAt:   op.CreatedAt,
		}

		if op.Order != nil {
			if data, err := json.Marshal(op.Order); err == nil {
				responses[i].Order = data
			}
		}

		if op.Part != nil {
			if data, err := json.Marshal(op.Part); err == nil {
				responses[i].Part = data
			}
		}

		if op.Supplier != nil {
			if data, err := json.Marshal(op.Supplier); err == nil {
				responses[i].Supplier = data
			}
		}
	}

	totalPages := int64(0)
	if paginate.PerPage() > 0 {
		totalPages = (total + int64(paginate.PerPage()) - 1) / int64(paginate.PerPage())
	}

	return &response.PageData[order_part_dto.OrderPartResponse]{
		Data:        responses,
		Total:       total,
		TotalPages:  totalPages,
		CurrentPage: paginate.Page(),
		PageSize:    paginate.PerPage(),
	}, nil
}

func (s *orderPartService) Create(ctx context.Context, orderPartDto *order_part_dto.OrderPartCreate) (int64, error) {
	orderPartModel := &order_part_model.OrderPart{
		OrderId:     orderPartDto.OrderId,
		PartId:      orderPartDto.PartId,
		SupplierId:  orderPartDto.SupplierId,
		IncomePrice: orderPartDto.IncomePrice,
		Price:       orderPartDto.Price,
	}

	if err := s.db.WithContext(ctx).Create(orderPartModel).Error; err != nil {
		return 0, err
	}

	return orderPartModel.Id, nil
}

func (s *orderPartService) Update(ctx context.Context, id int64, orderPartDto *order_part_dto.OrderPartUpdate) error {
	updates := make(map[string]interface{})

	if orderPartDto.SupplierId != nil {
		updates["supplier_id"] = *orderPartDto.SupplierId
	}
	if orderPartDto.IncomePrice != nil {
		updates["income_price"] = *orderPartDto.IncomePrice
	}
	if orderPartDto.Price != nil {
		updates["price"] = *orderPartDto.Price
	}

	if len(updates) == 0 {
		return nil
	}

	return s.db.WithContext(ctx).Model(&order_part_model.OrderPart{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (s *orderPartService) DeleteOrRestore(ctx context.Context, id int64) error {
	var orderPart order_part_model.OrderPart

	if err := s.db.WithContext(ctx).Unscoped().Where("id = ?", id).First(&orderPart).Error; err != nil {
		return err
	}

	if orderPart.DeletedAt.Valid {
		return s.db.WithContext(ctx).Model(&order_part_model.OrderPart{}).
			Where("id = ?", id).
			Update("deleted_at", nil).Error
	}

	return s.db.WithContext(ctx).Where("id = ?", id).Delete(&order_part_model.OrderPart{}).Error
}
