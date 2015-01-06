package mockstar

import (
	"fmt"
	"reflect"
	"runtime"
)

type expectation struct {
	actualValue interface{}
}

func Expect(actual interface{}) *expectation {
	return &expectation{actual}
}

func (e *expectation) ToEqual(expected interface{}) {
	if !reflect.DeepEqual(expected, e.actualValue) {
		Fatal(fmt.Sprintf("assertion failed:\n\texpected:(%v)\n\tactual:(%v)\n", expected, e.actualValue))
	}
}

func (e *expectation) ToBeTrue() {
	e.ToEqual(true)
}

func (e *expectation) ToBeNil() {
	e.ToEqual(nil)
}

func Fatal(message string) {
	stack := make([]byte, 10000)
	runtime.Stack(stack, true)
	T.Fatal(message + "\n" + string(stack))
}
