package tracer

import (
	"fmt"
	"time"
)

func Debug(arg ...any) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Print(now + " >>> DEBUG >>> ")
	fmt.Println(arg...)
}

func Error(arg ...any) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Print(now + " >>> ERROR >>> ")
	fmt.Println(arg...)
}
