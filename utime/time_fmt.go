package utime

import "time"

const (
	YearOnly = "2006"
	YM       = "2006-01"
	DateOnly = "2006-01-02"
	YMDH     = "2006-01-02 15"
	YMDHm    = "2006-01-02 15:04"
	DateTime = "2006-01-02 15:04:05"
	TimeOnly = "15:04:05"
)

const (
	YMTerse       = "200601"
	DateOnlyTerse = "20060102"
	YMDHTerse     = "2006010215"
	YMDHmTerse    = "200601021504"
	DateTimeTerse = "20060102150405"
	TimeOnlyTerse = "150405"
)

func St2Terse(stTime, fromLayout, toLayout string) (string, error) {
	pt, err := time.Parse(fromLayout, stTime)
	if err != nil {
		return "", err
	}
	ft := pt.Format(toLayout)
	return ft, nil
}
