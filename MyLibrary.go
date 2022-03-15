package MyLibrary

import (
	"github.com/axgle/mahonia"
	"log"
)

func Logplus() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("开启成功")
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
