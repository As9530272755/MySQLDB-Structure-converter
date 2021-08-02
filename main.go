package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	types := map[string]string{
		"INT":      "sql.NullInt32",
		"VARCHAR":  "string",
		"TINYINT":  "byte",
		"CHAR":     "string",
		"SMALLINT": "int16",
	}

	driverName := "mysql"
	dsn := "root:root@tcp(10.0.0.10:3306)/hellodb?charset=utf8mb4&loc=Local&parseTime=true"
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	SQL := "SHOW TABLES;"

	rows, err := db.Query(SQL)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			file, _ := os.Create(tableName + ".go")

			findTable := "select * from " + tableName

			findRows, err := db.Query(findTable)
			if err != nil {
				fmt.Println(err)
				return
			}
			columns, err := findRows.ColumnTypes()
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Fprintf(file, "type \t%s struct \t{\n", tableName)
				for _, column := range columns {

					fmt.Fprintf(file, "\t%s \t%s\n", column.Name(), types[column.DatabaseTypeName()])

				}
				fmt.Fprintf(file, "\n}")
			}
		}
	}
}
