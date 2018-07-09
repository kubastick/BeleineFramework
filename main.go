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

func (p Page) Render() string{

	var result string
	for _,a := range p.components {
		result += a.render()
	}
	return result
}

///Get unique component ID
func getGlobalID() int {
	x := globalID
	globalID++
	return x
}
