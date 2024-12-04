package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"strings"
	"regexp"
	
	"github.com/NickMoorman123/receipt-processor/errors"
	"github.com/NickMoorman123/receipt-processor/objects"
)

type Response interface {
	Json() []byte
	StatusCode() int
}

func WriteResponse(w http.ResponseWriter, res Response) {
	w.WriteHeader(res.StatusCode())
	_, _ = w.Write(res.Json())
}

func WriteError(w http.ResponseWriter, err error) {
	res, ok := err.(*errors.Error)
	if !ok {
		log.Println(err)
		res = errors.ErrInternal
	}
	WriteResponse(w, res)
}

func Unmarshal(w http.ResponseWriter, data []byte, v interface{}) error {
	if d := string(data); d == "null" || d == "" {
		WriteError(w, errors.ErrObjectIsRequired)
		return errors.ErrObjectIsRequired
	}
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Println(err)
		WriteError(w, errors.ErrBadRequest)
	}
	return err
}

func calculateReceiptPoints(receipt *objects.Receipt) error {
	receipt.Retailer = strings.TrimSpace(receipt.Retailer)
	receipt.PurchaseDate = strings.TrimSpace(receipt.PurchaseDate)
	receipt.PurchaseTime = strings.TrimSpace(receipt.PurchaseTime)
	receipt.Total = strings.TrimSpace(strings.Replace(receipt.Total, ",", "", -1))

	points, err := getRetailerPoints(receipt.Retailer)
	if err != nil {
		return err
	}
	receipt.Points += points

	points, err = getPurchaseDatePoints(receipt.PurchaseDate)
	if err != nil {
		return err
	}
	receipt.Points += points

	points, err = getPurchaseTimePoints(receipt.PurchaseTime)
	if err != nil {
		return err
	}
	receipt.Points += points

	points, totalCheck, err := getItemPoints(receipt.Items)
	if err != nil {
		return err
	}
	receipt.Points += points

	points, err = getTotalAmountPoints(receipt.Total, totalCheck)
	if err != nil {
		return err
	}
	receipt.Points += points

	return nil
}

func getRetailerPoints(retailer string) (int, error) {
	if retailer == "" {
		return 0, errors.ErrReceiptRetailerIsRequired
	}
	if !validRetailerCharacters(retailer) {
		return 0, errors.ErrReceiptRetailerInvalidCharacter
	}

	alphanumeric := regexp.MustCompile("[A-Za-z0-9]")
	points := len(alphanumeric.FindAllStringIndex(retailer, -1))
	return points, nil
}

func getPurchaseDatePoints(purchaseDate string) (int, error) {
	if purchaseDate == "" {
		return 0, errors.ErrReceiptDateIsRequired
	}
	parsedDate, dateErr := time.Parse("2006-01-02", purchaseDate)
	if dateErr != nil {
		return 0, errors.ErrInvalidDateFormat
	}

	points := 0
	if parsedDate.Day() % 2 == 1 {
		points += 6
	}
	return points, nil
}

func getPurchaseTimePoints(purchaseTime string) (int, error) {
	if purchaseTime == "" {
		return 0, errors.ErrReceiptTimeIsRequired
	}
	parsedTime, timeErr := time.Parse("15:04", purchaseTime)
	if timeErr != nil {
		return 0, errors.ErrInvalidTimeFormat
	}

	points := 0
	if parsedTime.Hour() == 14 || parsedTime.Hour() == 15 {
		points += 10
	}
	return points, nil
}

func getItemPoints(items []objects.Item) (int, int, error) {
	if len(items) == 0 {
		return 0, 0, errors.ErrReceiptItemsAreRequired
	}
	points := (len(items) / 2) * 5

	totalCentsCheck := 0
	for _, item := range items {
		item.ShortDescription = strings.TrimSpace(item.ShortDescription)
		item.Price = strings.TrimSpace(strings.Replace(item.Price, ",", "", -1))
		
		if item.ShortDescription == "" {
			return 0, 0, errors.ErrItemDescriptionIsRequired
		}
		if !validShortDescriptionCharacters(item.ShortDescription) {
			return 0, 0, errors.ErrItemDescriptionInvalidCharacter
		}
		if item.Price == "" {
			return 0, 0, errors.ErrItemPriceIsRequired
		}
		if !validMoneyFormat(item.Price) {
			return 0, 0, errors.ErrInvalidPriceFormat
		}

		itemCents := toCents(item.Price)
		if len(item.ShortDescription) % 3 == 0 {
			// equivalent of 20% of price, rounded up
			if (itemCents * 2) % 1000 == 0 {
				points += itemCents * 2 / 1000
			} else {
				points += (itemCents * 2 / 1000) + 1
			}
		}
		totalCentsCheck += itemCents
	}
	
	return points, totalCentsCheck, nil
}

func getTotalAmountPoints(total string, totalCentsCheck int) (int, error) {
	if total == "" {
		return 0, errors.ErrReceiptTotalIsRequired
	}
	if !validMoneyFormat(total) {
		return 0, errors.ErrInvalidTotalFormat
	}
	if toCents(total) != totalCentsCheck {
		return 0, errors.ErrReceiptTotalIsIncorrect
	}
	
	points := 0
	cents, _ := strconv.Atoi(total[len(total)-2:])
	if cents % 25 == 0 {
		points += 25
		if cents == 0 {
			points += 50
		}
	}
	return points, nil
}

func validRetailerCharacters(retailer string) bool {
	return getFirstVal(regexp.MatchString("^[\\w\\s\\-&]+$", retailer))
}

func validShortDescriptionCharacters(retailer string) bool {
	return getFirstVal(regexp.MatchString("^[\\w\\s\\-]+$", retailer))
}

func validMoneyFormat(moneyString string) bool {
	return getFirstVal(regexp.MatchString("^\\d+\\.\\d{2}$", moneyString))
}

func toCents(moneyString string) int {
	return getFirstVal(strconv.Atoi(strings.Replace(moneyString, ".", "", -1)))
}

func getFirstVal[T, U any](val T, _ U) T {
    return val
}