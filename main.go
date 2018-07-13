/*

 ____            ___                               ____                                                               __
/\  _`\         /\_ \           __                /\  _`\                                                            /\ \
\ \ \L\ \     __\//\ \      __ /\_\    ___      __\ \ \L\_\_ __    __      ___ ___      __   __  __  __    ___   _ __\ \ \/'\
 \ \  _ <'  /'__`\\ \ \   /'__`\/\ \ /' _ `\  /'__`\ \  _\/\`'__\/'__`\  /' __` __`\  /'__`\/\ \/\ \/\ \  / __`\/\`'__\ \ , <
  \ \ \L\ \/\  __/ \_\ \_/\  __/\ \ \/\ \/\ \/\  __/\ \ \/\ \ \//\ \L\.\_/\ \/\ \/\ \/\  __/\ \ \_/ \_/ \/\ \L\ \ \ \/ \ \ \\`\
   \ \____/\ \____\/\____\ \____\\ \_\ \_\ \_\ \____\\ \_\ \ \_\\ \__/.\_\ \_\ \_\ \_\ \____\\ \___x___/'\ \____/\ \_\  \ \_\ \_\
    \/___/  \/____/\/____/\/____/ \/_/\/_/\/_/\/____/ \/_/  \/_/ \/__/\/_/\/_/\/_/\/_/\/____/ \/__//__/   \/___/  \/_/   \/_/\/_/


*/

package beleine

import (
	"fmt"
	"strconv"
)

//variables
var globalID int

type Component interface {
	Render() string
	RenderJS() string
	GetID() string
}

//Root Page struct
type Page struct {
	components []Component
	title      string
	js string
}

//Return complete page as string
func (p *Page) Render() string {
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
	<div class="container">
	`, p.title)

	//HTML
	for _, a := range p.components {
		result += a.Render()
	}

	//Javascript
	result += `<script>`
	for _, a := range p.components {
		if a.RenderJS() != "" {
			result += a.RenderJS() + ";"
		}
	}
	result += p.js
	result += `</script>`

	//Footer
	result += `</div></body></html>`
	return result
}

//Adds component to page
func (p *Page) Attach(component interface{Component}) {
	p.components = append(p.components, component)
}

//Sets "<title>" attribute of page
func (p *Page) SetTitle(title string) {
	p.title = title
}

//Attach Javascript code
func (p *Page) AttachJSCode(code string) {
	p.js += code
}

//Get unique component ID
func GetGlobalID() string {
	x := globalID
	globalID++
	return "a" + strconv.Itoa(x)
}
