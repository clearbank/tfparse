package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/terraform/configs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: tfparse <path to configuration>")
		os.Exit(1)
	}

	json, err := tfvariable(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(json)
}

func tfvariable(dir string) (string, error) {
	parser := configs.NewParser(nil)
	if !parser.IsConfigDir(dir) {
		return "", fmt.Errorf("%s is not a valid Terraform directory", dir)
	}
	module, diags := parser.LoadConfigDir(dir)
	if diags.HasErrors() {
		return "", fmt.Errorf(diags.Error())
	}
	output, err := json.MarshalIndent(module.Variables, "", "  ")
	return string(output), err
}
