package utils

import (
	"time"
)

func Now()(string) {
	return time.Now().String()
}
func MustThrowError()(){
	panic("gogogogo")
	
}
