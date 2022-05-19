package errorhandler

import (
	"sync"
	"time"

	"bitbucket.org/test-webcrawl/utils"
	"go.uber.org/zap"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

var (
	dumpFileMutex sync.Mutex
	timezone, _   = time.LoadLocation("Asia/Kolkata")
)

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func LogError(ex Exception) {
	utils.LogD.Error("Error", zap.Any("", ex))
}
