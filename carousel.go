package beleine

import "fmt"

//Carousel

type Carousel struct {
	id string
	components []CarouselItem
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

func(c *Carousel) render() string {
	result := `
<div id="carouselExampleSlidesOnly" class="carousel slide" data-ride="carousel">
<div class="carousel-inner">
    `
    for _,i := range c.components {
    	result += `<div class="carousel-item">`
    	result += fmt.Sprintf(` <img class="d-block w-100" src="%s" alt="%s">`,i.ImageSource,i.AltText)
	}
}
