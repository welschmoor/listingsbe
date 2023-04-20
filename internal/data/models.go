package data

import (
	"database/sql"
	"errors"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when // looking up a listing that doesn't exist in our database.
var (
	ErrNotFoundRecord = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the ListingModel. We'll add other models to this, // like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Listings interface {
		Insert(listing *Listing) error
		Select(id int64) (*Listing, error)
		SelectAll(title string, categories []string, filters Filters) ([]*Listing, Metadata, error)
		Update(listing *Listing) error
		Delete(id int64) error
	}
}

// For ease of use, we also add a New() method which returns a Models struct containing // the initialized ListingModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Listings: ListingModel{DB: db},
	}
}

// You can then call NewMockModels() whenever you need it in your unit tests in place of the ‘real’ NewModels() function
func NewMockModels(db *sql.DB) Models {
	return Models{
		Listings: MockListingModel{},
	}
}
