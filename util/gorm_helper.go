package util

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func RawJSON(ctx context.Context, db *gorm.DB, sqlString string) (result string, err error) {
	if db == nil {
		err = errors.New("gorm db is nil")
		return
	}

	rows, err := db.WithContext(ctx).Raw(sqlString).Rows()
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	var (
		count     = len(columns)
		tableData = make([]map[string]interface{}, 0)
		values    = make([]interface{}, count)
		valuePtrs = make([]interface{}, count)
	)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return Byte2Str(jsonData), nil
}
