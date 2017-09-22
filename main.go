package main

import (
	"github.com/gotk3/gotk3/gtk"
	"os"
	"log"
)

func main() {

	gtk.Init(&os.Args)

	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Error getting new Builder: " + err.Error())
		os.Exit(1)
	}

	err = b.AddFromFile("./glade/main.glade")
	if err != nil {
		log.Fatal("Error adding glade as resource: " + err.Error())
		os.Exit(1)
	}

	obj, err := b.GetObject("window")
	if err != nil {
		log.Fatal("Error getting window from Builder: " + err.Error())
		os.Exit(1)
	}

	if window, ok := obj.(*gtk.Window); ok {
		window.Connect("destroy", func() {
			gtk.MainQuit()
		})
		window.SetDefaultSize(800, 600)
		window.ShowAll()
		gtk.Main()
		os.Exit(0)
	}

	log.Fatal("Well... Nope... ")
	os.Exit(1)
}