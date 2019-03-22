package log

import (
	"fmt"
	"runtime"
)

func GetFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return fmt.Sprintf("%s:", f.Name())
}
