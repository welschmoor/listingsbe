package data

import (
	"fmt"
	"strconv"
)

// Declare a custom Price type, which has the underlying type int32 (the same as our
// Listing struct field).
type Price int64

// Implement a MarshalJSON() method on the Price type so that it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the listing
// price (in our case, it will return a string in the format "<price> dallas").
func (r Price) MarshalJSON() ([]byte, error) {
	// Generate a string containing the listng price in the required format.
	jsonValue := fmt.Sprintf("%d dallas", r)
	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quotedJSONValue := strconv.Quote(jsonValue)
	// Convert the quoted string value to a byte slice and return it.
	return []byte(quotedJSONValue), nil
}
