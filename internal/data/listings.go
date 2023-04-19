package data

import (
	"letsgofurther/internal/validator"
	"time"
)

type Listing struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Categories  []string  `json:"categories"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ValidateListing(v *validator.Validator, listing *Listing) {
	v.Check(listing.Title != "", "title", "must be provided")
	v.Check(len(listing.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(listing.Description != "", "description", "must be provided")
	v.Check(len(listing.Description) <= 1000, "description", "must not be more than 1000 bytes long")

	v.Check(listing.Price > 0, "price", "must be provided")
	v.Check(listing.Price <= 1_000_000, "price", "must be less than one million")

	v.Check(listing.Categories != nil, "categories", "must be provided")
	v.Check(len(listing.Categories) >= 1, "categories", "must contain at least 1 genre")
	v.Check(len(listing.Categories) <= 5, "categories", "must not contain more than 5 genres")
	v.Check(validator.Unique(listing.Categories), "categories", "must not contain duplicate values")
}
