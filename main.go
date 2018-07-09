//BeleineFramework

package main

//variables
var globalID int

type component interface {
	render() string
}

type Page struct {
	components []component
}

func (p *Page) Render() string{

	var result string
	for _,a := range p.components {
		result += a.render()
	}
	return result
}

func (p *Page) Attach(component component) {
	p.components = append(p.components,component)
}

///Get unique component ID
func getGlobalID() int {
	x := globalID
	globalID++
	return x
}
