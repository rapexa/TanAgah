package stringResource

import (
	"github.com/gin-gonic/gin"
	"sync"
)

func GetStrings() Strings {
	return Getter().Strings
}

type StringGetter struct {
	Strings Strings
}

var lockLng = &sync.Mutex{}

var singleInstance *StringGetter

func Getter() *StringGetter {
	if singleInstance == nil {
		lockLng.Lock()
		defer lockLng.Unlock()
		if singleInstance == nil {
			singleInstance = &StringGetter{
				Strings: &ImplementationStrings{},
			}
		}
	}
	return singleInstance
}

func GetLng(c *gin.Context) string {
	return c.Request.Header.Get("language")
}
