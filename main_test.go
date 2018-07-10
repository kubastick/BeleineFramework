package main

import (
	"testing"
	"fmt"
)

func TestInput(t *testing.T) {
	input := MakeInput()
	input.SetHint("Username")
	if input.render() != fmt.Sprintf(`<input id="%s" type="text" class="form-control " placeholder="Username">`,input.GetInputId()) {t.Fail()}
	input.SetSize(2)
	if input.render() != fmt.Sprintf(`<input id="%s" type="text" class="form-control form-control-lg" placeholder="Username">`,input.GetInputId()) {t.Fail()}
	input.SetInputType("password")
	if input.render() != fmt.Sprintf(`<input id="%s" type="password" class="form-control form-control-lg" placeholder="Username">`,input.GetInputId()) {t.Fail()}
}

func TestButton(t *testing.T) {
	button := MakeButton()
	button.SetText("Click")
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-primary ">Click</button>`,button.GetBtnId()) {t.Fail()}
	button.SetSize(0)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-primary btn-sm">Click</button>`,button.GetBtnId()) {t.Fail()}
	button.SetButtonType(1)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-secondary btn-sm">Click</button>`,button.GetBtnId()) {t.Fail()}
	button.SetOutline(true)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-outline-secondary btn-sm">Click</button>`,button.GetBtnId()) {t.Fail()}
	button.SetState(false)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-outline-secondary btn-sm" disabled>Click</button>`,button.GetBtnId()) {t.Fail()}
}

func TestLabel(t *testing.T) {
	label := MakeLabel()
	label.SetText("Hello world")
	if label.render() != fmt.Sprintf(`<p id="%s">Hello world</p>`,label.GetLabelId()) {t.Fail()}
	label.SetSize(1)
	if label.render() != fmt.Sprintf(`<h1 id="%s">Hello world</h1>`,label.GetLabelId()) {t.Fail()}
	label.SetSize(6)
	if label.render() != fmt.Sprintf(`<h6 id="%s">Hello world</h6>`,label.GetLabelId()) {t.Fail()}
}

func TestAlert(t *testing.T) {
	alert := MakeAlert()
	alert.SetAlertType(0)
	alert.SetStrongText("Hello world")
	alert.SetText("This is alert description")

	if alert.render() != fmt.Sprintf(`
	<div id="%s" class=".alert-success">
  		<strong>Hello world</strong> This is alert description
	</div>
	`,alert.GetAlertId()) {t.Fail()}

	alert.SetAlertType(6)
	alert.SetStrongText("World hello")
	if alert.render() != fmt.Sprintf(`
	<div id="%s" class=".alert-light">
  		<strong>World hello</strong> This is alert description
	</div>
	`,alert.GetAlertId()) {t.Fail()}
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