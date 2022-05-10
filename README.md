# exif-rename
Rename jpg files according to EXIF datetime

# clone
```
mkdir $GOPATH/github.com/devldavydov
cd $GOPATH/github.com/devldavydov
gh repo clone devldavydov/exif-rename

cd $GOPATH/github.com/devldavydov/exif-rename
go mod tidy
```
# build
```
cd $GOPATH/github.com/devldavydov/exif-rename
go install ./cmd/exif-rename
```

# run
```
export PATH=$PATH:$GOPATH/bin
exif-rename -path <path to images folder> [-dry]
```