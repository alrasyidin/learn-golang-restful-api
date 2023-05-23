package repository

import (
	"belajar_golang_api/helper"
	"belajar_golang_api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (categoryRepo *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO categories (name) values(?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)

	helper.HandleIfPanicError(err)

	id, err := result.LastInsertId()
	helper.HandleIfPanicError(err)
	category.Id = id
	return category
}

func (categoryRepo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)

	helper.HandleIfPanicError(err)

	return category
}

func (categoryRepo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)

	helper.HandleIfPanicError(err)
}

func (categoryRepo *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (domain.Category, error) {
	sql := "SELECT id, name FROM categories WHERE id=?"
	rows, err := tx.QueryContext(ctx, sql, categoryId)

	helper.HandleIfPanicError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.HandleIfPanicError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (categoryRepo *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, sql)
	helper.HandleIfPanicError(err)
	defer rows.Close()

	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}

		err = rows.Scan(&category.Id, &category.Name)

		helper.HandleIfPanicError(err)
		categories = append(categories, category)
	}

	return categories
}
