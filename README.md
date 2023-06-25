# GO-GS
## Golang wrapper on ghost-script
### Process large number of pdfs with all gs properties
### Extract images from all kind of pdfs

## Todo List

- [ ] adding Ghostscript details or binary
- [ ] create codebase for go-gs
- [ ] create workerpool and thread
- [ ] write tests

### Ghostscript
- Download ghostscript binaries or executable from the site https://ghostscript.com/releases/gsdnld.html
- Update path variable to use Ghostscript (or specify binary/exe path in env variable) -- decide
- install cgo for executing gs

### Build go app
```
go build .
```

### Run go binary/executable
```
go-gs -config="./tmp/config.json" -worker=4 -thread=4 -path=<folder containing pdfs>
```