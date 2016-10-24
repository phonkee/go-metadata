/*
Source provides way how to specify choices rest endpoint for given field

This is proposal:

    md := New()
    userlist := md.Action(ACTION_POST).Field("result").From([]User{})
    statusfield := userlist.Field("status")
    statusfield.
        Source("/api/user/status").
        Action(action).
        Mapping("result", "value", "display")

*/
package metadata

/*
Source is source for field value. This describes rest endpoint together with (optional) Metadata information.
*/
type Source interface {

	// Action sets action
	Action(action Action) Source

	// GetAction Returns action
	GetAction() Action

	// ResultField points to resultField
	Result(field ...string) Source

	// Value sets value key
	//Value(value string) Action

	// Display sets display field name
	//Display(display string) Action

	// Path sets path to source
	//Path(string) Source

	// Return data
	//GetData() map[string]interface{}

	// MarshalJSON satisfies json marshal interface
	//MarshalJSON() ([]byte, error)
}

/*
newSource returns default source
*/
func newSource() Source {
	return &source{}
}

/*
source is implementation of Source interface
*/
type source struct {
	// actionData
	action Action

	// Result field (mapping to array)
	resultField     Field
	resultFieldName string

	// value fieldname
	valueField string

	// display fieldname
	displayField string
}

/*
Action sets action to source
*/
func (s *source) Action(action Action) Source {
	s.action = action
	return s
}

/*
GetAction returns action, if not given, blank action is returned
*/
func (s *source) GetAction() Action {
	if s.action == nil {
		s.Action(newAction())
	}

	return s.action
}

/*
Result points to correct field
*/
func (s *source) Result(field ...string) (result Source) {

	result = s
	if s.action == nil {
		panic("Please set Source.Action first")
	}

	if !s.action.HasField(field...) {
		loggerWarning(s.action.isDebug(), "Result field does not exist: %#v", field)
		return
	}

	// set resultFieldName to provide in GetData
	s.resultFieldName = field[len(field)-1]

	resultField := s.action.Field(field...)

	// check if we have FIELD_ARRAY otherwise bye bye!
	if resultField.GetType() != FIELD_ARRAY {
		loggerError(s.action.isDebug(), "Result field must be FIELD_ARRAY: %#v", field)
		return
	}

	s.resultField = resultField
	return s
}

/*
GetData returns data (for json marshalling etc..)
*/
func (s *source) GetData() (result map[string]interface{}) {
	result = map[string]interface{}{}

	// result field available
	if s.resultField != nil {

		//s.resultFieldName

	}

	return
}
