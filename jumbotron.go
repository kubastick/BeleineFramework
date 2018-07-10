package beleine

import "fmt"

type Jumbotron struct {
	id string
	components []component
	fluid bool
}

func MakeJumbotron() Jumbotron {
	return Jumbotron{id:getGlobalID()}
}

func (j *Jumbotron) Attach(c component) {
	j.components = append(j.components,c)
}

func (j *Jumbotron) render() string {
	fluidTag := ""
	if j.fluid {fluidTag=" jumbotron-fluid"}
	result := fmt.Sprintf(`<div id="%s" class="jumbotron%s">`,j.id,fluidTag)

	for _,c := range j.components {
		result += c.render()
	}
	result += `</div>`
	return result
}

func (j *Jumbotron) renderJS() string {
	return ""
}

func (j *Jumbotron) SetFluid (fluid bool) {
	j.fluid = fluid
}

func (j *Jumbotron) GetID() string {
	return j.id
}

