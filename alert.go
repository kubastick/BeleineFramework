package beleine

//ALERT

import (
	"fmt"
)

/*
Alert types:
	ALERT_SUCCES = 0
	ALERT_IFNO = 1
	ALERT_WARNING = 2
	ALERT_DANGER = 3
	ALERT_PRIMARY = 4
	ALERT_SECONDARY = 5
	ALERT_LIGHT = 6
	ALERT_DARK = 7
*/

type Alert struct {
	id string
	strongText string
	text string
	alertType int
	js string
	//TODO closable bool
}

func MakeAlert() Alert {
	return Alert{id:getGlobalID()}
}

func (a *Alert) GetAlertId() string {
	return a.id
}

func (a *Alert) SetText(text string) {
	a.text = text
}

func (a *Alert) SetStrongText(text string) {
	a.strongText = text
}

func (a *Alert) SetAlertType(alertType int) {
	a.alertType = alertType
}

func (a *Alert) render() string {
	return fmt.Sprintf(`
	<div id="%s" class="%s">
  		<strong>%s</strong> %s
	</div>
	`, a.id, a.getDivClassName(a.alertType),a.strongText,a.text)
}

func (a *Alert) renderJS() string {
	return a.js
}

func (a *Alert) getDivClassName(num int) string {
	switch num {
	case 0:
		return ".alert-success"
	case 1:
		return ".alert-info"
	case 2:
		return ".alert-warning"
	case 3:
		return ".alert-danger"
	case 4:
		return ".alert-primary"
	case 5:
		return ".alert-secondary"
	case 6:
		return ".alert-light"
	case 7:
		return ".alert-dark"
	}
	panic("illegal alert type number")
}


