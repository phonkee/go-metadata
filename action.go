package metadata

import (
	"encoding/json"
	"reflect"
	"strings"
)

/*
Action interface that describes single action (or http method)
*/
type Action interface {

	// Description sets description for action
	Description(string) Action

	// GetDescription returns description
	GetDescription() string

	// Field adds or retrieves fields (recursively)
	Field(...string) Field

	// HasField returns whether field is set
	HasField(names ...string) bool

	// GetFields return all fields
	GetFields() map[string]Field

	// From inspects given value and makes appropriate steps
	From(v interface{}) Action

	// GetData returns dynamic data (for json etc..)
	GetData() map[string]interface{}

	// MarshalJSON satisfies json marshaller
	MarshalJSON() ([]byte, error)
}

/*
NewAction creates fresh new action
*/
func NewAction() Action {
	return &action{
		fields: map[string]Field{},
	}
}

/*
action is implementation of Action interface
*/
type action struct {

	// description of action
	description string

	// store fields
	fields map[string]Field
}

/*
Description sets action description
*/
func (a *action) Description(description string) Action {
	a.description = strings.TrimSpace(description)
	return a
}

/*
GetDescription returns action description
*/
func (a *action) GetDescription() string {
	return a.description
}

/*
Field adds or retrieves field
*/
func (a *action) Field(names ...string) Field {
	if len(names) == 0 {
		panic("please provide top level field")
	}

	if _, ok := a.fields[names[0]]; !ok {
		a.fields[names[0]] = NewField()
	}

	if len(names) > 1 {
		rest := names[1:]
		return a.fields[names[0]].Type(FIELD_STRUCT).Field(rest...)
	}

	return a.fields[names[0]]
}

/*
HasField returns whether field is set
*/
func (a *action) HasField(names ...string) (ok bool) {
	if len(names) == 0 {
		panic("please provide top level field")
	}

	field, ok := a.fields[names[0]]

	if len(names) > 1 {
		if ok {
			remaining := names[1:]
			return field.HasField(remaining...)
		}
	}
	return ok
}

/*
GetFields returns field mappings
*/
func (a *action) GetFields() map[string]Field {
	return a.fields
}

/*
Read target structure and add fields
*/
func (a *action) From(target interface{}) Action {
	typ := reflect.TypeOf(target)
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		} else {
			break
		}
	}
	if typ.Kind() != reflect.Struct {
		panic("Metadata.Action.From supports only structs")
	}

	// clear fields
	a.fields = map[string]Field{}

	f := NewField().From(target)
	for name, field := range f.GetFields() {
		a.fields[name] = field
	}

	return a
}

/*
GetData returns data for json marshalling etc..
*/
func (a *action) GetData() (result map[string]interface{}) {
	result = map[string]interface{}{}

	if len(a.fields) > 0 {
		result["fields"] = a.fields
	}

	return
}

/*
MarshalJSON returns json representation of metadata
*/
func (a *action) MarshalJSON() (result []byte, err error) {
	result, err = json.Marshal(a.GetData())
	return
}
