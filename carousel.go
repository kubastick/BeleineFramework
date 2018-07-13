package beleine

import "fmt"

//Carousel
//TODO:Finish rendering captions
//TODO:Add JS functions, and api
type Carousel struct {
	id string
	components *[]CarouselItem
	controls bool
	indicators bool
	captions bool
	crossFade bool
	js string
}

type CarouselItem struct {
	ImageSource string
	AltText string
	Title string
	Caption string
}

//Create Carousel struct
func NewCarousel() Carousel {
	return Carousel{id:getGlobalID()}
}

//Return ID of the carousel
func (c *Carousel) GetID() string {
	return c.id
}

//Set Carousel items
func (c *Carousel) SetCarouselItems(items *[]CarouselItem) {
	c.components = items
}

//Enable/disable Carousel controls
func (c *Carousel) SetControlsEnabled(enabled bool) {
	c.controls = enabled
}

//Enable/disable Carousel indicators
func (c *Carousel) SetIndicatorsEnabled(enabled bool) {
	c.indicators = enabled
}

//Enable/disable Carousel caption text
func (c *Carousel) SetCaptionsEnabled(enabled bool) {
	c.captions = enabled
}

//Returns js code that sets interval in ms
func (c *Carousel) SetIntervalJS (interval int) string {
	return fmt.Sprintf(`$('%s').carousel({interval: %d})`,c.id,interval)
}

//Sets interval of carousel rotating
func (c *Carousel) SetInterval (interval int) {
	c.js += c.SetIntervalJS(interval) +";"
}

//Sets JS code function, to be executed after click
func (c *Carousel) SetOnClickListener(listener string) {
	c.js += fmt.Sprintf("%s.onclick = function(){%s};",c.id,listener)
}

func(c *Carousel) render() string {
	result := fmt.Sprintf(`
<div id="%s" class="carousel slide" data-ride="carousel">
<div class="carousel-inner">
    `,c.id)
    if c.indicators {
    	first := ` class="active"`
    	result += `<ol class="carousel-indicators">`

    	for i:=0;i<len(*c.components);i++ {
			result += fmt.Sprintf(`<li data-target="#%s" data-slide-to="0"%s"></li>`,c.id,first)
			if first != "" {
				first = ""
			}
		}
	}
	first := " active"
    for _,i := range *c.components {
    	result += fmt.Sprintf(`<div class="carousel-item%s">`,first)
    	result += fmt.Sprintf(` <img class="d-block w-100" src="%s" alt="%s">`,i.ImageSource,i.AltText)
    	if c.captions {
    		result += fmt.Sprintf(`
		<div class="carousel-caption d-none d-md-block">
    	<h5>%s</h5>
    	<p>%s</p>
  		</div>`,i.Title,i.Caption)
		}
    	result += `</div>`
    	if first != "" {
			first = ""
		}
	}

	if c.controls {
		result += `<a class="carousel-control-prev" href="#carouselExampleControls" role="button" data-slide="prev">
    <span class="carousel-control-prev-icon" aria-hidden="true"></span>
    <span class="sr-only">Previous</span>
  </a>
  <a class="carousel-control-next" href="#carouselExampleControls" role="button" data-slide="next">
    <span class="carousel-control-next-icon" aria-hidden="true"></span>
    <span class="sr-only">Next</span>
  </a>`
	}
	return result
}

func (c *Carousel) renderJS() string {
	return c.js
}
