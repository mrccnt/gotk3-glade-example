[![license](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![goreportcard](https://goreportcard.com/badge/github.com/mrccnt/gotk3-glade-example)](https://goreportcard.com/report/github.com/mrccnt/gotk3-glade-example)

# Go GTK3 Glade Example

Experimenting with gtk using [gotk3](https://github.com/gotk3/gotk3).

![screenshot](screenshot.png)

## Environment

    Ubuntu 22.04.03 LTS
    Go 1.21.0
    GTK 3.24.33
    
## Dependencies

Install dev dependencies for `gtk`, `cairo` and `glib`

```bash
    sudo apt-get install libgtk-3-dev libcairo2-dev libglib2.0-dev
```

Check your GTK3 version:

```bash
    apt-cache policy libgtk-3-0 | grep Installed
    #
    #> Installed: 3.24.33-1ubuntu2
```

## Run, Install and Build

To run or build the code we need to pass a build tag with our given gtk3 version:

```bash
    go run     -v -tags gtk_3_24 main.go
    go install -v -tags gtk_3_24
    go build   -v -tags gtk_3_24 -o dist/gotk3-glade-example
```

## Links

 * [GOTK3: Go bindings for GTK+3](https://github.com/gotk3/gotk3)
 * [GOTK3: Examples](https://github.com/gotk3/gotk3-examples/tree/master/gtk-examples)
 * [Glade User Interface Designer](https://glade.gnome.org/)
