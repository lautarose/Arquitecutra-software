FROM golang:1.17
run mkdir /gin
add . /gin
workdir /gin
run go mod init gin
run go mod tidy
run go build -o gin .
run chmod +x /gin/gin
entrypoint ["/gin/gin"]