// /这个要改
package log

import (
	"fmt"
		"runtime/debug"
)

func Info(msg interface{}, parameters ...interface{}) {
	zaploger.Info(fmt.Sprintf(fmt.Sprint(msg), parameters...))
}

func Warn(msg interface{}, parameters ...interface{}) {
	zaploger.Warn(fmt.Sprintf(fmt.Sprint(msg), parameters...))
}

func Error(msg interface{}, parameters ...interface{}) {
	zaploger.Error(fmt.Sprintf(fmt.Sprint(msg), parameters...))
}

func Debug(msg interface{}, parameters ...interface{}) {
	zaploger.Debug(fmt.Sprintf(fmt.Sprint(msg), parameters...))
}

func DPanic(msg interface{}, parameters ...interface{}) {
	zaploger.DPanic(fmt.Sprintf(fmt.Sprint(msg), parameters...))
}

func Panic(msg interface{}, parameters ...interface{}) {
	zaploger.Panic(fmt.Sprintf(fmt.Sprint(msg), parameters...))
}

func Fatal(msg interface{}, parameters ...interface{}) {
	zaploger.Fatal(fmt.Sprintf(fmt.Sprint(msg), parameters...))
}

///会抛出异常
func Recover(msg interface{}, parameters ...interface{}) {
	err:=recover()
	if err != nil{
		zaploger.Panic(fmt.Sprintf(fmt.Sprint(msg), parameters...)+ " debug.Stack: " +string(debug.Stack()))
	}
	 
}

///不会抛出异常
func DRecover(msg interface{}, parameters ...interface{}) {
	err:=recover()
	if err != nil{
		zaploger.DPanic(fmt.Sprintf(fmt.Sprint(msg), parameters...)+ " debug.Stack: " +string(debug.Stack()))
	}
	 
}

