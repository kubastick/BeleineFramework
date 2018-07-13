package beleine

import "fmt"

type Jumbotron struct {
	id         string
	components []Component
	fluid      bool
	js         string
}

//Create Jumbotron struct
func NewJumbotron() Jumbotron {
	return Jumbotron{id: GetGlobalID()}
}

//Attach component to Jumbotron
func (j *Jumbotron) Attach(c interface{Component}) {
	j.components = append(j.components, c)
}

func (j *Jumbotron) Render() string {
	fluidTag := ""
	if j.fluid {
		fluidTag = " jumbotron-fluid"
	}
	result := fmt.Sprintf(`<div id="%s" class="jumbotron%s">`, j.id, fluidTag)

	for _, c := range j.components {
		result += c.Render()
	}
	result += `</div>`
	return result
}

func (j *Jumbotron) RenderJS() string {
	return j.js
}

//Enable/disable Jumbotron fluid
func (j *Jumbotron) SetFluid(fluid bool) {
	j.fluid = fluid
}

//Returns component ID
func (j *Jumbotron) GetID() string {
	return j.id
}

//Sets JS code function, to be executed after click
func (j *Jumbotron) SetOnClickListener(listener string) {
	j.js += fmt.Sprintf("%s.onclick = function(){%s};", j.id, listener)
}
