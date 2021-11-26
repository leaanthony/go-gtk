package main

import (
	"os"

	"github.com/leaanthony/go-gtk/_example/builder/callback"
	"github.com/leaanthony/go-gtk/gtk"
)

//"github.com/leaanthony/go-gtk/example/builder/callback"
func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()

	builder.AddFromFile("hello.ui")
	obj := builder.GetObject("window1")

	window := gtk.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", gtk.MainQuit)

	callback.Init(builder)

	gtk.Main()
}
