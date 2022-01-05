package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/teetgerink/nsxv-golang-library.git/nsxv"
)

// This script generates the contents of the `interfaces.go` for you to copy-paste
// It's kept here in case we ever need it again.
func main() {

	item := reflect.TypeOf(nsxv.APIClient{})

	// Generate interface methods
	for index := 0; index < item.NumField(); index++ {
		prop := item.Field(index)

		fmt.Printf("Get%v() I%v\n", prop.Name, prop.Name)
	}

	// Generate return methods of the api client
	for index := 0; index < item.NumField(); index++ {
		prop := item.Field(index)

		fmt.Printf("func (a *ApiClient) Get%v() I%v {\n"+
			"   return a.%v\n"+
			"}\n\n", prop.Name, prop.Name, prop.Name)
	}

	// Generate the interfaces
	for index := 0; index < item.NumField(); index++ {
		prop := item.Field(index)

		fmt.Printf("type I%v interface {\n", prop.Name)

		for funcIndex := 0; funcIndex < prop.Type.NumMethod(); funcIndex++ {
			method := prop.Type.Method(funcIndex)
			fmt.Printf("   %v(", method.Name)

			if method.Type.Kind() != reflect.Func {
				continue
			}

			var inArgs []string

			for argIndex := 1; argIndex < method.Type.NumIn(); argIndex++ {
				arg := method.Type.In(argIndex)
				inArgs = append(inArgs, arg.String())
			}

			fmt.Printf("%v)", strings.Join(inArgs, ", "))

			var outArgs []string

			for returnIndex := 0; returnIndex < method.Type.NumOut(); returnIndex++ {
				returnValue := method.Type.Out(returnIndex)
				outArgs = append(outArgs, returnValue.String())
			}

			fmt.Printf("(%v)\n", strings.Join(outArgs, ", "))
		}

		fmt.Printf("}\n\n")
	}
}
