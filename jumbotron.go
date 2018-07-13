package beleine

import "fmt"

type Jumbotron struct {
	id         string
	components []component
	fluid      bool
	js         string
}

//Create Jumbotron struct
func NewJumbotron() Jumbotron {
	return Jumbotron{id: getGlobalID()}
}

//Attach component to Jumbotron
func (j *Jumbotron) Attach(c component) {
	j.components = append(j.components, c)
}

func (j *Jumbotron) render() string {
	fluidTag := ""
	if j.fluid {
		fluidTag = " jumbotron-fluid"
	}
	result := fmt.Sprintf(`<div id="%s" class="jumbotron%s">`, j.id, fluidTag)

	for _, c := range j.components {
		result += c.render()
	}
	result += `</div>`
	return result
}

func (j *Jumbotron) renderJS() string {
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
