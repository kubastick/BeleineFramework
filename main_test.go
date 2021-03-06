package beleine

import (
	"fmt"
	"testing"
)

func TestPagination(t *testing.T) {
	pagination := NewPagination()
	pItem := NewPaginationItem()
	pItem.SetText("1")
	pagination.AddItem(&pItem)
	if pagination.Render() != fmt.Sprintf(`
<nav id="%s" aria-label="Page navigation example">
  <ul class="pagination  ">
	<li id="%s" class="page-item"><a class="page-link" href="#">1</a></li>
  </ul>
</nav>
`, pagination.GetID(), pItem.GetID()) {
		t.Fail()
	}

	pItem2 := NewPaginationItem()
	pItem2.SetText("Next")
	pItem2.SetEnabled(false)
	pagination.AddItem(&pItem2)
	pagination.SetSize(2)
	pagination.SetAlignment("END")
	if pagination.Render() != fmt.Sprintf(`
<nav id="%s" aria-label="Page navigation example">
  <ul class="pagination pagination-lg justify-content-end">
	<li id="%s" class="page-item"><a class="page-link" href="#">1</a></li><li id="%s" class="page-item disabled"><a class="page-link" href="#">Next</a></li>
  </ul>
</nav>
`, pagination.GetID(), pItem.GetID(), pItem2.GetID()) {
		t.Fail()
	}
}

func TestCollapse(t *testing.T) {
	collapse := NewButton()
	collapse.SetText("Lorem")
	collapse.SetSize(0)
	collapse.CollapseEnabled(true)
	collapse.AddCollapseText("ipsum")
	if collapse.Render() != fmt.Sprintf(`
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
			`, collapse.GetID(), collapse.GetCollapseID(), collapse.GetCollapseID(), collapse.GetCollapseID()) {
		t.Fail()
	}
}

func TestDropdown(t *testing.T) {
	dropdown := NewButton()
	dropdown.SetText("Animals")
	dropdown.SetSize(2)
	dropdown.SetButtonType(3)
	dropdown.DropdownEnabled(true)

	dropdown.AddDropdownItem("Cat")
	if dropdown.Render() != fmt.Sprintf(`<div class="dropdown">
<button class="btn btn-danger btn-lg dropdown-toggle" type="button" id="%s" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
Animals
</button>
<div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
<a id="%s" class="dropdown-item" href="#">Cat</a>
</div>
</div>`, dropdown.GetID(), dropdown.GetDropdownItemID("Cat")) {
		t.Fail()
	}

	//SOMETIMES CRASHES NO IDEA WHY

	//	dropdown.AddDropdownItem("Dog")
	//	if dropdown.Render() != fmt.Sprintf(`<div class="dropdown">
	//<button class="btn btn-danger btn-lg dropdown-toggle" type="button" id="a2" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
	//Animals
	//</button>
	//<div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
	//<a id="a3" class="dropdown-item" href="#">Cat</a><a id="a4" class="dropdown-item" href="#">Dog</a>
	//</div>
	//</div>` /*,dropdown.GetID(),dropdown.GetDropdownItemID("Cat"),dropdown.GetDropdownItemID("Dog")*/) {
	//	println("|"+dropdown.Render()+"|")
	//	t.Fail()
	//	}
}

