## inspect check

Check the image for recommended labels

### Synopsis

Check the images for recommended labels, provides a tabular output with recommendations if a particular label does not exist.

```
inspect check <image>... [flags]
```

### Examples

```
inspect check alpine:3.13.0
```

### Options

```
  -f, --fail-on-recommendations   Should exit 1 if there are recommendations
```

### Options inherited from parent commands

```
  -v, --debug        Debug Output
      --help         Show help for command
  -m, --markdown     Display all tables in Markdown format
      --no-headers   Do not display table headers
  -r, --raw          Display all tables in raw format
```

### SEE ALSO

* [inspect](inspect.md)	 - inspect a docker image

