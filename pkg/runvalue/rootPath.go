package runvalue

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// RootPath 项目根目录
var (
	RootPath = func() string {
		notMain := false
		switch runtime.GOOS {
		case `linux`:
			if os.Args[0][len(os.Args[0])-5:] == `.test` {
				notMain = true
			}
		case `windows`:
			nowPath := filepath.Base(os.Args[0])
			if nowPath == `main.exe` {
				break
			}
			if nowPath[:7] == `___Test` || nowPath[len(nowPath)-9:] == `.test.exe` {
				notMain = true
			}
		}
		if notMain {
			_, onPath, _, _ := runtime.Caller(0)
			return path.Join(onPath, `..`, `..`, `..`)
		}
		return `.`
	}()
)
