package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/teetgerink/nsxv-golang-library.git/nsxv"
)

//PACKAGENAME name for the package where the generated interfaces are located
const PACKAGENAME = "nsxv"

//WORKINGDIR the working directory is based on the package name
const WORKINGDIR = "./" + PACKAGENAME

// This script generates the contents of the `interfaces.go`

func main() {
	workingDir, _ := os.Getwd()
	fmt.Printf("Current working directory %v \n", workingDir)
	os.Chdir(WORKINGDIR)
	workingDir, _ = os.Getwd()
	fmt.Printf("New working directory %v \n", workingDir)

	item := reflect.TypeOf(nsxv.APIClient{})
	iFile, err := os.Create("interfaces.go")

	if err != nil {
		panic(err)
	}
	defer iFile.Close()

	// Generate interface methods

	// generate Package name
	generateLines(iFile, fmt.Sprintf("package %v\n\n ", PACKAGENAME))
	// generate imports
	generateLines(iFile, fmt.Sprintf("import (\n\"context\"\n,\"net/http\"\n)\n\n"))

	generateLines(iFile, fmt.Sprintf("// Compile-time interface checks\n"))
	generateLines(iFile, fmt.Sprintf("var _ IApiClient = &ApiClient{}\n\n"))
	generateLines(iFile, fmt.Sprintf("// IApiClient is the main interface for the library\n"))
	generateLines(iFile, fmt.Sprintf("type IApiClient interface {\n"))
	generateLines(iFile, fmt.Sprintf(" GetContext() context.Context\n"))

	for index := 0; index < item.NumField(); index++ {
		prop := item.Field(index)
		generateLines(iFile, fmt.Sprintf(" Get%v() I%v\n", prop.Name, prop.Name))

	}

	generateLines(iFile, fmt.Sprintf("}\n\n"))
	generateLines(iFile, fmt.Sprintf("type ApiClient struct {\n"))
	generateLines(iFile, fmt.Sprintf("  APIClient \n  Context context.Context\n}\n\n"))
	generateLines(iFile, fmt.Sprintf("func (a *ApiClient) GetContext() context.Context {\n"))
	generateLines(iFile, fmt.Sprintf("  return a.Context\n}\n\n"))

	// Generate return methods of the api client
	for index := 0; index < item.NumField(); index++ {
		prop := item.Field(index)

		funcName := fmt.Sprintf("func (a *ApiClient) Get%v() I%v {\n"+
			"   return a.%v\n"+
			"}\n\n", prop.Name, prop.Name, prop.Name)

		generateLines(iFile, funcName)

	}

	// Generate the interfaces
	for index := 0; index < item.NumField(); index++ {
		prop := item.Field(index)

		generateLines(iFile, fmt.Sprintf("type I%v interface {\n", prop.Name))

		for funcIndex := 0; funcIndex < prop.Type.NumMethod(); funcIndex++ {
			method := prop.Type.Method(funcIndex)
			generateLines(iFile, fmt.Sprintf("   %v(", method.Name))

			if method.Type.Kind() != reflect.Func {
				continue
			}

			var inArgs []string

			for argIndex := 1; argIndex < method.Type.NumIn(); argIndex++ {
				arg := method.Type.In(argIndex)

				inArgs = append(inArgs, arg.String())
			}
			generateLines(iFile, fmt.Sprintf("%v)", strings.Join(inArgs, ", ")))

			var outArgs []string

			for returnIndex := 0; returnIndex < method.Type.NumOut(); returnIndex++ {
				returnValue := method.Type.Out(returnIndex)
				outArgs = append(outArgs, returnValue.String())
			}
			generateLines(iFile, fmt.Sprintf("(%v)\n", strings.Join(outArgs, ", ")))
		}

		generateLines(iFile, fmt.Sprintf("}\n\n"))
	}

}

func generateLines(iFile *os.File, value string) {
	_, err := fmt.Fprint(iFile, value)
	if err != nil {
		fmt.Printf("error occured %v", err)
		panic(err)
	}
	fmt.Print(value)

}
