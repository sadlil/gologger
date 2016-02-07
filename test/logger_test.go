package test

import (
	"testing"
	"github.com/sadlil/gologger"
)

func TestLogger(t *testing.T) {
	l := gologger.New("hello", gologger.Console)
	l.Info("a", 1, "hello")
}