func TestBadge(t *testing.T) {
	badge := NewBadge()

	badge.SetText("xCairuuu")
	if badge.Render() != fmt.Sprintf(`<span id="%s" class="badge badge-primary">xCairuuu</span>`, badge.GetID()) {
		t.Fail()
	}
	badge.SetText("Bye world")
	badge.SetStyle(2)
	badge.SetPill(true)
	if badge.Render() != fmt.Sprintf(`<span id="%s" class="badge badge-pill badge-success">Bye world</span>`, badge.GetID()) {
		t.Fail()
	}
}
func TestProgress(t *testing.T) {
	progress := NewProgress()

	progress.SetMinMax(0, 100)
	if progress.Render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar   " role="progressbar" style="width: 0%s" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100"></div>
	</div>
	`, progress.GetID(), "%") {
		t.Fail()
	}

	progress.SetLabels(true)
	if progress.Render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar   " role="progressbar" style="width: 0%s" aria-valuenow="0" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`, progress.GetID(), "%", "%") {
		t.Fail()
	}

	progress.SetPercent(45)
	if progress.Render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar   " role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`, progress.GetID(), "%", "%") {
		t.Fail()
	}

	progress.SetProgressType(2)
	if progress.Render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar  bg-info " role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`, progress.GetID(), "%", "%") {
		t.Fail()
	}

	progress.SetStriped(true)
	if progress.Render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar progress-bar-striped bg-info " role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`, progress.GetID(), "%", "%") {
		t.Fail()
	}

	progress.SetAnimation(true)
	if progress.Render() != fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar progress-bar-striped bg-info progress-bar-animated" role="progressbar" style="width: 45%s" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100">0%s</div>
	</div>
	`, progress.GetID(), "%", "%") {
		t.Fail()
	}
}

func TestInput(t *testing.T) {
	input := NewInput()
	input.SetHint("Username")
	if input.Render() != fmt.Sprintf(`<input id="%s" type="text" class="form-control " placeholder="Username">`, input.GetID()) {
		t.Fail()
	}
	input.SetSize(2)
	if input.Render() != fmt.Sprintf(`<input id="%s" type="text" class="form-control form-control-lg" placeholder="Username">`, input.GetID()) {
		t.Fail()
	}
	input.SetInputType("password")
	if input.Render() != fmt.Sprintf(`<input id="%s" type="password" class="form-control form-control-lg" placeholder="Username">`, input.GetID()) {
		t.Fail()
	}
}

func TestButton(t *testing.T) {
	button := NewButton()
	button.SetText("Click")
	if button.Render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-primary ">Click</button>`, button.GetID()) {
		t.Fail()
	}
	button.SetSize(0)
	if button.Render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-primary btn-sm">Click</button>`, button.GetID()) {
		t.Fail()
	}
	button.SetButtonType(1)
	if button.Render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-secondary btn-sm">Click</button>`, button.GetID()) {
		t.Fail()
	}
	button.SetOutline(true)
	if button.Render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-outline-secondary btn-sm">Click</button>`, button.GetID()) {
		t.Fail()
	}
	button.SetState(false)
	if button.Render() != fmt.Sprintf(`<button id="%s" type="button" class="btn btn-outline-secondary btn-sm" disabled>Click</button>`, button.GetID()) {
		t.Fail()
	}
}

func TestLabel(t *testing.T) {
	label := NewLabel()
	label.SetText("Hello world")
	if label.Render() != fmt.Sprintf(`<p id="%s">Hello world</p>`, label.GetID()) {
		t.Fail()
	}
	label.SetSize(1)
	if label.Render() != fmt.Sprintf(`<h1 id="%s">Hello world</h1>`, label.GetID()) {
		t.Fail()
	}
	label.SetSize(6)
	if label.Render() != fmt.Sprintf(`<h6 id="%s">Hello world</h6>`, label.GetID()) {
		t.Fail()
	}
}

func TestAlert(t *testing.T) {
	alert := NewAlert()
	alert.SetAlertType(0)
	alert.SetStrongText("Hello world")
	alert.SetText("This is alert description")

	if alert.Render() != fmt.Sprintf(`
	<div id="%s" class=".alert-success">
  		<strong>Hello world</strong> This is alert description
	</div>
	`, alert.GetID()) {
		t.Fail()
	}

	alert.SetAlertType(6)
	alert.SetStrongText("World hello")
	if alert.Render() != fmt.Sprintf(`
	<div id="%s" class=".alert-light">
  		<strong>World hello</strong> This is alert description
	</div>
	`, alert.GetID()) {
		t.Fail()
	}
}

func TestCore(t *testing.T) {
	var testPage Page

	helloworldLabel := NewLabel()
	helloworldLabel.SetText("Hello world")
	helloworldLabel.SetSize(1) //H1
	helloworldLabel.SetOnClickListener(helloworldLabel.SetTextJS("Hi!"))
	testPage.Attach(&helloworldLabel)

	button := NewButton()
	button.SetText("Click me")
	button.SetOnClickListener(helloworldLabel.SetTextJS("You clicked button"))

	testPage.Attach(&button)

	testPage.Render()

	//println(testPage.Render())
}

