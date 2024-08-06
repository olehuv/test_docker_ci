# syntax=docker/dockerfile:1
FROM golang:1.21 AS build
WORKDIR /src
COPY <<EOF ./main.go
package main

import "fmt"
import "time"

func main() {
  fmt.Println("HI AVIHEYA! Hi, Oleh! This is an image 1")
  currentTime := time.Now()
  fmt.Println("Current time:", currentTime.Format("2006-01-02 15:04:05"))
}
EOF
RUN go build -o /bin/hello ./main.go

FROM scratch
COPY --from=0 /bin/hello /bin/hello
CMD ["/bin/hello"]
