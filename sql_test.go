package arn_test

import (
	"testing"

	"github.com/banzaicloud/arn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceName_Value(t *testing.T) {
	for _, test := range tests {
		t.Run(test.arn, func(t *testing.T) {
			if test.valid == false {
				t.SkipNow()
			}

			v, err := test.expected.Value()
			require.NoError(t, err)

			s, ok := v.(string)
			require.True(t, ok)
			assert.Equal(t, test.arn, s)
		})
	}
}

func TestResourceName_Scan(t *testing.T) {
	for _, test := range tests {
		t.Run(test.arn, func(t *testing.T) {
			if test.valid == false {
				t.SkipNow()
			}

			rn := new(arn.ResourceName)

			err := rn.Scan(test.arn)
			require.NoError(t, err)
			assert.Equal(t, test.expected, rn)

			err = rn.Scan([]byte(test.arn))
			require.NoError(t, err)
			assert.Equal(t, test.expected, rn)
		})
	}
}

func TestResourceName_Scan_Invalid(t *testing.T) {
	rn := new(arn.ResourceName)

	err := rn.Scan(1234)
	require.Error(t, err)
}