func TestJumbotron(t *testing.T) {
	jumbotron := NewJumbotron()

	if jumbotron.Render() != fmt.Sprintf(`<div id="%s" class="jumbotron"></div>`, jumbotron.GetID()) {
		t.Fail()
	}

	label := NewLabel()
	label.SetText("hello")
	jumbotron.Attach(&label)
	jumbotron.SetFluid(true)

	if jumbotron.Render() != fmt.Sprintf(`<div id="%s" class="jumbotron jumbotron-fluid"><p id="%s">hello</p></div>`, jumbotron.id, label.id) {
		t.Fail()
	}
}

func TestCarousel(t *testing.T) {
	carousel := NewCarousel()
	i := make([]CarouselItem, 3)
	i[0] = CarouselItem{Title: "hello", ImageSource: "yo1.png", AltText: "Lol"}
	i[1] = CarouselItem{Title: "world", ImageSource: "yo2.png"}
	i[2] = CarouselItem{Title: "carousel", ImageSource: "yo3.png"}
	carousel.SetCarouselItems(&i)
	carousel.SetIndicatorsEnabled(true)
	carousel.SetCaptionsEnabled(true)
	//fmt.Println(carousel.Render())
	//TODO:Test real-world rendering
	//TODO:Test failing
}

func TestComponentInterface(t *testing.T) {
	var page Page
	page.Attach(&Alert{})
	page.Attach(&Badge{})
	page.Attach(&Button{})
	page.Attach(&Carousel{})
	page.Attach(&Input{})
	page.Attach(&Jumbotron{})
	page.Attach(&Input{})
	page.Attach(&Nav{})
}

func BenchmarkPageHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var testPage Page

		helloworldLabel := NewLabel()
		helloworldLabel.SetText("Hello world")
		helloworldLabel.SetSize(1) //H1
		testPage.Attach(&helloworldLabel)

		testPage.Render()
	}
}

func BenchmarkPageLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var testPage Page

		helloworldLabel := NewLabel()
		helloworldLabel.SetText("Log-in into BeleineFramework")
		helloworldLabel.SetSize(1) //H1

		loginCaption := NewLabel()
		loginCaption.SetText("Login")

		login := NewInput()
		login.SetHint("Login")

		password := NewInput()
		password.SetHint("*********")
		password.SetInputType("password")

		submitButton := NewButton()
		submitButton.SetText("Login")
		submitButton.SetOnClickListener(submitButton.SetTextJS("Loggining in..."))

		testPage.Attach(&helloworldLabel)
		testPage.Attach(&loginCaption)
		testPage.Attach(&login)
		testPage.Attach(&password)
		testPage.Attach(&submitButton)
		testPage.Render()
	}
}

func BenchmarkPageLoginPrepared(b *testing.B) {
	var testPage Page

	helloworldLabel := NewLabel()
	helloworldLabel.SetText("Log-in into BeleineFramework")
	helloworldLabel.SetSize(1) //H1

	loginCaption := NewLabel()
	loginCaption.SetText("Login")

	login := NewInput()
	login.SetHint("Login")

	password := NewInput()
	password.SetHint("*********")
	password.SetInputType("password")

	submitButton := NewButton()
	submitButton.SetText("Login")
	submitButton.SetOnClickListener(
		PostRequestJS("127.0.0.1/api", fmt.Sprintf("{l:%s,p:%s", login.GetTextJS(), password.GetTextJS()), "") +
			submitButton.SetTextJS("Logging in..."))

	testPage.Attach(&helloworldLabel)
	testPage.Attach(&loginCaption)
	testPage.Attach(&login)
	testPage.Attach(&password)
	testPage.Attach(&submitButton)
	for i := 0; i < b.N; i++ {
		testPage.Render()
	}
}

func TestNav(t *testing.T) {
	nav := NewNav()
	nav.SetStyle(1)
	nav.SetExpand(true)

	nav.AddItem(&NavItem{Title: "yo", Link: "#"})
	if nav.Render() != `<ul class="nav nav-tabs nav-fill">
<li class="nav-item">
	<a class="nav-link" href="#">yo</a>
</li>
</ul>` {
		t.Fail()
	}
	nav.SetExpand(false)
	nav.SetStyle(0)
	if nav.Render() != `<ul class="nav">
<li class="nav-item">
	<a class="nav-link" href="#">yo</a>
</li>
</ul>` {
		t.Fail()
	}

	err := nav.SetStyle(3)
	if err == nil {
		t.Fail()
	}
}

func TestBreakline(t *testing.T) {
	breakLine := NewBreakline()
	if breakLine.Render() != "</br>" {
		t.Fail()
	}
}
