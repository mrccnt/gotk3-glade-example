# Glade

Handle `*.glade` interface files and compile `.gresource` files.

## Files

    main.glade          # Glade interface file (edit with interface designer)
    resources.xml       # Edit/extend resources.xml manually to include all assets
    main.gresource      # Compressed resourse file

## Generate GResource

    # Build compressed .gresource bundle from glade files
    #
    glib-compile-resources --target=main.gresource resources.xml
