package arn

import (
	"errors"
	"strings"
)

// ErrInvalid is returned when an ARN fails validation checks.
var ErrInvalid = errors.New("invalid ARN")

// ResourceName represents the actual resource name.
type ResourceName struct {
	// Scheme is always "arn". Placed here for future compatibility.
	Scheme string

	// Partition is the partition that the resource is in.
	Partition string

	// Service identifies the actual product.
	Service string

	// Region is the region the resource resides in.
	// This component can sometimes be omitted, depends on the implementation.
	Region string

	// AccountID is the ID of the account that owns the resource.
	// This can be ignored in most of the cases when the resource itself can be identified
	// (eg. when the resource is unique).
	AccountID string

	// ResourceType is the optional type of resource.
	// Whether it is required or not depends entirely on the implementation.
	ResourceType string

	// ResourceDelimiter delimits the ResourceType and the Resource.
	// It is necessary to be stored here in order to restore the original Scheme.
	ResourceDelimiter string

	// Resource is the actual identifier of the resource.
	Resource string
}

// Parse accepts an ARN parses it into a ResourceName.
func Parse(arn string) (*ResourceName, error) {
	components, delimiter := parse(arn)

	err := validate(components)
	if err != nil {
		return nil, err
	}

	return &ResourceName{
		Scheme:            components[0],
		Partition:         components[1],
		Service:           components[2],
		Region:            components[3],
		AccountID:         components[4],
		ResourceType:      components[6],
		ResourceDelimiter: delimiter,
		Resource:          components[5],
	}, nil
}

// parse parses an ARN into a string slice.
func parse(arn string) ([]string, string) {
	components := strings.SplitN(arn, ":", 6)

	// Leave this to validate to fail
	if len(components) < 6 {
		return components, ""
	}

	// Increase slice capacity with one for the resource type.
	// Additional splitting done bellow.
	components = append(components, "")

	var delimiter string

	// Check the resource delimiter (if any)
	colonDelimiter := strings.Index(components[5], ":")
	slashDelimiter := strings.Index(components[5], "/")

	if colonDelimiter > -1 && (colonDelimiter <= slashDelimiter || slashDelimiter == -1) {
		delimiter = ":"
	} else if slashDelimiter > -1 && (colonDelimiter > slashDelimiter || colonDelimiter == -1) {
		delimiter = "/"
	}

	// There is a resource delimiter
	if delimiter != "" {
		resourceComponents := strings.SplitN(components[5], delimiter, 2)

		// Resource is first, resource type is second since it's optional
		components[5] = resourceComponents[1]
		components[6] = resourceComponents[0]
	}

	return components, delimiter
}

// Validate checks whether the input is a valid ARN.
func Validate(arn string) error {
	components, _ := parse(arn)

	return validate(components)
}

// validate checks whether a partially parsed ARN is valid.
func validate(components []string) error {
	if len(components) < 6 {
		return ErrInvalid
	}

	return nil
}

// ResourceValue returns the last component of the ARN.
func (n *ResourceName) ResourceValue() string {
	if n.ResourceType == "" {
		return n.Resource
	}

	return n.ResourceType + n.ResourceDelimiter + n.Resource
}

// String returns the original ARN.
func (n *ResourceName) String() string {
	components := []string{
		n.Scheme,
		n.Partition,
		n.Service,
		n.Region,
		n.AccountID,
		n.ResourceValue(),
	}

	return strings.Join(components, ":")
}
