# Beleine Framework
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
