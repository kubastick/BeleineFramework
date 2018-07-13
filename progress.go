package beleine

//PROGRESS

import (
	"errors"
	"fmt"
)

/*
Button types:
	DEFAULT = 0
	SUCCES = 1
	IFNO = 2
	WARNING = 3
	DANGER = 4
*/

type Progress struct {
	id           string
	max          int
	min          int
	showLabels   string
	percent      int
	progressType string
	striped      string
	animated     string
	js           string
}

//Creates new Progress struct
func NewProgress() Progress {
	return Progress{id: getGlobalID()}
}

//Returns ID of the progress
func (p *Progress) GetID() string {
	return p.id
}

//Sets min and max value of the progress bar
func (p *Progress) SetMinMax(min, max int) {
	p.min = min
	p.max = max
}

//Sets label (default false)
func (p *Progress) SetLabels(showLabels bool) {
	if showLabels {
		p.showLabels = fmt.Sprintf("%d%s", p.percent, "%")
	} else {
		p.showLabels = ""
	}
}

//Sets percent of now value of progress
func (p *Progress) SetPercent(percent int) {
	p.percent = percent
}

//Sets progress bar type (default 0)
func (p *Progress) SetProgressType(progressType int) error {
	switch progressType {
	case 0:
		p.progressType = ""
	case 1:
		p.progressType = "bg-success"
	case 2:
		p.progressType = "bg-info"
	case 3:
		p.progressType = "bg-warning"
	case 4:
		p.progressType = "bg-danger"
	default:
		return errors.New("Type does not exist")
	}
	return nil
}

//Sets striped progress bar (default false)
func (p *Progress) SetStriped(striped bool) {
	if striped {
		p.striped = "progress-bar-striped"
	} else {
		p.striped = ""
	}
}

//Sets animated progress bar (default false)
func (p *Progress) SetAnimation(animated bool) {
	if animated {
		p.animated = "progress-bar-animated"
	} else {
		p.animated = ""
	}
}

//Sets JS code function, to be executed after click
func (p *Progress) SetOnClickListener(listener string) {
	p.js += fmt.Sprintf("%s.onclick = function(){%s}", p.id, listener)
}

func (p *Progress) render() string {
	return fmt.Sprintf(`
	<div id="%s" class="progress">
  		<div class="progress-bar %s %s %s" role="progressbar" style="width: %d%s" aria-valuenow="%d" aria-valuemin="%d" aria-valuemax="%d">%s</div>
	</div>
	`, p.id, p.striped, p.progressType, p.animated, p.percent, "%", p.percent, p.min, p.max, p.showLabels)
}

func (p *Progress) renderJS() string {
	return p.js
}
