Description
===========

A simple Bloom Filter implementation

Example
=======

```
  package main
  
  import (
    "fmt"
    "github.com/lazybeaver/bloomfilter"
  )
  
  func main() {
    bf := bloomfilter.New(10, 2)
    bf.Add([]byte("somekey"))
    if bf.Contains([]byte("somekey")) {
      fmt.Println("Found")
    }
  }
```
