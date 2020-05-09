# Go basic http server using Docker 

![Gopher image](doc/gopher.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

- [Meet Go Language](https://golang.org/)

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

- [main.go](src/main.go)
```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HolaServidor)
    http.ListenAndServe(":8080", nil)
}

func HolaServidor(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Memo Docker")
}
```

* To download it, visit [here](https://golang.org/)

* To run it locally, run ```$ go run src/main.go ``` in the current directory

* To run it in docker
1. Is needed to build it, ```docket build -t go-app```
2. Is needed to run the docker image with ```docker run --name go-container -p 8080:8080 go-app```

