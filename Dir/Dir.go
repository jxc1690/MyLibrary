package Dir

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type file struct {
	Name    string      // base name of the file
	Size    int64       // length in bytes for regular files; system-dependent for others
	Mode    os.FileMode // file mode bits
	ModTime time.Time   // modification time
	IsDir   bool        // abbreviation for Mode().IsDir()
	Sys     any         // underlying data source (can return nil)
}

func ListFile(fileDir, 后缀名 string, 目录继续匹配 bool) (ret []file) {
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		return nil
	}
	for _, onefile := range files {
		if onefile.IsDir() {
			if 目录继续匹配 {
				for _, s := range listFile(fileDir+onefile.Name(), 后缀名) {
					ret = append(ret, s)
				}
			}
		} else {
			if 后缀名 == "*" {
				ret = append(ret, fftof(onefile))
				continue
			}
			ls := strings.Split(onefile.Name(), ".")
			if len(ls) > 1 {
				if ls[len(ls)-1] == 后缀名 {
					ret = append(ret, fftof(onefile))
				}
			}
		}
	}
	return
}
func listFile(fileDir, 后缀名 string) (ret []file) {
	fileDir += "/"
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		return nil
	}
	for _, onefile := range files {
		if onefile.IsDir() {
			for _, s := range listFile(fileDir+onefile.Name(), 后缀名) {
				ret = append(ret, s)
			}
		} else {
			if 后缀名 == "*" {
				ret = append(ret, fftof(onefile))
				continue
			}
			ls := strings.Split(onefile.Name(), ".")
			if len(ls) > 1 {
				if ls[len(ls)-1] == 后缀名 {
					s := fftof(onefile)
					s.Name = fileDir + s.Name
					ret = append(ret, s)
				}
			}
		}
	}
	return
}

func fftof(info os.FileInfo) file {
	return file{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
		Sys:     info.Sys(),
	}
}
