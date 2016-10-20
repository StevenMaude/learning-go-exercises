package main

import "fmt"

func main() {
	a := "package main\n\nimport \"fmt\"\n\nfunc main() {"
	b := "%s\n\ta := %#v\n\tc := %#v\n\tfmt.Printf(b, a, a, b)\n}\n"
	fmt.Printf(b, a, a, b)
}
