package zenorm

import (
	"database/sql"
)

func parseRows(rows *sql.Rows, err error) ([]map[string]string, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	ret := []map[string]string{}
	rowp := make([]interface{}, len(columns))
	for rows.Next() {
		row := make([]string, len(columns))
		for i := range row {
			rowp[i] = &row[i]
		}

		if err := rows.Scan(rowp...); err != nil {
			return nil, err
		}
		rowm := map[string]string{}
		for i := range columns {
			rowm[columns[i]] = row[i]
		}
		ret = append(ret, rowm)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
