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

func TestAlert(t *testing.T) {
	alert := MakeAlert()
	alert.SetAlertType(0)
	alert.SetStrongText("Hello world")
	alert.SetText("This is alert description")

	if alert.render() != `
	<div class=".alert-success">
  		<strong>Hello world</strong> This is alert description
	</div>
	` {t.Fail()}

	alert.SetAlertType(6)
	alert.SetStrongText("World hello")
	if alert.render() != `
	<div class=".alert-light">
  		<strong>World hello</strong> This is alert description
	</div>
	` {t.Fail()}
}

func TestCore (t *testing.T) {
	var testPage Page

	helloworldLabel := MakeLabel()
	helloworldLabel.SetText("Hello world")
	helloworldLabel.SetSize(1) //H1
	testPage.Attach(&helloworldLabel)
	//fmt.Println(testPage.Render())
	testPage.Render()
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