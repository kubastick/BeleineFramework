package beleine

import (
	"testing"
	"fmt"
)

func TestCollapse(t *testing.T) {
	collapse := MakeButton()
	collapse.SetText("Lorem")
	collapse.SetSize(0)
	collapse.CollapseEnabled(true)
	collapse.AddCollapseText("ipsum")
	if collapse.render() != fmt.Sprintf(`
<p>
  <button id="%s" class="btn btn-primary btn-sm" type="button" data-toggle="collapse" data-target="#%s" aria-expanded="false" aria-controls="%s">
	Lorem
  </button>
</p>
<div class="collapse" id="%s">
  <div class="card card-body">
    ipsum
  </div>
</div>
			`,collapse.GetID(),collapse.GetCollapseID(),collapse.GetCollapseID(),collapse.GetCollapseID()) {t.Fail()}
}

func TestDropdown(t *testing.T) {
	dropdown := MakeButton()
	dropdown.SetText("Animals")
	dropdown.SetSize(2)
	dropdown.SetButtonType(3)
	dropdown.DropdownEnabled(true)

	dropdown.AddDropdownItem("Cat")
	if dropdown.render() != fmt.Sprintf(`
<div class="dropdown">
  <button class="btn btn-danger btn-lg dropdown-toggle" type="button" id="%s" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
    Animals
  </button>
  <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
	<a id="%s" class="dropdown-item" href="#">Cat</a>
</div>
</div>
			`,dropdown.GetID(),dropdown.GetDropdownItemID("Cat")) {t.Fail()}

	dropdown.AddDropdownItem("Dog")
	if dropdown.render() != fmt.Sprintf(`
<div class="dropdown">
  <button class="btn btn-danger btn-lg dropdown-toggle" type="button" id="%s" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
    Animals
  </button>
  <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
	<a id="%s" class="dropdown-item" href="#">Cat</a>
	<a id="%s" class="dropdown-item" href="#">Dog</a>
</div>
</div>
			`,dropdown.GetID(),dropdown.GetDropdownItemID("Cat"),dropdown.GetDropdownItemID("Dog")) {t.Fail()}
}

func TestBadge(t *testing.T) {
	badge := MakeBadge()

	badge.SetText("xCairuuu")
	if badge.render() != fmt.Sprintf(`<span id="%s" class="badge badge-primary">xCairuuu</span>`,badge.GetID()) {t.Fail()}
	badge.SetText("Bye world")
	badge.SetStyle(2)
	badge.SetPill(true)
	if badge.render() != fmt.Sprintf(`<span id="%s" class="badge badge-pill badge-success">Bye world</span>`,badge.GetID()) {t.Fail()}
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

	button := MakeButton()
	button.SetText("Click me")
	button.SetOnClickListener(helloworldLabel.SetTextJS("You clicked button"))

	testPage.Attach(&button)

	testPage.Render()

	//println(testPage.Render())
}

func TestJumbotron(t *testing.T){
	jumbotron := MakeJumbotron()

	if jumbotron.render() != fmt.Sprintf(`<div id="%s" class="jumbotron"></div>`,jumbotron.GetID()) {t.Fail()}

	label := MakeLabel()
	label.SetText("hello")
	jumbotron.Attach(&label)
	jumbotron.SetFluid(true)

	if jumbotron.render() !=  fmt.Sprintf(`<div id="%s" class="jumbotron jumbotron-fluid"><p id="%s">hello</p></div>`,jumbotron.id,label.id) {t.Fail()}
}

func TestCarousel(t *testing.T){
	carousel := MakeCarousel()
	i := make([]CarouselItem,3)
	i[0] = CarouselItem{Title:"hello",ImageSource:"yo1.png",AltText:"Lol"}
	i[1] = CarouselItem{Title:"world",ImageSource:"yo2.png"}
	i[2] = CarouselItem{Title:"carousel",ImageSource:"yo3.png"}
	carousel.SetCarouselItems(&i)
	carousel.SetIndicatorsEnabled(true)
	carousel.SetCaptionsEnabled(true)
	//fmt.Println(carousel.render())
	//TODO:Test real-world rendering
	//TODO:Test failing
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