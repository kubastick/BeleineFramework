//BeleineFramework

package main

import "fmt"

//variables
var globalID int

type component interface {
	render() string
}

type Page struct {
	components []component
	title string
}

func (p *Page) Render() string{
	//Header
	result := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
  		<title>%s</title>
  		<meta charset="utf-8">
  		<meta name="viewport" content="width=device-width, initial-scale=1">
  		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css">
  		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  		<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.0/umd/popper.min.js"></script>
  		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.1.0/js/bootstrap.min.js"></script>
		</head>
	<body>
	`,p.title)

	for _,a := range p.components {
		result += a.render()
	}

	//Footer
	result += `</body></html>`
	return result
}

func (p *Page) Attach(component component) {
	p.components = append(p.components,component)
}

func (p *Page) SetTitle(title string) {
	p.title = title
}

///Get unique component ID
func getGlobalID() int {
	x := globalID
	globalID++
	return x
}
