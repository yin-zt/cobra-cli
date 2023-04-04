package core

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yin-zt/cobra-cli/utils"
	"strings"
)

func (this *Cli) OperateSqlite(filename, sqlStr, table string) error {
	var (
		err      error
		v        interface{}
		t        string
		s        string
		isUpdate bool
		s2       string
		rows     []map[string]interface{}
		count    int64
		db       *sql.DB
	)

	if table == "" {
		t = "data"
	} else {
		t = table
	}
	s = sqlStr
	f := filename
	_ = s

	if s == "" {
		v, s = this.StdinJson()
	}

	// 此操作是判断sql执行语句是select还是update操作
	s2 = strings.TrimSpace(strings.ToLower(s))
	if strings.HasPrefix(s2, "select") {
		isUpdate = true
	} else {
		isUpdate = false
	}
	if v != nil {
		if db, err = this.SqliteInsert(f, t, v.([]interface{})); err != nil {
			corelog.Error(err)
			fmt.Println(err)
			return err
		}
		_ = db
		fmt.Println("ok")
	}
	if !isUpdate {
		count, err = this.SqliteExec(f, s)
		if err != nil {
			corelog.Error(err)
			fmt.Println(err)
			return err
		}
		fmt.Println(fmt.Sprintf("ok(%d)", count))
	} else {
		rows, err = this.SqliteQuery(f, s)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(utils.JsonEncodePretty(rows))
	}
	return nil
}

// SqliteQuery 适配在sqlite3上执行查询命令
func (this *Cli) SqliteQuery(filename string, s string) ([]map[string]interface{}, error) {
	var (
		err     error
		db      *sql.DB
		rows    *sql.Rows
		records []map[string]interface{}
	)

	if filename == "" {
		filename = ":memory:"
	}
	if db, err = sql.Open("sqlite3", filename); err != nil {
		return nil, err
	}

	rows, err = db.Query(s)

	if err != nil {
		corelog.Error(err)
		return nil, err
	}
	defer rows.Close()

	records = []map[string]interface{}{}
	for rows.Next() {
		record := map[string]interface{}{}

		columns, err := rows.Columns()
		if err != nil {
			corelog.Error(
				err, "unable to obtain rows columns",
			)
			continue
		}

		pointers := []interface{}{}
		for _, column := range columns {
			var value interface{}
			pointers = append(pointers, &value)
			record[column] = &value
		}

		err = rows.Scan(pointers...)
		if err != nil {
			corelog.Error(err, "can't read result records")
			continue
		}

		for key, value := range record {
			indirect := *value.(*interface{})
			if value, ok := indirect.([]byte); ok {
				record[key] = string(value)
			} else {
				record[key] = indirect
			}
		}

		records = append(records, record)
	}

	return records, nil

}

// SqliteInsert 适配在sqlite3上执行插入命令
func (this *Cli) SqliteInsert(filename string, table string, records []interface{}) (*sql.DB, error) {

	var (
		err error
		db  *sql.DB
	)

	if filename == "" {
		filename = ":memory:"
	}
	if db, err = sql.Open("sqlite3", filename); err != nil {
		return nil, err
	}

	Push := func(db *sql.DB, table string, records []interface{}) error {
		hashKeys := map[string]struct{}{}

		keyword := []string{"ALTER",
			"CLOSE",
			"COMMIT",
			"CREATE",
			"DECLARE",
			"DELETE",
			"DENY",
			"DESCRIBE",
			"DOMAIN",
			"DROP",
			"EXECUTE",
			"EXPLAN",
			"FETCH",
			"GRANT",
			"INDEX",
			"INSERT",
			"OPEN",
			"PREPARE",
			"PROCEDURE",
			"REVOKE",
			"ROLLBACK",
			"SCHEMA",
			"SELECT",
			"SET",
			"SQL",
			"TABLE",
			"TRANSACTION",
			"TRIGGER",
			"UPDATE",
			"VIEW",
			"GROUP"}
		_ = keyword
		for _, record := range records {
			switch record.(type) {
			case map[string]interface{}:
				for key, _ := range record.(map[string]interface{}) {
					if strings.HasPrefix(key, "`") {
						continue
					}
					key2 := fmt.Sprintf("`%s`", key)
					record.(map[string]interface{})[key2] = record.(map[string]interface{})[key]
					delete(record.(map[string]interface{}), key)
					hashKeys[key2] = struct{}{}
				}
			}
		}

		keys := []string{}

		for key, _ := range hashKeys {
			keys = append(keys, key)
		}

		//		db.Exec("DROP TABLE data")
		query := fmt.Sprintf("CREATE TABLE %s ("+strings.Join(keys, ",")+")", table)
		if _, err := db.Exec(query); err != nil {
			//fmt.Println(query)
			corelog.Error(err)
		}

		for _, record := range records {
			recordKeys := []string{}
			recordValues := []string{}
			recordArgs := []interface{}{}

			switch record.(type) {
			case map[string]interface{}:

				for key, value := range record.(map[string]interface{}) {
					recordKeys = append(recordKeys, key)
					recordValues = append(recordValues, "?")
					recordArgs = append(recordArgs, value)
				}

			}

			query := fmt.Sprintf("INSERT INTO %s ("+strings.Join(recordKeys, ",")+
				") VALUES ("+strings.Join(recordValues, ", ")+")", table)

			statement, err := db.Prepare(query)
			if err != nil {
				corelog.Error(
					err, "can't prepare query: %s", query,
				)
				continue

			}

			_, err = statement.Exec(recordArgs...)
			if err != nil {
				corelog.Error(
					err, "can't insert record",
				)

			}
			statement.Close()
		}

		return nil
	}

	err = Push(db, table, records)
	if err != nil {
		return nil, err
	}

	return db, err

}

// SqliteExec 适配在sqlite3上执行操作命令
func (this *Cli) SqliteExec(filename string, s string) (int64, error) {
	var (
		err    error
		db     *sql.DB
		result sql.Result
	)
	if filename == "" {
		filename = ":memory:"
	}
	if db, err = sql.Open("sqlite3", filename); err != nil {
		return -1, err
	}

	if result, err = db.Exec(s); err != nil {
		return -1, err
	}
	return result.RowsAffected()

}
