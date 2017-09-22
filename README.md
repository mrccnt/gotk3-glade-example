# Go GTK3 Glade Example

Experimenting with go and gtk using [gotk3](https://github.com/gotk3/gotk3) and [glade](https://glade.gnome.org/).

## Environment

I am using Ubuntu 16.04 LTS for this and Go version 1.8.3

    Ubuntu 16.04 LTS
    Go 1.8.3
    GTK 3.18

The `run.sh` and  `build.sh` scripts make use of the corresponding gtk build tags.

## Shell
    
```bash
    # Install dev dependencies for gtk, cairo and glib.
    #
    sudo apt-get install libgtk-3-dev libcairo2-dev libglib2.0-dev
    
    # Make glide available
    #
    curl https://glide.sh/get | sh
    
    # Install dependencies in ./vendor directory
    #
    glide install
    
    # Link dependecies in ./src directory
    # 
    ./link.sh
    
    # Run
    #
    ./run.sh
    
    # Build
    #
    ./build.sh
    
    # Run binary
    #
    bin/goform
```