ini
===

Parse INI file like "flag" in Go


```
go get -u github.com/c4pt0r/ini
```


Example:

```
package main

import (
	"log"
	"github.com/c4pt0r/ini"
)

var conf = ini.NewConf("test.ini")

var (
	v1 = conf.String("section1", "field1", "v1")
	v2 = conf.Int("section1", "field2", 0)
)

func main() {
	conf.Parse()

	log.Println(*v1, *v2)
}
```
