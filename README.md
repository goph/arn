# [Amazon Resource Name](http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)

[![Build Status](https://img.shields.io/travis/goph/arn.svg?style=flat-square)](https://travis-ci.org/goph/arn)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/arn?style=flat-square)](https://goreportcard.com/report/github.com/goph/arn)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/arn)

**[Amazon Resource Name](http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) parser and validator library.**

Amazon Resource Name (ARN) is Amazon's standard for identifying resources.
Although there are Amazon specific components in it (eg. region), in most of the cases it perfectly fits.

This library provides a way to validate and parse ARNs.

Read about ARNs in the [official documentation](http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).

## Installation

```bash
$ go get github.com/goph/arn
```


## Usage

```go
package main

import(
	"fmt"
	
	"github.com/goph/arn"
)

func main() {
	resourceName, err := arn.Parse("arn:aws:rds:eu-west-1:123456789012:db:mysql-db")
	if err != nil {
		panic(err)
	}

	fmt.Println(resourceName.String())
}
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
