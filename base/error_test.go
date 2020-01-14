package base

import (
	"testing"
)

func TestErrorMessage(t *testing.T) {
	e := NewEpError(1, "error")

	if e.Error() != "error" {
		t.Errorf("错误信息不一致")
	}
}

func TestErrorType(t *testing.T) {
	e1 := &EpError{}
	var e2 interface{} = e1

	_, ok := e2.(error)
	if !ok {
		t.Errorf("EpError 没有实现error的方法")
	}
}
