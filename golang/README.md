#SHORTCODE EXTRACTOR<br/>

#Golang Edition<br/>

```
replace inside bracket shortcode with dynamic value
```
<br/>

#INSTALATION<br/>
```
go get -u github.com/jrsmth97/shorcode-extractor/golang
```

#EXAMPLE USAGE
```
package main

import (
	"fmt"
	scextractor "github.com/jrsmth97/shorcode-extractor/golang"
)

func main() {
	var variables = make(map[string]interface{})
	variables["variable"] = "hello"
	variables["variable2"] = "world"
	variables["variable3"] = "say hi"

	scextractor.SetText("test string with {variable} inserted and (variable2) also {notexistvariable}")
	scextractor.SetBracketType(scextractor.Curly)
	scextractor.SetVariables(variables)
	result := scextractor.Extract()

	fmt.Println(result)
}

```