package test

import (
	"lbf.com/go-dp/prototype_dp"
	"testing"
)

func TestPrototype(t *testing.T) {

	hanli := prototype_dp.CreateImmortal("韩立", 9999)

	clonedHanli := hanli.Clone()

	t.Log(hanli)
	t.Log(clonedHanli)

	t.Log(hanli == clonedHanli)

}
