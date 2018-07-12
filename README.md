# Beleine Framework
[![Build Status](https://travis-ci.com/kubastick/BeleineFramework.svg?token=bpH3SmbkspnyvKjSgrmq&branch=master)](https://travis-ci.com/kubastick/BeleineFramework)
[![GoDoc](https://godoc.org/github.com/kubastick/BeleineFramework?status.svg)](https://godoc.org/github.com/kubastick/BeleineFramework)  
Tried of html,css and javascript ?  
Beleine Framework is the fastest way to build web applications in Go  
It gives you high level api, and every component, that you will need for web development  
Code example:
```
var testPage Page

helloworldLabel := MakeLabel()
helloworldLabel.SetText("Hello world")
helloworldLabel.SetSize(1) //H1
testPage.Attach(&helloworldLabel)
fmt.Println(testPage.Render())
```

Is this not looking easy ?

Now more powerful thing:

```
var interactivePage Page

clickMeLabel := MakeLabel()
clickMeLabel.SetText("Foo")
clickMeLabel.SetOnClickListener(clickMeLabel.SetTextJS("Bar"))
interactivePage.Attach(&clickMeLabel)

fmt.Println(interactivePage.Render())
```
