# go-face examples
## Russian plate numbers

## Requirements

To compile face you need to have [go-face](github.com/Danile71/go-face) and [gocv](gocv.io/x/gocv)

![screen](./images/screen.jpg)

### Ubuntu:
```
1. go build -tags gocv
2. ./plate
```

### Mac m1:
```
1. export LIBRARY_PATH=$LIBRARY_PATH:/opt/homebrew/lib
2. CPLUS_INCLUDE_PATH=/opt/homebrew/include go build -tags gocv
3. ./plate
```
