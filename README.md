# tfparse

A command line tool for parsing Terraform configuration files.

Currently it just prints out all input variables in json format.

This is useful if you want to generate tfvars files and need to know what input variables a configuration expects.

(Inspired by [tfjson](https://github.com/palantir/tfjson))

## Installation

```
$ go get github.com/clearbank/tfparse
```

## Usage

Running against the following configuration:

```hcl
variable "environment" {
  type        = string
  description = "Name of the environment"
}
```

Output:

```json
$ tfparse /terraform-module-folder

{
  "environment": {
    "Name": "environment",
    "Description": "Name of the environment",
    "Default": {},
    "Type": "string",
    "ParsingMode": 76,
    "DescriptionSet": true,
    "DeclRange": {
      "Filename": "/terraform-module-folder/example.tf",
      "Start": {
        "Line": 1,
        "Column": 1,
        "Byte": 0
      },
      "End": {
        "Line": 1,
        "Column": 23,
        "Byte": 22
      }
    }
  }
}
```