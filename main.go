package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/kardianos/osext"
	"os"
	"log"
	"path"
)

const WindowName = "window"
const ListboxName = "listbox"
const UiMain = "glade/main.glade"
// const ResMain = "glade/main.gresource"

func main() {

	gtk.Init(&os.Args)

	bldr, err := builder()
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
	window.SetDefaultSize(365, 470)
	window.Connect("destroy", destroy)
	window.ShowAll()

	loadlist(bldr, []string{"hello world", "what ever", "lorem ipsum", "foo bar"})

	gtk.Main()
}

func builder() (*gtk.Builder, error) {

	filename, err := osext.Executable()
	if err != nil {
		return nil, err
	}

	b, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	//b, err := gtk.BuilderNewFromResource(path.Dir(path.Dir(filename)) + "/" + ResMain)
	//if err != nil {
	//	return nil, err
	//}

	// err = b.AddFromResource(path.Dir(path.Dir(filename)) + "/" + ResMain)
	err = b.AddFromFile(path.Dir(path.Dir(filename)) + "/" + UiMain)
	if err != nil {
		return nil, err
	}

	return b, nil
}

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

func loadlist(b *gtk.Builder, data []string) {

	box, err := listbox(b)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	for index, element := range data {

		lbl, err :=  gtk.LabelNew(element)
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

func destroy() {
	gtk.MainQuit()
}