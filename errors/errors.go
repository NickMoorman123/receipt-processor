package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	// ErrInternal HTTP 500
	ErrInternal = &Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
	}
	// ErrUnprocessableEntity HTTP 422
	ErrUnprocessableEntity = &Error{
		Code:    http.StatusUnprocessableEntity,
		Message: "Unprocessable Entity",
	}
	// ErrObjectIsRequired HTTP 400
	ErrObjectIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Request object should be provided",
	}
	// ErrBadRequest HTTP 400
	ErrBadRequest = &Error{
		Code:    http.StatusBadRequest,
		Message: "Error invalid argument",
	}
	// ErrReceiptNotFound HTTP 404
	ErrReceiptNotFound = &Error{
		Code:    http.StatusNotFound,
		Message: "Receipt not found",
	}
	// ErrReceiptRetailerIsRequired HTTP 400
	ErrReceiptRetailerIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt retailer is required",
	}
	// ErrReceiptRetailerIsRequired HTTP 400
	ErrReceiptRetailerInvalidCharacter = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt retailer has invalid character",
	}
	// ErrReceiptDateIsRequired HTTP 400
	ErrReceiptDateIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt purchase date is required",
	}
	// ErrInvalidDateFormat HTTP 400
	ErrInvalidDateFormat = &Error{
		Code:    http.StatusBadRequest,
		Message: "Date should be passed in ISO 8601 YYYY-mm-DD format",
	}
	// ErrReceiptTimeIsRequired HTTP 400
	ErrReceiptTimeIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt purchase time is required",
	}
	// ErrInvalidTimeFormat HTTP 400
	ErrInvalidTimeFormat = &Error{
		Code:    http.StatusBadRequest,
		Message: "Time should be passed in ISO 8601 hh:mm format",
	}
	// ErrReceiptItemsAreRequired HTTP 400
	ErrReceiptItemsAreRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt must have at least one item",
	}
	// ErrItemDescriptionIsRequired HTTP 400
	ErrItemDescriptionIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Item description is required",
	}
	// ErrItemDescriptionIsRequired HTTP 400
	ErrItemDescriptionInvalidCharacter = &Error{
		Code:    http.StatusBadRequest,
		Message: "Item description has invalid character",
	}
	// ErrItemPriceIsRequired HTTP 400
	ErrItemPriceIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Item price is required",
	}
	// ErrInvalidPriceFormat HTTP 400
	ErrInvalidPriceFormat = &Error{
		Code:    http.StatusBadRequest,
		Message: "Price should be passed in #.## format",
	}
	// ErrReceiptTotalIsRequired HTTP 400
	ErrReceiptTotalIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt total is required",
	}
	// ErrInvalidPriceFormat HTTP 400
	ErrInvalidTotalFormat = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt total should be passed in #.## format",
	}
	// ErrReceiptTotalIsIncorrect HTTP 400
	ErrReceiptTotalIsIncorrect = &Error{
		Code:    http.StatusBadRequest,
		Message: "Receipt total must equal total of item prices",
	}
)

// Error main object for error
type Error struct {
	Code    int
	Message string
}

func (err *Error) Error() string {
	return err.String()
}

func (err *Error) String() string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("error: code=%s message=%s", http.StatusText(err.Code), err.Message)
}

// JSON convert Error in json
func (err *Error) Json() []byte {
	if err == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(err)
	return res
}

// StatusCode get status code
func (err *Error) StatusCode() int {
	if err == nil {
		return http.StatusOK
	}
	return err.Code
}