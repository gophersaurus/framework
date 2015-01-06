package mockstar

import (
	"reflect"
	"testing"
)

var T *testing.T

func NewMock() *Mock {
	return &Mock{map[string][]expectedCall{}}
}

type Mock struct {
	expectedCalls map[string][]expectedCall
}

func (m *Mock) When(methodName string, args ...interface{}) *mockResult {
	call := newExpectedCall(args)
	calls, ok := m.expectedCalls[methodName]
	if ok {
		m.expectedCalls[methodName] = append(calls, *call)
	} else {
		m.expectedCalls[methodName] = []expectedCall{*call}
	}
	return call.result
}

func (m *Mock) CalledVarArgs(methodName string, args ...interface{}) *Args {
	size := len(args)
	if size <= 0 {
		return m.Called(methodName)
	}
	lastIndex := size - 1
	lastElem := args[lastIndex]
	val := reflect.ValueOf(lastElem)
	kind := val.Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		length := val.Len()
		newArgs := args[:lastIndex]
		for index := 0; index < length; index++ {
			newArgs = append(newArgs, val.Index(index).Interface())
		}
		return m.Called(methodName, newArgs...)
	} else {
		return m.Called(methodName, args...)
	}
}

func (m *Mock) Called(methodName string, args ...interface{}) *Args {
	call := m.GetCall(methodName, args...)
	argsObj := Args(args)
	results, ok := call.result.call(&argsObj)
	if !ok {
		Fatal("Method Call exheeds expected count: " + methodName)
	}
	return results
}

func (m *Mock) GetCall(methodName string, args ...interface{}) *expectedCall {
	callList, ok := m.expectedCalls[methodName]
	if !ok {
		Fatal("Unexpected Method Call: " + methodName)
	}
	for _, call := range callList {
		if len(*call.params) > len(args) {
			Fatal("invalid arguments included in mock: " + methodName)
		}
		check := true
		for index, param := range *call.params {
			check = check && (reflect.DeepEqual(param, Any) || reflect.DeepEqual(param, args[index]))
		}
		if check {
			return &call
		}
	}
	Fatal("call not found: " + methodName)
	return nil
}

func (m *Mock) HasCalled(methodName string, args ...interface{}) *callCounter {
	call := m.GetCall(methodName, args...)
	result := call.result
	counter := callCounter{len(result.callArgs), result.expectedCallCount}
	return &counter
}

type callCounter struct {
	count    int
	expected int
}

func (c *callCounter) Once() bool {
	return c.count == 1
}

func (c *callCounter) Expected() bool {
	return c.count == c.expected
}

func (c *callCounter) Twice() bool {
	return c.count == 2
}

func (c *callCounter) Times(count int) bool {
	return c.count == count
}

func (c *callCounter) AtLeast(count int) bool {
	return c.count >= count
}

type expectedCall struct {
	params *Args
	result *mockResult
}

func newExpectedCall(args []interface{}) *expectedCall {
	params := Args(args)
	return &expectedCall{&params, newMockResult()}
}

func (e *expectedCall) GetCallCount() int {
	return len(e.result.callArgs)
}

func (e *expectedCall) GetCallParams(index int) *Args {
	args := e.result.callArgs[index]
	return &args
}

type mockResult struct {
	resultSet         []func(args *Args) *Args
	callArgs          []Args
	expectedCallCount int
}

func newMockResult() *mockResult {
	return &mockResult{[]func(args *Args) *Args{}, []Args{}, -1}
}

func (m *mockResult) Then(do func(args *Args) *Args) *mockResult {
	m.resultSet = append(m.resultSet, do)
	return m
}

func (m *mockResult) Return(returnArgs ...interface{}) *mockResult {
	return m.Then(func(args *Args) *Args {
		out := Args(returnArgs)
		return &out
	})
}

func (m *mockResult) Panic(body interface{}) *mockResult {
	return m.Then(func(args *Args) *Args {
		panic(body)
		return nil
	})
}

func (m *mockResult) Expect(calls int) *mockResult {
	m.expectedCallCount = calls
	return m
}

func (m *mockResult) call(in *Args) (*Args, bool) {
	m.callArgs = append(m.callArgs, *in)
	size := len(m.resultSet)
	resultIndex := len(m.callArgs) - 1
	if size <= 0 {
		return nil, false
	}
	index := resultIndex % size
	fn := m.resultSet[index]
	args := fn(in)
	return args, (resultIndex < m.expectedCallCount || m.expectedCallCount < 0)
}

type Args []interface{}

func (a *Args) Get(index int) interface{} {
	size := len(*a)
	if index >= size || index < 0 {
		return nil
	}
	out := (*a)[index]
	if reflect.DeepEqual(out, Any) {
		return nil
	}
	return out
}

func (a *Args) ValueOf(index int) reflect.Value {
	return reflect.ValueOf(a.Get(index))
}

func (a *Args) TypeOf(index int) reflect.Type {
	return reflect.TypeOf(a.Get(index))
}

type any string

const Any any = "any value" // TODO -- better way to do this?
