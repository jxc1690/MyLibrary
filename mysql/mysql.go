package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type sql struct {
	Db *sqlx.DB
}
type nsql struct {
	IP, PORT, PASSNAME, PASSWORD, SQLNAME string
}

func New(Ip string, Port int, Passname, Password, Sqlname string) (*sql, error) {
	sqlstring := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Passname, Password, Ip, Port, Sqlname)
	db, err := sqlx.Open("mysql", sqlstring)
	if err != nil {
		return nil, err
	}
	return &sql{Db: db}, nil
}
func Newstr(Ip, Port, Passname, Password, Sqlname string) (*sql, error) {
	sqlstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Passname, Password, Ip, Port, Sqlname)
	db, err := sqlx.Open("mysql", sqlstring)
	if err != nil {
		return nil, err
	}
	return &sql{Db: db}, nil
}

func News(sqls nsql) (*sql, error) {
	sqlstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sqls.IP, sqls.PORT, sqls.PASSNAME, sqls.PASSWORD, sqls.SQLNAME)
	db, err := sqlx.Open("mysql", sqlstring)
	if err != nil {
		return nil, err
	}
	return &sql{Db: db}, nil
}

// GetMap Get using this sql.DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (p sql) GetMap(query string) ([]map[string]interface{}, error) {
	stmt, err := p.Db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	// 查询
	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	// 数据列
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 列的个数
	count := len(columns)

	// 返回值 Map切片
	mData := make([]map[string]interface{}, 0)
	// 一条数据的各列的值（需要指定长度为列的个数，以便获取地址）
	values := make([]interface{}, count)
	// 一条数据的各列的值的地址
	valPointers := make([]interface{}, count)
	for rows.Next() {

		// 获取各列的值的地址
		for i := 0; i < count; i++ {
			valPointers[i] = &values[i]
		}

		// 获取各列的值，放到对应的地址中
		rows.Scan(valPointers...)

		// 一条数据的Map (列名和值的键值对)
		entry := make(map[string]interface{})

		// Map 赋值
		for i, col := range columns {
			var v interface{}

			// 值复制给val(所以Scan时指定的地址可重复使用)
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				// 字符切片转为字符串
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}

		mData = append(mData, entry)
	}

	return mData, nil
}
