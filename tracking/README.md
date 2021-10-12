# go-face examples
## Tracking

## Requirements

To compile face you need to have [go-face](github.com/Danile71/go-face) and [gocv](gocv.io/x/gocv)

![screen](./images/screen.jpg)

[youtube](https://youtu.be/JFRfxLJ9CIM)


### Ubuntu:
```
1. go build -tags gocv
2. ./tracking
```

### Mac m1:
```
1. export LIBRARY_PATH=$LIBRARY_PATH:/opt/homebrew/lib
2. CPLUS_INCLUDE_PATH=/opt/homebrew/include go build -tags gocv
3. ./tracking
```
