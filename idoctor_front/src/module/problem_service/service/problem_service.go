package problem_service

import (
	"context"

	problem_dto "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/dto"
	problem_model "github.com/Mirsadikovv/idoctor_platform/src/module/problem_service/model"

	"github.com/Mirsadikovv/shared/pg"
	"github.com/Mirsadikovv/shared/request"
	"gorm.io/gorm"
)

type ProblemService interface {
	Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*problem_dto.ProblemPage, error)
	Find(ctx context.Context, filter pg.Filter) ([]problem_dto.Problem, error)
	FindOne(ctx context.Context, filter pg.Filter) (*problem_dto.Problem, error)
	CreateOrUpdate(ctx context.Context, problemDto *problem_dto.ProblemCreate) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type problemService struct {
	db *gorm.DB
}

func NewProblemService(db *gorm.DB) ProblemService {
	return &problemService{
		db: db,
	}
}

func (s *problemService) Find(ctx context.Context, filter pg.Filter) ([]problem_dto.Problem, error) {
	return pg.FindWithScan[problem_model.Problem, problem_dto.Problem](s.db, filter)
}

func (s *problemService) FindOne(ctx context.Context, filter pg.Filter) (*problem_dto.Problem, error) {
	return pg.FindOneWithScan[problem_model.Problem, problem_dto.Problem](s.db, filter)
}

func (s *problemService) Page(ctx context.Context, paginate *request.Paginate, filter pg.Filter) (*problem_dto.ProblemPage, error) {
	return pg.PageWithScan[problem_model.Problem, problem_dto.Problem](s.db.WithContext(ctx), paginate, filter)
}

func (s *problemService) CreateOrUpdate(ctx context.Context, problemDto *problem_dto.ProblemCreate) (int64, error) {
	problemModel := &problem_model.Problem{
		Name:  problemDto.Name,
		Price: problemDto.Price,
	}

	if problemDto.Id == 0 {
		if err := pg.Create(s.db.WithContext(ctx), problemModel, "id"); err != nil {
			return 0, err
		}

		return problemModel.Id, nil
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", problemDto.Id)
	}

	if _, err := pg.Update[problem_model.Problem](s.db.WithContext(ctx), problemModel, filter, "id"); err != nil {
		return 0, err
	}

	return problemDto.Id, nil
}

func (s *problemService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Delete(&problem_model.Problem{}).Error
}
