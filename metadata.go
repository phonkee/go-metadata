package metadata

import (
	"encoding/json"
	"strings"
)

/*
MetaData interface that describes heep endpoint methods
*/
type Metadata interface {
	// action returns either existing or newly created action
	Action(method string) Action

	// Name sets name to metadata
	Name(string) Metadata

	// GetName returns name of metadata
	GetName() string

	// Description sets description of metadata
	Description(string) Metadata

	// GetDescription returns description
	GetDescription() string

	// ActionCreate is alias to create "POST" action
	ActionCreate() Action

	// ActionUpdate is alias to create "POST" action
	ActionUpdate() Action

	// ActionRetrieve is alias to create "GET" action
	ActionRetrieve() Action

	// ActionDelete is alias to create "DELETE" action
	ActionDelete() Action

	// GetData returns dynamic data (for json etc..)
	GetData() map[string]interface{}

	// MarshalJSON satisfies json marshaller
	MarshalJSON() ([]byte, error)

	// @TODO: implement RemoveAction method
	// RemoveAction(method string) Metadata

}

/*
New returns new metadata instance
*/
func New() Metadata {
	return &metadata{
		actions: map[string]Action{},
	}
}

/*
Metadata interface implementation
*/
type metadata struct {
	// map to actions
	actions     map[string]Action

	// name of metadata
	name        string

	// description of metadata
	description string
}

/*
Actions either returns existing action, or it creates new one
*/
func (m *metadata) Action(method string) Action {

	method = cleanMethod(method)

	// check if action exists
	if _, ok := m.actions[method]; !ok {
		m.actions[method] = NewAction()
	}

	return m.actions[method]
}

/*
ActionCreate is alias to create "POST" action
*/
func (m *metadata) ActionCreate() Action {
	return m.Action(ACTION_CREATE)
}

/*
ActionUpdate is alias to create "POST" action
*/
func (m *metadata) ActionUpdate() Action {
	return m.Action(ACTION_UPDATE)
}

/*
ActionRetrieve is alias to create "GET" action
*/
func (m *metadata) ActionRetrieve() Action {
	return m.Action(ACTION_RETRIEVE)
}

/*
ActionDelete is alias to create "DELETE" action
*/
func (m *metadata) ActionDelete() Action {
	return m.Action(ACTION_DELETE)
}

/*
Name sets name of metadata
*/
func (m *metadata) Name(name string) Metadata {
	m.name = strings.TrimSpace(name)
	return m
}

/*
GetName returns name of metadata
*/
func (m *metadata) GetName() string {
	return m.name
}

/*
Description sets description of metadata
*/
func (m *metadata) Description(description string) Metadata {
	m.description = strings.TrimSpace(description)
	return m
}

/*
GetDescription returns description of metadata
*/
func (m *metadata) GetDescription() string {
	return m.description
}

/*
GetData returns data for json marshalling etc..
*/
func (m *metadata) GetData() (result map[string]interface{}) {
	result = map[string]interface{}{}

	if m.name != "" {
		result["name"] = m.name
	}

	if m.description != "" {
		result["description"] = m.description
	}

	if len(m.actions) > 0 {
		result["actions"] = m.actions
	}

	return
}

/*
MarshalJSON returns json representation of metadata
*/
func (m *metadata) MarshalJSON() (result []byte, err error) {
	result, err = json.Marshal(m.GetData())
	return
}
