package data

import (
	"database/sql"
	"errors"
	"letsgofurther/internal/validator"
	"time"

	"github.com/lib/pq"
)

type Listing struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Categories  []string  `json:"categories"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     int32     `json:"version"`
}

func ValidateListing(v *validator.Validator, listing *Listing) {
	v.Check(listing.Title != "", "title", "must be provided")
	v.Check(len(listing.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(listing.Description != "", "description", "must be provided")
	v.Check(len(listing.Description) <= 1000, "description", "must not be more than 1000 bytes long")

	v.Check(listing.Price > 0, "price", "must be provided")
	v.Check(listing.Price <= 1_000_000, "price", "must be less than one million")

	v.Check(listing.Categories != nil, "categories", "must be provided")
	v.Check(len(listing.Categories) >= 1, "categories", "must contain at least 1 category")
	v.Check(len(listing.Categories) <= 5, "categories", "must not contain more than 5 category")
	v.Check(validator.Unique(listing.Categories), "categories", "must not contain duplicate values")
}

/* MODEL */

// Define a Listings Model struct type which wraps a sql.DB connection pool.
type ListingModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the listings table.
func (lm ListingModel) Insert(listing *Listing) error {
	rows := lm.DB.QueryRow(
		`INSERT INTO listings (title, description, price, categories) VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version`,
		listing.Title,
		listing.Description,
		listing.Price,
		pq.Array(listing.Categories),
	)
	// if err != nil {
	// 	return err
	// }
	// defer rows.Close()

	err := rows.Scan(
		&listing.ID,
		&listing.CreatedAt,
		&listing.Version,
	)
	if err != nil {
		return err
	}

	return nil
}

// Add a placeholder method for fetching a specific record from the listings table.
func (lm ListingModel) Select(id int64) (*Listing, error) {
	if id < 1 {
		return nil, ErrNotFoundRecord
	}

	var lis Listing
	rows := lm.DB.QueryRow(
		`SELECT id, title, description, price, categories, created_at, updated_at, version
		FROM listings
		WHERE id = $1;`,
		id,
	)

	err := rows.Scan(
		&lis.ID,
		&lis.Title,
		&lis.Description,
		&lis.Price,
		pq.Array(&lis.Categories),
		&lis.CreatedAt,
		&lis.UpdatedAt,
		&lis.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFoundRecord
		default:
			return nil, err
		}
	}

	return &lis, nil
}

// Add a placeholder method for updating a specific record in the listings table.
func (lm ListingModel) Update(listing *Listing) error {
	return nil
}

// Add a placeholder method for deleting a specific record from the listings table.
func (lm ListingModel) Delete(id int64) error {
	return nil
}

/* MOCK MODEL */

type MockListingModel struct{}

func (lm MockListingModel) Insert(listing *Listing) error { // Mock the action...
	return nil
}
func (lm MockListingModel) Select(id int64) (*Listing, error) { // Mock the action...
	return &Listing{}, nil
}
func (lm MockListingModel) Update(listing *Listing) error { // Mock the action...
	return nil
}
func (lm MockListingModel) Delete(id int64) error { // Mock the action...
	return nil
}
