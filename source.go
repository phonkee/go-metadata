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

import (
	"encoding/json"
	"net/url"
	"strings"
)

/*
Source is source for field value. This describes rest endpoint together with (optional) Metadata information.
*/
type Source interface {

	// Debug sets debug to source
	Debug() Source

	// isDebug returns whether debugging is enabled
	isDebug() bool

	// Action sets action
	Action(action Action) Source

	// GetAction Returns action
	GetAction() Action

	// ResultField points to resultField
	Result(field ...string) Source

	// IsValid returns whether source is setup correct
	IsValid() bool

	// Value sets value key
	Value(value string) Source

	// GetValue returns value field name
	GetValue() string

	// Display sets display field name
	//Display(display string) Action

	// Path sets path to source
	Path(string) Source

	// GetPath returns path set for this source
	GetPath() string

	// Return data
	GetData() map[string]interface{}

	// MarshalJSON satisfies json marshal interface
	MarshalJSON() ([]byte, error)
}

/*
newSource returns default source
*/
func newSource(path ...string) (result Source) {
	result = &source{}

	if len(path) > 0 {
		result.Path(path[0])
	}

	// set default
	result.Value(SOURCE_DEFAULT_VALUE_FIELD)

	return
}

/*
source is implementation of Source interface
*/
type source struct {
	// actionData
	action Action

	// debug enabled
	debug bool

	// path for given source (rest endpoint)
	path string

	// Result field (mapping to array)
	resultFieldPath []string

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
Enable debug for source
*/
func (s *source) Debug() Source {
	s.debug = true
	return s
}

/*
isDebug returns whether debug is enabled
*/
func (s *source) isDebug() bool {
	return s.debug
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

	resultField := s.action.Field(field...)

	// check if we have FIELD_ARRAY otherwise bye bye!
	if resultField.GetType() != FIELD_ARRAY {
		loggerError(s.action.isDebug(), "Result field must be FIELD_ARRAY: %#v", field)
		return
	}

	s.resultFieldPath = field
	return s
}

/*
IsValid returns whether source is correclty set and can be shown
*/
func (s *source) IsValid() bool {
	return s.GetPath() != ""
}

/*
Value sets field name within `Result`
*/
func (s *source) Value(value string) Source {
	s.valueField = value
	return s
}

/*
GetValue returns value field name
*/
func (s *source) GetValue() string {
	return s.valueField
}

/*
GetData returns data (for json marshalling etc..)
*/
func (s *source) GetData() (result map[string]interface{}) {
	result = map[string]interface{}{}

	// if not path provided we bail out.
	if s.GetPath() == "" {
		return
	}

	// add path
	result["path"] = s.GetPath()

	// result field available
	if len(s.resultFieldPath) > 0 {
		result["result"] = strings.Join(s.resultFieldPath, ".")

		if s.action != nil {
			result["metadata"] = s.action
		}
	}

	return
}

/*
Path sets path (rest endpoint)
*/
func (s *source) Path(path string) (self Source) {
	self = s

	var (
		parsed *url.URL
		err    error
	)

	if parsed, err = url.Parse(path); err != nil {
		loggerError(s.isDebug(), "cannot parse path: %#v", path)
		return
	}

	s.path = parsed.Path
	return
}

/*
GetPath sets path (rest endpoint)
*/
func (s *source) GetPath() string {
	return s.path
}

/*
MarshalJSON satisfies json marshal interface
*/
func (s *source) MarshalJSON() (result []byte, err error) {
	data := s.GetData()
	result, err = json.Marshal(data)
	return
}
