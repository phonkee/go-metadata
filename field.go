package metadata

import (
	"encoding/json"
	"reflect"
	"strings"
)

/*
Field interface
*/
type Field interface {
	// adds field
	AddField(string, Field) (self Field)

	// Field method returns field by names (recursively)
	Field(...string) Field

	// returns whether Field has field with given name
	HasField(...string) bool

	// Fields return mapping
	GetFields() map[string]Field

	// Description sets help text for given field
	Description(string) Field

	// GetDescription is getter for HelpText
	GetDescription() string

	// Label sets label of field
	Label(string) Field

	// GetLabel is getter for Label
	GetLabel() string

	// RemoveField removes field from fields
	RemoveField(string) (self Field)

	// set required flag to field
	Required(bool) (self Field)

	// returns whether field is required
	IsRequired() bool

	// type is type of field
	Type(string) Field

	// GetType returns field type
	GetType() string

	// From reads target and sets data to field
	From(interface{}) Field

	// Data returns data representation in dynamic fashion (interface{})
	GetData() map[string]interface{}

	// MarshalJSON satisfies json marshaller
	MarshalJSON() ([]byte, error)

	// Choices returns choices
	Choices() Choices
}

/*
NewField returns fresh Field instance
*/
func NewField() Field {
	return &field{
		fields:  map[string]Field{},
		choices: newChoices(),
	}
}

/*
field is implementation of Field interface
*/
type field struct {

	// fields of field (struct, list)
	fields map[string]Field

	// description property
	description string

	// label of field
	label string

	// required flag
	required bool

	// typ is type of field
	typ string

	// choices
	choices Choices
}

/*
AddField adds sub field
*/
func (f *field) AddField(name string, field Field) Field {
	f.fields[name] = field
	return f
}

/*
Field get or creates field (property) of field
*/
func (f *field) Field(names ...string) Field {
	if len(names) == 0 {
		panic("please provide top level field")
	}

	if _, ok := f.fields[names[0]]; !ok {
		f.fields[names[0]] = NewField()
	}

	if len(names) > 1 {
		rest := names[1:]
		return f.fields[names[0]].Type(FIELD_STRUCT).Field(rest...)
	}

	return f.fields[names[0]]
}

/*
HasField returns whether field has field with name
*/
func (f *field) HasField(names ...string) bool {
	if len(names) == 0 {
		panic("please provide top level field")
	}

	field, ok := f.fields[names[0]]

	if len(names) > 1 {
		if ok {
			remaining := names[1:]
			return field.HasField(remaining...)
		}
	}
	return ok
}

/*
Fields returns mapping of fields
*/
func (f *field) GetFields() map[string]Field {
	return f.fields
}

/*
Description sets description (for verbose metadata)
*/
func (f *field) Description(description string) Field {
	f.description = strings.TrimSpace(description)
	return f
}

/*
GetDescription is getter for help text
*/
func (f *field) GetDescription() string {
	return f.description
}

/*
Label sets label (for verbose metadata)
*/
func (f *field) Label(label string) Field {
	f.label = strings.TrimSpace(label)
	return f
}

/*
GetLabel is getter for label
*/
func (f *field) GetLabel() string {
	return f.label
}

/*
RemoveField removes sub field and returns self for method chaining
*/
func (f *field) RemoveField(name string) (self Field) {
	delete(f.fields, name)

	return f
}

/*
Required sets required flag on field
*/
func (f *field) Required(required bool) Field {
	f.required = required
	return f
}

/*
IsRequired whether required flag is set on field
*/
func (f *field) IsRequired() bool {
	return f.required
}

/*
Type sets type of field
*/
func (f *field) Type(typ string) Field {
	f.typ = typ
	return f
}

/*
GetType returns field type
*/
func (f *field) GetType() string {
	return f.typ
}

/*
From inspects given target and sets information to field
*/
func (f *field) From(target interface{}) Field {
	typ := reflect.TypeOf(target)
	required := true
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
			required = false
		} else {
			break
		}
	}

	nf := GetField(typ)
	*f = *(nf.(*field))

	// if pointer was provided set required to false
	f.Required(required)

	return f
}

/*
Data returns data (e.g. for json marshalling)
*/
func (f *field) GetData() (result map[string]interface{}) {
	result = map[string]interface{}{
		"required": f.required,
		"type":     f.typ,
	}

	if len(f.fields) > 0 {
		result["fields"] = f.fields
	}

	if f.label != "" {
		result["label"] = f.label
	}

	if f.description != "" {
		result["description"] = f.description
	}

	if f.choices.Count() > 0 {
		result["choices"] = f.choices
	}

	return
}

/*
MarshalJSON returns json representation of metadata
*/
func (f *field) MarshalJSON() (result []byte, err error) {
	result, err = json.Marshal(f.GetData())
	return
}

/*
Choices returns choices
*/
func (f *field) Choices() Choices {
	return f.choices
}