package arn

import (
	"database/sql/driver"
	"fmt"
)

// Value converts the ResourceName into a SQL driver value which can be used to
// directly use the ResourceName as parameter to a SQL query.
func (n *ResourceName) Value() (driver.Value, error) {
	if n == nil {
		return nil, nil
	}

	return n.String(), nil
}

// Scan implements the sql.Scanner interface. It supports converting from
// string, []byte, or nil into a ResourceName value. Attempting to convert from
// another type will return an error.
func (n *ResourceName) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		resourceName, err := Parse(string(v))
		*n = *resourceName

		return err
	case string:
		resourceName, err := Parse(v)
		*n = *resourceName

		return err
	default:
		return fmt.Errorf("scan: unable to scan type %T into ResourceName", v)
	}
}
