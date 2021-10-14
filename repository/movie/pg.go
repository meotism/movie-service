package movie

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/phamtrung99/movie-service/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type pgRepository struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getClient}
}

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.Movie, error) {
	db := r.getClient(ctx)
	movie := &model.Movie{}

	err := db.Where("id = ?", ID).
		First(movie).Error

	if err != nil {
		return nil, errors.Wrap(err, "get movie by id")
	}

	return movie, nil
}

func (r *pgRepository) Find(
	ctx context.Context,
	conditions []model.Condition,
	paginator *model.Paginator,
	orders []string,
) (*model.MovieResult, error) {
	// Build query
	db := r.getClient(ctx)
	query := db.Model(&model.Movie{})

	// Where
	for _, condition := range conditions {
		switch strings.ToLower(condition.Type) {
		case model.ConditionTypeNot:
			query.Not(condition.Pattern, condition.Values...)
		case model.ConditionTypeOr:
			query.Or(condition.Pattern, condition.Values...)
		default:
			query.Where(condition.Pattern, condition.Values...)
		}
	}

	// Order
	for _, order := range orders {
		query.Order(order)
	}

	// Paging
	var result model.MovieResult

	if paginator.Limit >= 0 {
		if paginator.Page <= 0 {
			paginator.Page = 1
		}

		if paginator.Limit == 0 {
			paginator.Limit = model.PageSize
		}

		result.Page = paginator.Page
		result.Limit = paginator.Limit
		query.Count(&result.Total).Scopes(paginator.Paginate())
	}

	err := query.Find(&result.Data).Error

	return &result, err
}

func (r *pgRepository) SearchMovie(
	ctx context.Context,
	paginator *model.Paginator,
	str string,
	filter *model.MovieFilter,
	orders []string) (*model.MovieResult, error) {

	db := r.getClient(ctx)
	query := db.Model(&model.Movie{})

	// Order
	for _, order := range orders {
		query.Order(order)
	}

	filterActor := ""
	filterCate := ""
	filterRate := ""
	filterAdult := ""
	filterName := ""

	if filter.ActorID != 0 {
		filterActor = "JOIN movie_actors ON movie_actors.movie_id = movies.id AND movie_actors.actor_id = " + strconv.FormatInt(filter.ActorID, 10)
	}

	if filter.CateID != 0 {
		filterCate = "JOIN movie_categories ON movie_categories.movie_id = movies.id AND movie_categories.cate_id = " + strconv.FormatInt(filter.CateID, 10)
	}

	if filter.IsAdult != -1 {
		filterAdult = "AND is_adult = " + strconv.Itoa(filter.IsAdult)
	}

	if filter.MinRating != -1 {
		filterRate = "AND rating_average > " + strconv.Itoa(filter.MinRating)
	}

	if str != "" {
		filterName = "AND original_title LIKE " + "'%" + str + "%'"
	}

	queryStr := "SELECT * FROM movies " +
		filterActor + " " +
		filterCate + " WHERE isnull(movies.deleted_at) " +
		filterAdult + " " +
		filterRate + " " +
		filterName

	// Paging
	var result model.MovieResult
	var kq *gorm.DB

	if paginator.Limit >= 0 {
		if paginator.Page <= 0 {
			paginator.Page = 1
		}

		if paginator.Limit == 0 {
			paginator.Limit = model.PageSize
		}

		result.Page = paginator.Page
		result.Limit = paginator.Limit
		limitQuery := " LIMIT " + fmt.Sprint(result.Limit) + " OFFSET " + fmt.Sprint((result.Page-1)*result.Limit)
		kq = query.Raw(fmt.Sprint(queryStr, limitQuery)).Find(&result.Data)

		result.Total = query.Raw(queryStr).Find(&[]model.Movie{}).RowsAffected
	}

	return &result, kq.Error
}

func (r *pgRepository) Insert(ctx context.Context, movie *model.Movie) (*model.Movie, error) {
	db := r.getClient(ctx)
	err := db.Create(movie).Error

	return movie, errors.Wrap(err, "create movie")
}

func (r *pgRepository) Delete(ctx context.Context, id int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", id).Delete(&model.Movie{}).Error

	return errors.Wrap(err, "delete movie fail")
}

func (r *pgRepository) Update(ctx context.Context, movie *model.Movie) (*model.Movie, error) {
	db := r.getClient(ctx)
	err := db.Save(movie).Error

	return movie, errors.Wrap(err, "update movie")
}
