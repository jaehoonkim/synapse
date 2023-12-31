// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package v2

import (
	"fmt"
	"strings"
)

const (
	// PollingTypeRegular is a PollingType of type Regular.
	PollingTypeRegular PollingType = iota
	// PollingTypeSmart is a PollingType of type Smart.
	PollingTypeSmart
)

const _PollingTypeName = "regularsmart"

var _PollingTypeNames = []string{
	_PollingTypeName[0:7],
	_PollingTypeName[7:12],
}

// PollingTypeNames returns a list of possible string values of PollingType.
func PollingTypeNames() []string {
	tmp := make([]string, len(_PollingTypeNames))
	copy(tmp, _PollingTypeNames)
	return tmp
}

var _PollingTypeMap = map[PollingType]string{
	PollingTypeRegular: _PollingTypeName[0:7],
	PollingTypeSmart:   _PollingTypeName[7:12],
}

// String implements the Stringer interface.
func (x PollingType) String() string {
	if str, ok := _PollingTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PollingType(%d)", x)
}

var _PollingTypeValue = map[string]PollingType{
	_PollingTypeName[0:7]:                   PollingTypeRegular,
	strings.ToLower(_PollingTypeName[0:7]):  PollingTypeRegular,
	_PollingTypeName[7:12]:                  PollingTypeSmart,
	strings.ToLower(_PollingTypeName[7:12]): PollingTypeSmart,
}

// ParsePollingType attempts to convert a string to a PollingType.
func ParsePollingType(name string) (PollingType, error) {
	if x, ok := _PollingTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _PollingTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return PollingType(0), fmt.Errorf("%s is not a valid PollingType, try [%s]", name, strings.Join(_PollingTypeNames, ", "))
}
