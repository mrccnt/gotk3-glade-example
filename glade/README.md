# Glade

If you need to package more files or other assets this may come in handy. But do not forget to implement loading those
resources in code! It is not part of the example here. ;)

## Files

    main.glade          # Glade interface file (edit with interface designer)
    resources.xml       # Edit/extend resources.xml manually to include all assets
    main.gresource      # Compressed resourse file

## Generate GResource

Build compressed `.gresource` bundle from interface and resources:

```bash
    glib-compile-resources --target=main.gresource resources.xml
```

