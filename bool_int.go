package vast

import (
	"encoding/xml"
	"strconv"
)

// BoolInt is a custom type that represents a boolean value
// that marshals to "0" (false) or "1" (true) in XML attributes.
// When used as a pointer (*BoolInt), it supports three states:
// - nil: attribute not present
// - false: marshals to "0"
// - true: marshals to "1"
// This is created because go's xml/encoder marshals "0" to false and "1" to true
type BoolInt bool

// MarshalXMLAttr implements xml.MarshalerAttr to output "0" or "1"
func (b BoolInt) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	var val string
	if b {
		val = "1"
	} else {
		val = "0"
	}

	return xml.Attr{
		Name:  name,
		Value: val,
	}, nil
}

// UnmarshalXMLAttr implements xml.UnmarshalerAttr to parse "0"/"1" or "false"/"true"
func (b *BoolInt) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		*b = false
		return nil
	}

	// Try parsing as boolean (handles "true"/"false"/"1"/"0")
	parsed, err := strconv.ParseBool(attr.Value)
	if err != nil {
		return err
	}

	*b = BoolInt(parsed)
	return nil
}

// Bool returns the boolean value
func (b BoolInt) Bool() bool {
	return bool(b)
}
