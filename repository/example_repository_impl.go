package repository

import (
	"context"
	"database/sql"
	"errors"
	"simplify-go/helper"
	"simplify-go/model/entity"
)

type PhotoRepositoryImpl struct{}

func NewPhotosRepositoryImpl() PhotosRepository {
	return &PhotoRepositoryImpl{}
}

func (repository *PhotoRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, Photo entity.Photos) entity.Photos {
	SQL := "INSERT INTO photo (id, name, email, created_at, updated_at) VALUES (?,?,?,?,?)"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		Photo.Id,
		Photo.Name,
		Photo.Email,
		Photo.CreatedAt,
		Photo.UpdatedAt,
	)
	helper.PanicIfError(err)

	return Photo
}

func (repository *PhotoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, Photo entity.Photos) entity.Photos {
	SQL := "UPDATE photo SET name=?, updated_at=? WHERE id=?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		Photo.Name,
		Photo.UpdatedAt,
		Photo.Id,
	)
	helper.PanicIfError(err)

	return Photo
}

func (repository *PhotoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, photo entity.Photos) {
	SQL := "DELETE FROM photo WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		photo.Id,
	)
	helper.PanicIfError(err)
}

func (repository *PhotoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, photoId string) (entity.Photos, error) {
	SQL := "SELECT id, name, email, created_at, updated_at FROM photo WHERE id=?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		photoId,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	Photo := entity.Photos{}
	if rows.Next() {
		err := rows.Scan(
			&Photo.Id,
			&Photo.Name,
			&Photo.Email,
			&Photo.CreatedAt,
			&Photo.UpdatedAt,
		)
		helper.PanicIfError(err)

		return Photo, nil
	}

	return Photo, errors.New("photo not found")
}

func (repository *PhotoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Photos {
	SQL := "SELECT id, name, email, created_at, updated_at FROM photo"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	Photo := []entity.Photos{}
	for rows.Next() {
		photo := entity.Photos{}
		err := rows.Scan(
			&photo.Id,
			&photo.Name,
			&photo.Email,
			&photo.CreatedAt,
			photo.UpdatedAt,
		)
		helper.PanicIfError(err)
		Photo = append(Photo, photo)
	}

	return Photo
}
