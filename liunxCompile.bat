:: 在win上进行liunx可执行交叉文件编译

set GOOS=linux
set GOARCH=amd64

go build main.go