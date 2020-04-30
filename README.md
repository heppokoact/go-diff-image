# go-diff-image

A diff tool for images.

It is forked from [murooka/go-diff-image](https://github.com/murooka/go-diff-image) and modified to compile into C shared libraries.

## Build

```sh
go get -u github.com/sergi/go-diff/diffmatchpatch

# Unix
go build -o diffImage.so -buildmode=c-shared ./cmd/diff-image/main.go
# Windows
go build -o diffImage.dll -buildmode=c-shared ./cmd/diff-image/main.go
```
