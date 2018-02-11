package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
)

// WindowName is the defined identifier for the main window in the glade template
const WindowName = "window"

// ListboxName is the defined identifier for the list box in the glade template
const ListboxName = "listbox"

// UiMain represent the path to our glade file
const UiMain = "glade/main.glade"

func main() {

	gtk.Init(&os.Args)

	bldr, err := builder(UiMain)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	window, err := window(bldr)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	window.SetTitle("GO GTK3 Glade Example")
	window.SetDefaultSize(365, 490)
	window.Connect("destroy", destroy)
	window.ShowAll()

	loadlist(bldr, []string{"hello world", "what ever", "lorem ipsum", "foo bar"})

	gtk.Main()
}

// builder returns pointer to gtk.builder loaded with glade resource (if resource is given)
func builder(filename string) (*gtk.Builder, error) {

	b, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	if filename != "" {
		err = b.AddFromFile(filename)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

// windows returns gtk.Window object from the glade resource
func window(b *gtk.Builder) (*gtk.Window, error) {

	obj, err := b.GetObject(WindowName)
	if err != nil {
		return nil, err
	}

	window, ok := obj.(*gtk.Window)
	if !ok {
		return nil, err
	}

	return window, nil
}

// windows returns gtk.ListBox object from the glade resource
func listbox(b *gtk.Builder) (*gtk.ListBox, error) {

	obj, err := b.GetObject(ListboxName)
	if err != nil {
		return nil, err
	}

	lb, ok := obj.(*gtk.ListBox)
	if !ok {
		return nil, err
	}

	return lb, nil
}

// loadlist populates ListBox with data
func loadlist(b *gtk.Builder, data []string) {

	box, err := listbox(b)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	for index, element := range data {

		lbl, err := gtk.LabelNew(element)
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}

		lbl.SetXAlign(0)
		lbl.SetMarginStart(5)

		row, err := gtk.ListBoxRowNew()
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}

		row.Add(lbl)

		box.Insert(row, index)
	}

	box.ShowAll()
}

// destroy is the triggered handler when closing/destroying the gui window
func destroy() {
	gtk.MainQuit()
}
