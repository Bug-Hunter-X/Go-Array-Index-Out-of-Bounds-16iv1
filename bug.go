```go
package main

import "fmt"

func main() {
    var a [5]int
    for i := 0; i <= 5; i++ { //Error: Index out of bounds
        a[i] = i
    }
    fmt.Println(a)
}
```