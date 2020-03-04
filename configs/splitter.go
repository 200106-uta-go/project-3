package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("istio-demo.yaml")
	if err != nil {
		fmt.Println("File reading error", err)
	}

	dataString := string(data)

	//var kind, name string

	dataSlice := strings.Split(dataString, "---")
	for _, s := range dataSlice {
		sSlice := strings.Split(s, "\n")
		for _, ss := range sSlice {
			if strings.HasPrefix(ss, "kind:") {

				fmt.Println(ss)
			}
			if strings.HasPrefix(ss, "  name:") {
				fmt.Println(ss)
			}

			// ssSlice := strings.Split(ss, " ")
			// for i, sss := range ssSlice {
			// 	if sss == "kind:" {
			// 		kind = ssSlice[i+1]
			// 		break
			// 	}
			// 	if sss == "name:" {
			// 		name = ssSlice[i+1]
			// 	}
			// }
		}
		// fmt.Println(kind + "-" + name)
	}
}
