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
	// NotifierTypeNaV is a NotifierType of type NaV.
	NotifierTypeNaV NotifierType = iota
	// NotifierTypeConsole is a NotifierType of type Console.
	NotifierTypeConsole
	// NotifierTypeWebhook is a NotifierType of type Webhook.
	NotifierTypeWebhook
	// NotifierTypeRabbitmq is a NotifierType of type Rabbitmq.
	NotifierTypeRabbitmq
)

const _NotifierTypeName = "NaVconsolewebhookrabbitmq"

var _NotifierTypeNames = []string{
	_NotifierTypeName[0:3],
	_NotifierTypeName[3:10],
	_NotifierTypeName[10:17],
	_NotifierTypeName[17:25],
}

// NotifierTypeNames returns a list of possible string values of NotifierType.
func NotifierTypeNames() []string {
	tmp := make([]string, len(_NotifierTypeNames))
	copy(tmp, _NotifierTypeNames)
	return tmp
}

var _NotifierTypeMap = map[NotifierType]string{
	NotifierTypeNaV:      _NotifierTypeName[0:3],
	NotifierTypeConsole:  _NotifierTypeName[3:10],
	NotifierTypeWebhook:  _NotifierTypeName[10:17],
	NotifierTypeRabbitmq: _NotifierTypeName[17:25],
}

// String implements the Stringer interface.
func (x NotifierType) String() string {
	if str, ok := _NotifierTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NotifierType(%d)", x)
}

var _NotifierTypeValue = map[string]NotifierType{
	_NotifierTypeName[0:3]:                    NotifierTypeNaV,
	strings.ToLower(_NotifierTypeName[0:3]):   NotifierTypeNaV,
	_NotifierTypeName[3:10]:                   NotifierTypeConsole,
	strings.ToLower(_NotifierTypeName[3:10]):  NotifierTypeConsole,
	_NotifierTypeName[10:17]:                  NotifierTypeWebhook,
	strings.ToLower(_NotifierTypeName[10:17]): NotifierTypeWebhook,
	_NotifierTypeName[17:25]:                  NotifierTypeRabbitmq,
	strings.ToLower(_NotifierTypeName[17:25]): NotifierTypeRabbitmq,
}

// ParseNotifierType attempts to convert a string to a NotifierType.
func ParseNotifierType(name string) (NotifierType, error) {
	if x, ok := _NotifierTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _NotifierTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return NotifierType(0), fmt.Errorf("%s is not a valid NotifierType, try [%s]", name, strings.Join(_NotifierTypeNames, ", "))
}
