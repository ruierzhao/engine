package assets

// To generate file with fonts binary data install "go-bindata" from:
// https://github.com/go-bindata/go-bindata
// > go install github.com/go-bindata/go-bindata@latest

//go:generate go-bindata -o data.go -pkg assets fonts cursors
//go:generate g3nicodes -pkg icon icon/codepoints icon/icodes.go
//FreeMono.ttf FreeSans.ttf FreeSansBold.ttf MaterialIcons-Regular.ttf
