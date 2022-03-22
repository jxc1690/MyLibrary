package MyLibrary

//package MyLibrary

import (
	"database/sql"
	"fmt"
	"github.com/axgle/mahonia"
	"log"
	"www.github.com/jxc1690/MyLibrary/mysql"
)

type s struct {
	ID int    `db:"ID"`
	IP string `db:"IP"`
}
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Queryx(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowx(query string, args ...interface{}) *sql.Row
}

func main() {
	Logplus()
	/*
		s, err := install.New("AAAA", "AAAAA", "测试用")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(s.Unistall())
		s.Run(func() {
			log.Println("测试")
		})
	*/
	sql, _ := mysql.New("127.0.0.1", 3306, "root", "a58230330", "data")
	var ss s
	sql.Db.Get(&ss, "select ID,IP from socks")
	fmt.Println(ss)
	m, _ := sql.GetMap("select * from socks")
	log.Println(m, m[0]["IP"])
}

func Logplus() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func GBKtoUTF_B(s []byte) string {
	return GBKtoUTF(string(s))
}

func GBKtoUTF(s string) string {
	return mahonia.NewDecoder("GBK").ConvertString(s)
}

func UTFtoGBk(s string) []byte {
	return []byte(mahonia.NewEncoder("GBK").ConvertString(s))
}
