package arn_test

import (
	"fmt"
	"testing"

	"github.com/goph/arn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	arn      string
	valid    bool
	expected *arn.ResourceName
}{
	{
		"",
		false,
		nil,
	},
	{
		"::::",
		false,
		nil,
	},
	{
		"arn:partition:service:region:",
		false,
		nil,
	},
	{
		":::::",
		true,
		&arn.ResourceName{},
	},
	{
		"arn:::::",
		true,
		&arn.ResourceName{
			Scheme: "arn",
		},
	},
	{
		"arn:partition:service:region:account-id:resource",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "service",
			Region: "region",
			AccountID: "account-id",
			ResourceType: "",
			ResourceDelimiter: "",
			Resource: "resource",
		},
	},
	{
		"arn:partition:service:region:account-id:resource-type/resource",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "service",
			Region: "region",
			AccountID: "account-id",
			ResourceType: "resource-type",
			ResourceDelimiter: "/",
			Resource: "resource",
		},
	},
	{
		"arn:partition:service:region:account-id:resource-type:resource",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "service",
			Region: "region",
			AccountID: "account-id",
			ResourceType: "resource-type",
			ResourceDelimiter: ":",
			Resource: "resource",
		},
	},
	{
		"arn:partition:service:region:account-id:resource-type/resource:path",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "service",
			Region: "region",
			AccountID: "account-id",
			ResourceType: "resource-type",
			ResourceDelimiter: "/",
			Resource: "resource:path",
		},
	},
	{
		"arn:partition:service:region:account-id:resource-type:resource/path",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "service",
			Region: "region",
			AccountID: "account-id",
			ResourceType: "resource-type",
			ResourceDelimiter: ":",
			Resource: "resource/path",
		},
	},
	{
		"arn:partition:service::account-id:resource-type:resource",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "service",
			Region: "",
			AccountID: "account-id",
			ResourceType: "resource-type",
			ResourceDelimiter: ":",
			Resource: "resource",
		},
	},
	{
		"arn:partition:service:::resource-type:resource",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "service",
			Region: "",
			AccountID: "",
			ResourceType: "resource-type",
			ResourceDelimiter: ":",
			Resource: "resource",
		},
	},
	{
		"arn:partition::::resource-type:resource",
		true,
		&arn.ResourceName{
			Scheme: "arn",
			Partition: "partition",
			Service: "",
			Region: "",
			AccountID: "",
			ResourceType: "resource-type",
			ResourceDelimiter: ":",
			Resource: "resource",
		},
	},
}

func TestParse(t *testing.T) {
	for _, test := range tests {
		t.Run(test.arn, func(t *testing.T) {
			err := arn.Validate(test.arn)

			if test.valid {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Equal(t, err, arn.ErrInvalid)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	for _, test := range tests {
		t.Run(test.arn, func(t *testing.T) {
			resourceName, err := arn.Parse(test.arn)

			if test.valid {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Equal(t, err, arn.ErrInvalid)
			}

			assert.Equal(t, test.expected, resourceName)
		})
	}
}

func TestResourceName_ResourceValue(t *testing.T) {
	tests := []struct {
		resourceName *arn.ResourceName
		expected     string
	}{
		{
			&arn.ResourceName{
				ResourceType:      "type",
				ResourceDelimiter: ":",
				Resource:          "resource",
			},
			"type:resource",
		},
		{
			&arn.ResourceName{
				ResourceType:      "type",
				ResourceDelimiter: "/",
				Resource:          "resource",
			},
			"type/resource",
		},
		{
			&arn.ResourceName{
				ResourceType:      "type",
				ResourceDelimiter: ":",
				Resource:          "resource/path",
			},
			"type:resource/path",
		},
		{
			&arn.ResourceName{
				ResourceType:      "type",
				ResourceDelimiter: "/",
				Resource:          "resource:path",
			},
			"type/resource:path",
		},
		{
			&arn.ResourceName{
				ResourceDelimiter: ":",
				Resource:          "resource/path",
			},
			"resource/path",
		},
		{
			&arn.ResourceName{
				ResourceDelimiter: "/",
				Resource:          "resource:path",
			},
			"resource:path",
		},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			// These values are irrelevant, so they are hardcoded to avoid duplication
			test.resourceName.Scheme = "arn"
			test.resourceName.Partition = "partition"
			test.resourceName.Service = "service"
			test.resourceName.Region = ""
			test.resourceName.AccountID = ""

			assert.Equal(t, test.expected, test.resourceName.ResourceValue())
		})
	}
}

func TestResourceName_String(t *testing.T) {
	for _, test := range tests {
		if test.valid == false {
			continue
		}

		t.Run(test.arn, func(t *testing.T) {
			assert.Equal(t, test.arn, test.expected.String())
		})
	}
}

func ExampleParse() {
	resourceName, err := arn.Parse("arn:aws:rds:eu-west-1:123456789012:db:mysql-db")
	if err != nil {
		panic(err)
	}

	fmt.Println(resourceName.String())

	// Output:
	// arn:aws:rds:eu-west-1:123456789012:db:mysql-db
}
