package test

import (
	"testing"
	"worframe/share/initialize"
)

func TestKafkaSendMsg(t *testing.T) {
	_ = initialize.SendMsg()
	t.Skip()
}
