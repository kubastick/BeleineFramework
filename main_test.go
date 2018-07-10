package beleine

import (
	"testing"
	"fmt"
)

func TestBadge(t *testing.T) {
	badge := MakeBadge()

	badge.SetText("xCairuuu")
	if badge.render() != fmt.Sprintf(`<span id="%s" class="badge badge-secondary">xCairuuu</span>`,badge.GetID()) {t.Fail()}

}
func TestProgress(t *testing.T) {
	progress :=MakeProgress()

	progress.SetMinMax(0,100)
	if progress.render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar   " role="progressbar" style="width: 0%s" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></div>
	</div>
	`,progress.GetID(), "%") {t.Fail()}

	progress.SetLabels(true)
	if progress.render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar   " role="progressbar" style="width: 0%s" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`,progress.GetID(), "%", "%") {t.Fail()}

	progress.SetPercent(45)
	if progress.render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar   " role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`,progress.GetID(), "%", "%") {t.Fail()}

	progress.SetProgressType(2)
	if progress.render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar  bg-info " role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`,progress.GetID(), "%", "%") {t.Fail()}

	progress.SetStriped(true)
	if progress.render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar progress-bar-striped bg-info " role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`,progress.GetID(), "%", "%") {t.Fail()}

	progress.SetAnimation(true)
	if progress.render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar progress-bar-striped bg-info progress-bar-animated" role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`,progress.GetID(), "%", "%") {t.Fail()}
}

func TestInput(t *testing.T) {
	input := MakeInput()
	input.SetHint("Username")
	if input.render() != fmt.Sprintf(`<input id="%s" type="text" class="form-control " placeholder="Username">`,input.GetID()) {t.Fail()}
	input.SetSize(2)
	if input.render() != fmt.Sprintf(`<input id="%s" type="text" class="form-control form-control-lg" placeholder="Username">`,input.GetID()) {t.Fail()}
	input.SetInputType("password")
	if input.render() != fmt.Sprintf(`<input id="%s" type="password" class="form-control form-control-lg" placeholder="Username">`,input.GetID()) {t.Fail()}
}

func TestButton(t *testing.T) {
	button := MakeButton()
	button.SetText("Click")
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-primary ">Click</button>`,button.GetID()) {t.Fail()}
	button.SetSize(0)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-primary btn-sm">Click</button>`,button.GetID()) {t.Fail()}
	button.SetButtonType(1)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-secondary btn-sm">Click</button>`,button.GetID()) {t.Fail()}
	button.SetOutline(true)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-outline-secondary btn-sm">Click</button>`,button.GetID()) {t.Fail()}
	button.SetState(false)
	if button.render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-outline-secondary btn-sm" disabled>Click</button>`,button.GetID()) {t.Fail()}
}

func TestLabel(t *testing.T) {
	label := MakeLabel()
	label.SetText("Hello world")
	if label.render() != fmt.Sprintf(`<p id="%s">Hello world</p>`,label.GetID()) {t.Fail()}
	label.SetSize(1)
	if label.render() != fmt.Sprintf(`<h1 id="%s">Hello world</h1>`,label.GetID()) {t.Fail()}
	label.SetSize(6)
	if label.render() != fmt.Sprintf(`<h6 id="%s">Hello world</h6>`,label.GetID()) {t.Fail()}
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
	`,alert.GetID()) {t.Fail()}

	alert.SetAlertType(6)
	alert.SetStrongText("World hello")
	if alert.render() != fmt.Sprintf(`
	<div id="%s" class=".alert-light">
  		<strong>World hello</strong> This is alert description
	</div>
	`,alert.GetID()) {t.Fail()}
}

func TestCore (t *testing.T) {
	var testPage Page

	helloworldLabel := MakeLabel()
	helloworldLabel.SetText("Hello world")
	helloworldLabel.SetSize(1) //H1
	helloworldLabel.SetOnClickListener(helloworldLabel.SetTextJS("Hi!"))
	testPage.Attach(&helloworldLabel)
	testPage.Render()

	println(testPage.Render())
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