package datasql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/brunobotter/casa-codigo/internal/data/model"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/entity"
)

type bookRepository struct {
	conn executor
	data contract.RepoManager
}

func (r *bookRepository) Save(ctx context.Context, book entity.Book) (model.BookModel, error) {
	bookModel := model.ToBookModel(book)

	query := `
		INSERT INTO book (title, resume, sumary, price, page, isbn, publish_date, author_id, category_id, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	bookModel.CreatedAt = now

	result, err := r.conn.ExecContext(
		ctx,
		query,
		bookModel.Title,
		bookModel.Resume,
		bookModel.Summary,
		bookModel.Price,
		bookModel.Page,
		bookModel.ISBN,
		bookModel.PublishDate,
		bookModel.AuthorID,
		bookModel.CategoryID,
		bookModel.CreatedAt,
	)
	if err != nil {
		return model.BookModel{}, fmt.Errorf("failed to insert book: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.BookModel{}, fmt.Errorf("failed to get inserted ID: %w", err)
	}
	bookModel.ID = uint(id)
	return bookModel, nil
}

func (r *bookRepository) GetById(ctx context.Context, bookId int64) (model.BookByIdModel, error) {

	var bookModel model.BookByIdModel

	query := `
		SELECT 
			b.id,
			b.title,
			b.resume,
			b.price,
			b.page,
			b.isbn,
			b.publish_date,
			a.name AS author_name,
			c.name AS category_name
		FROM book b
		INNER JOIN authors a ON b.author_id = a.id
		INNER JOIN category c ON b.category_id = c.id
		WHERE b.id = ?
	`
	row := r.conn.QueryRowContext(ctx, query, bookId)
	err := row.Scan(
		&bookModel.ID,
		&bookModel.Title,
		&bookModel.Resume,
		&bookModel.Price,
		&bookModel.Page,
		&bookModel.ISBN,
		&bookModel.PublishDate,
		&bookModel.AuthorName,
		&bookModel.CategoryName,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.BookByIdModel{}, fmt.Errorf("book not found")
		}
		return model.BookByIdModel{}, fmt.Errorf("failed to fetch book: %w", err)
	}
	return bookModel, nil
}

func (r *bookRepository) GetAll(ctx context.Context) ([]model.BookByAllModel, error) {

	var books []model.BookByAllModel

	query := `
		SELECT 
			b.id,
			b.title
		FROM book b
	`
	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query books: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book model.BookByAllModel

		err := rows.Scan(
			&book.ID,
			&book.Title,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}

		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return books, nil
}
