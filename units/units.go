package units

import (
	"path"
	"runtime"
)

func Get_root_path() string {
	_, filename, _, _ := runtime.Caller(0)
	root_path := path.Dir(path.Dir(filename))
	return root_path
}
