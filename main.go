package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"os"
)

// WindowName is the defined identifier for the main getWindow in the glade template
const WindowName = "window"

// ListboxName is the defined identifier for the list box in the glade template
const ListboxName = "listbox"

// ButtonName is the defined identifier for the list box in the glade template
const ButtonName = "button"

// UIMain is the path to our glade file
const UIMain = "glade/main.glade"

func main() {

	gtk.Init(&os.Args)

	bldr, err := getBuilder(UIMain)
	if err != nil {
		panic(err)
	}

	window, err := getWindow(bldr)
	if err != nil {
		panic(err)
	}

	window.SetTitle("GO GTK3 Glade Example")
	window.SetDefaultSize(365, 490)
	_, err = window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	if err != nil {
		panic(err)
	}

	window.ShowAll()

	err = loadlist(bldr, []string{"hello world", "what ever", "lorem ipsum", "foo bar"})
	if err != nil {
		panic(err)
	}

	button, err := getButton(bldr)
	if err != nil {
		panic(err)
	}

	_, err = button.Connect("clicked", func() {
		fmt.Println("Click-Click!")
	})
	if err != nil {
		panic(err)
	}

	gtk.Main()
}

// getBuilder returns *gtk.getBuilder loaded with glade resource (if resource is given)
func getBuilder(filename string) (*gtk.Builder, error) {

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

// getWindow returns *gtk.Window object from the glade resource
func getWindow(b *gtk.Builder) (*gtk.Window, error) {

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

// getButton returns *gtk.Button object from the glade resource
func getButton(b *gtk.Builder) (*gtk.Button, error) {

	obj, err := b.GetObject(ButtonName)
	if err != nil {
		return nil, err
	}

	button, ok := obj.(*gtk.Button)
	if !ok {
		return nil, err
	}

	return button, nil
}

// getListbox returns *gtk.ListBox object from the glade resource
func getListbox(b *gtk.Builder) (*gtk.ListBox, error) {

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
func loadlist(b *gtk.Builder, data []string) error {

	lb, err := getListbox(b)
	if err != nil {
		return err
	}

	for index, element := range data {

		lbl, err := gtk.LabelNew(element)
		if err != nil {
			return err
		}

		lbl.SetXAlign(0)
		lbl.SetMarginStart(5)

		row, err := gtk.ListBoxRowNew()
		if err != nil {
			return err
		}

		row.Add(lbl)

		lb.Insert(row, index)
	}

	lb.ShowAll()

	return nil
}
