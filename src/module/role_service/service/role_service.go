package role_service

import (
	"context"

	role_dto "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/dto"
	role_model "github.com/Mirsadikovv/idoctor_platform/src/module/role_service/model"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
)

type RoleService interface {
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*role_dto.RolePage, error)
	Find(ctx context.Context, filter pg.Filter) ([]role_dto.Role, error)
	FindOne(ctx context.Context, filter pg.Filter) (*role_dto.Role, error)
	CreateOrUpdate(ctx context.Context, roleDto *role_dto.RoleCreate) (int64, error)
}

type roleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{
		db: db,
	}
}

func (r *roleService) Find(ctx context.Context, filter pg.Filter) ([]role_dto.Role, error) {
	return pg.FindWithScan[role_model.Role, role_dto.Role](r.db, filter)
}

func (r *roleService) FindOne(ctx context.Context, filter pg.Filter) (*role_dto.Role, error) {
	return pg.FindOneWithScan[role_model.Role, role_dto.Role](r.db, filter)
}

func (r *roleService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*role_dto.RolePage, error) {
	return pg.PageWithScan[role_model.Role, role_dto.Role](r.db.WithContext(ctx), paginate, filter)
}

func (r *roleService) CreateOrUpdate(ctx context.Context, roleDto *role_dto.RoleCreate) (int64, error) {

	roleModel := &role_model.Role{
		Name:        roleDto.Name,
		Description: roleDto.Description,
		Pages:       roleDto.Pages,
		Permissions: roleDto.Permissions,
	}

	if roleDto.ID == 0 {
		if err := pg.Create(r.db.WithContext(ctx), roleModel, "id"); err != nil {
			return 0, err
		}

		return roleModel.ID, nil
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", roleDto.ID)
	}

	if _, err := pg.Update[role_model.Role](r.db.WithContext(ctx), roleModel, filter, "id"); err != nil {
		return 0, err
	}

	return roleDto.ID, nil
}
