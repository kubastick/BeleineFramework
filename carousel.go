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
}

type CarouselItem struct {
	ImageSource string
	AltText string
	Title string
	Caption string
}

func (c *Carousel) SetCarouselItems(items *[]CarouselItem) {
	c.components = items
}

func (c *Carousel) SetControlsEnabled(enabled bool) {
	c.controls = enabled
}

func (c *Carousel) SetIndicatorsEnabled(enabled bool) {
	c.indicators = enabled
}

func (c *Carousel) SetCaptionsEnabled(enabled bool) {
	c.captions = enabled
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
	return ""
}
