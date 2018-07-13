# Beleine Framework
[![Build Status](https://travis-ci.com/kubastick/BeleineFramework.svg?token=bpH3SmbkspnyvKjSgrmq&branch=master)](https://travis-ci.com/kubastick/BeleineFramework)
[![GoDoc](https://godoc.org/github.com/kubastick/BeleineFramework?status.svg)](https://godoc.org/github.com/kubastick/BeleineFramework)  
Tried of html,css and javascript ?  
Beleine Framework is the fastest way to build web applications in Go  
It gives you high level api, and every component, that you will need for web development  
Code example:
```
var testPage beleine.Page

helloworldLabel := beleine.NewLabel()
helloworldLabel.SetText("Hello world")
helloworldLabel.SetSize(1) //H1
testPage.Attach(&helloworldLabel)
fmt.Println(testPage.Render())
```

Is this not looking easy ?

Now more powerful thing:

```
var interactivePage beleine.Page

clickMeLabel := beleine.NewLabel()
clickMeLabel.SetText("Foo")
clickMeLabel.SetOnClickListener(clickMeLabel.SetTextJS("Bar"))
interactivePage.Attach(&clickMeLabel)

fmt.Println(interactivePage.Render())
```

Looking easy?  
Create working web login form!
```
var testPage beleine.Page

helloworldLabel := beleine.NewLabel()
helloworldLabel.SetText("Log-in into BeleineFramework")
helloworldLabel.SetSize(1) //H1

loginCaption := beleine.NewLabel()
loginCaption.SetText("Login")

login := beleine.MakeInput()
login.SetHint("Login")

passwordCaption := beleine.NewLabel()
passwordCaption.SetText("Password")

password := beleine.NewInput()
password.SetHint("*********")
password.SetInputType("password")

submitButton := beleine.NewButton()
submitButton.SetText("Login")
submitButton.SetOnClickListener(
beleine.PostRequestJS("/api/", fmt.Sprintf("{l:%s,p:%s}",login.GetTextJS(),password.GetTextJS()),submitButton.SetTextJS("Logged in!")) +
    submitButton.SetTextJS("Logging in..."))

testPage.Attach(&helloworldLabel)
testPage.Attach(&loginCaption)
testPage.Attach(&login)
testPage.Attach(&passwordCaption)
testPage.Attach(&password)
testPage.Attach(&submitButton)
result:= testPage.Render()
```

## Custom components
Every component have to implement following methods:
```
func render() string //This should return html code
func renderJS() string //This should return js code (if any)
func GetID() string //This should return ID obtained by beleine.GetGlobalID()
```

## Installation
`go get github.com/kubastick/BeleineFramework`

## Performance
Tested on Intel Core i3 3120m@2.50Ghz
```
Page with one label - avarage 3366 ns/render - 297 089 renders per second
Typical login page - 10 694 ns/render -  93 510 renders per second
Prepared typical login page - 7697 ns/render - 129 921 renders per second
```

## Api stability
Nope, until final release