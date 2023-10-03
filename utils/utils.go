package utils

import (
	"encoding/json"
	"gboardist/database"
	"gboardist/global"
	"math"
	"time"

	"github.com/craftzbay/go_grc/v2/converter"
	"github.com/craftzbay/go_grc/v2/data"
	"gorm.io/datatypes"
)

func GetTableName(name ...interface{}) string {
	var tableName string
	var alias string
	for idx, v := range name {
		switch idx {
		case 0:
			tableName = v.(string)
		case 1:
			alias = v.(string)
		}
	}
	tName := global.Conf.DBSchema + "." + global.Conf.DBTablePrefix + "_" + tableName
	if alias != "" {
		tName += " as " + alias
	}
	return tName
}

func DateFilter(req *map[string]string) (string, string) {
	filters := *req
	// dates from {start_date} to {end_date} ---> 2023-07-05 - 2023-07-06
	startStr := filters["start_date"]
	endStr := filters["end_date"]
	// last {duration} days	---> {duration} == 7 ? 2023-06-29 - 2023-07-05
	durationStr := filters["duration"]
	// days in {year}-{month} ---> {year} == 2023 & {month} == 7 ? 2023-07-01 - 2023-07-31
	yearStr := filters["year"]
	monthStr := filters["month"]

	// no filter if none of the above is given
	if startStr == "" && endStr == "" && durationStr == "" && yearStr == "" && monthStr == "" {
		return "", ""
	}
	// conversions
	startTime := converter.DateStringToTime(startStr)
	endTime := converter.DateStringToTime(endStr)
	duration := converter.StringToInt(durationStr)
	year := converter.StringToInt(yearStr)
	month := converter.StringToInt(monthStr)

	if year == 0 {
		year = time.Now().Year()
	}
	if month == 0 {
		month = int(time.Now().Month())
	}
	if duration == 0 {
		if startTime.IsZero() && endTime.IsZero() {
			startTime = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
			endTime = time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Now().Location())
		} else {
			if startTime.IsZero() {
				startTime = time.Now()
			}
			if endTime.IsZero() {
				endTime = time.Now()
			}
		}
	} else {
		endTime = time.Now()
		startTime = endTime.Add(-time.Hour * 24 * time.Duration(duration))
	}
	endTime = endTime.Add(time.Hour * 24)

	return data.LocalDate(startTime).String(), data.LocalDate(endTime).String()
}

func CategoryFilter(categoryStr string) []uint {
	res := make([]uint, 0)
	categoryJson := datatypes.JSON(categoryStr)
	categoryArray := make([]uint, 0)

	if err := json.Unmarshal(categoryJson, &categoryArray); err != nil {
		return res
	}

	for _, elem := range categoryArray {
		res = append(res, CategorySubIds(elem)...)
		res = append(res, elem)
	}

	res = RemoveDuplicatesUint(res)

	return res
}

func StrToIntArray(str string) []uint {
	res := make([]uint, 0)
	strJson := datatypes.JSON(str)
	intArray := make([]uint, 0)

	if err := json.Unmarshal(strJson, &intArray); err != nil {
		return res
	}
	res = RemoveDuplicatesUint(intArray)

	return res
}

func GetOrder(req string) string {
	switch req {
	case "new":
		return "created_at DESC"
	case "old":
		return "created_at ASC"
	case "cheap":
		return "price ASC"
	case "expensive":
		return "price DESC"
	case "featured":
		return "view_count DESC"
	default:
		return ""
	}
}

type CategoryParent struct {
	Id       uint
	ParentId uint
}

var (
	categories  []CategoryParent
	categoryIds []uint
)

func CategorySubIds(parentId uint) []uint {
	db := database.DBconn

	if err := db.Table(GetTableName("category")).Find(&categories).Error; err != nil {
		return nil
	}

	categorySubIdsDFS(parentId)

	return categoryIds
}

func categorySubIdsDFS(parentId uint) {
	for _, c := range categories {
		if ContainsUint(categoryIds, c.Id) {
			continue
		}
		if c.ParentId == parentId {
			categoryIds = append(categoryIds, c.Id)
			categorySubIdsDFS(c.Id)
		}
	}
}

func MapToSlice[T any](m *map[interface{}]T) []T {
	req := *m
	res := make([]T, 0)
	for _, val := range req {
		res = append(res, val)
	}
	return res
}

func RemoveDuplicatesUint(list []uint) []uint {
	res := make([]uint, 0)

	for _, elem := range list {
		if !ContainsUint(res, elem) {
			res = append(res, elem)
		}
	}

	return res
}

func ContainsUint(list []uint, num uint) bool {
	for _, x := range list {
		if num == x {
			return true
		}
	}
	return false
}

func DayStartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func DayEndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, t.Location())
}

func CheckExists(table string, columns []string, values []interface{}) bool {
	var exists bool
	if len(columns) != len(values) {
		return false
	}
	db := database.DBconn
	tx := db.Table(GetTableName(table)).Select("count(*) > 0")

	for i := range columns {
		tx.Where(columns[i]+" = ?", values[i])
	}
	if err := tx.Take(&exists).Error; err != nil {
		return false
	}
	return exists
}

func RoundFloat(v float64) float64 {
	val := math.Floor(v*100) / 100
	return val
}
