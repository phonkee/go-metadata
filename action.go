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
	return &action{}
}

/*
action is implementation of Action interface
*/
type action struct {

	// description of action
	description string

	// field
	field Field
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

	if a.field == nil {
		a.field = newStructField()
	}

	return a.field.Field(names...)
}

/*
HasField returns whether field is set
*/
func (a *action) HasField(names ...string) bool {

	if len(names) == 0 {
		panic("please provide top level field")
	}

	if a.field == nil {
		return false
	}

	return a.field.HasField(names...)
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

	if !(typ.Kind() == reflect.Struct || typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice) {
		panic("Metadata.Action.From supports only structs/array/slice")
	}

	a.field = newField().From(target)
	return a
}

/*
GetData returns data for json marshalling etc..
*/
func (a *action) GetData() (result map[string]interface{}) {

	if a.field != nil {
		result = a.field.GetData()
	} else {
		result = map[string]interface{}{}
	}
	// no need to have this information
	delete(result, "label")
	delete(result, "description")
	delete(result, "required")

	if a.description != "" {
		result["description"] = a.description
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
