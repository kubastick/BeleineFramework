package main

import (
	"testing"
)

func TestLabel(t *testing.T) {
	label := MakeLabel()
	label.SetText("Hello world")
	if label.render() != `<p>Hello world</p>` {t.Fail()}
	label.SetSize(1)
	if label.render() != `<h1>Hello world</h1>` {t.Fail()}
	label.SetSize(6)
	if label.render() != `<h6>Hello world</h6>` {t.Fail()}
}
