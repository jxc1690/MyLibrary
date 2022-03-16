package MyLibrary

//package MyLibrary

import (
	"github.com/axgle/mahonia"
	"log"
	"www.github.com/jxc1690/MyLibrary/install"
)

func main() {
	s, err := install.New("AAAA", "AAAAA", "测试用")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(s.Unistall())
	s.Run(func() {
		log.Println("测试")
	})
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
