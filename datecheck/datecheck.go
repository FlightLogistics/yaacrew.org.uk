package datecheck

import (
	"strconv"
	"time"
)

type ExpiryDate struct {
	Day   int
	Month time.Month
	Year  int
}

type ExpiryUnit int

const (
	Day ExpiryUnit = iota
	Month
	Year
)

// AddMonth returns the time corresponding to adding the
// given number of months to t.
// For example, AddMonth(t, 2) applied to January 1, 2011
// (= t) returns March 1, 2011.
//
// AddMonth does not normalize its result in the same way
// that Date does, so, for example, adding one month to
// October 31 yields November 30.

func addMonth(t time.Time, months int) time.Time {
	lastMonthDay := func(t time.Time) int {
		return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1).Day()
	}

	newDate := func(t time.Time, day int) time.Time {
		return time.Date(t.Year(), t.Month(), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	// Creating 1st Date from t and adding months because AddDate() normalizes t.
	am := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).AddDate(0, months, 0)
	ad := lastMonthDay(am)

	if ld := lastMonthDay(t); t.Day() == ld {
		return newDate(am, ad)
	} else {
		if t.Day() > ad {
			return newDate(am, ad)
		}
		return t.AddDate(0, months, 0)
	}
}

func GetExpiryDate(dateOfCheck time.Time, durationToGoToExpiry int, durationUnits ExpiryUnit, calculateToEndOfMonth bool) ExpiryDate {

	var checkExpiryDate time.Time
	var exactExpiryDate time.Time

	if durationUnits == Month {
		exactExpiryDate = addMonth(dateOfCheck, durationToGoToExpiry)

	} else if durationUnits == Day {
		calculateToEndOfMonth = false
		exactExpiryDate = dateOfCheck.AddDate(0, 0, durationToGoToExpiry)

	} else if durationUnits == Year {
		exactExpiryDate.AddDate(durationToGoToExpiry, 0, 0)
	}

	if calculateToEndOfMonth {
		currentLocation := time.Now().Location()
		firstOfMonth := time.Date(exactExpiryDate.Year(), exactExpiryDate.Month(), 1, 0, 0, 0, 0, currentLocation)
		expiryDateToEndOfMonth := firstOfMonth.AddDate(0, 1, -1)
		checkExpiryDate = expiryDateToEndOfMonth
	} else {
		checkExpiryDate = exactExpiryDate
	}

	return ExpiryDate{
		Day:   checkExpiryDate.Day(),
		Month: checkExpiryDate.Month(),
		Year:  checkExpiryDate.Year(),
	}
}

func ConvertExpiryDateToString(expiryDate ExpiryDate) string {
	return strconv.Itoa(expiryDate.Day) + " " + expiryDate.Month.String() + " " + strconv.Itoa(expiryDate.Year)
}
