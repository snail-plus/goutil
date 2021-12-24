package goutil_test

import (
	"fmt"
	"testing"

	"github.com/snail-plus/goutil"
	"github.com/stretchr/testify/assert"
)

func TestFuncName(t *testing.T) {
	name := goutil.FuncName(goutil.PkgName)
	assert.Equal(t, "github.com/snail-plus/goutil.PkgName", name)

	name = goutil.FuncName(goutil.PanicIfErr)
	assert.Equal(t, "github.com/snail-plus/goutil.PanicIfErr", name)
}

func TestPkgName(t *testing.T) {
	name := goutil.PkgName(goutil.FuncName(goutil.PanicIfErr))
	assert.Equal(t, "github.com/snail-plus/goutil", name)
}

func TestPanicIfErr(t *testing.T) {
	goutil.PanicIfErr(nil)
}

func TestPanicf(t *testing.T) {
	assert.Panics(t, func() {
		goutil.Panicf("hi %s", "inhere")
	})
}

func TestGetCallStacks(t *testing.T) {
	msg := goutil.GetCallStacks(false)
	fmt.Println(string(msg))

	fmt.Println("-------------full stacks-------------")
	msg = goutil.GetCallStacks(true)
	fmt.Println(string(msg))
}
