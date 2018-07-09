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

func TestCore (t *testing.T) {
	var testPage Page

	helloworldLabel := MakeLabel()
	helloworldLabel.SetText("Hello world")
	helloworldLabel.SetSize(1) //H1
	testPage.Attach(&helloworldLabel)

	if testPage.Render() != `<h1>Hello world</h1>` {t.Fail()}

	helloworldLabel.SetText("Another text")

	if testPage.Render() != `<h1>Another text</h1>` {t.Fail()}
}

func BenchmarkPageHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var testPage Page

		helloworldLabel := MakeLabel()
		helloworldLabel.SetText("Hello world")
		helloworldLabel.SetSize(1) //H1
		testPage.Attach(&helloworldLabel)

		testPage.Render()
	}
}