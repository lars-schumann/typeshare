package proto

import "encoding/json"

// This struct has a unit field
type StructHasVoidType struct {
	ThisIsAUnit struct{} `json:"thisIsAUnit"`
}
// This enum has a variant associated with unit data
type EnumHasVoidTypeTypes string
const (
	EnumHasVoidTypeTypeVariantHasAUnit EnumHasVoidTypeTypes = "hasAUnit"
)
type EnumHasVoidType struct{ 
	Type EnumHasVoidTypeTypes `json:"type"`
	content interface{}
}

func (e *EnumHasVoidType) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    EnumHasVoidTypeTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	e.Type = enum.Tag
	switch e.Type {
	case EnumHasVoidTypeTypeVariantHasAUnit:
		var res struct{}
		e.content = &res

	}
	if err := json.Unmarshal(enum.Content, &e.content); err != nil {
		return err
	}

	return nil
}

func (e EnumHasVoidType) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    EnumHasVoidTypeTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = e.Type
    enum.Content = e.content
    return json.Marshal(enum)
}

func (e EnumHasVoidType) HasAUnit() struct{} {
	res, _ := e.content.(*struct{})
	return *res
}

func NewEnumHasVoidTypeTypeVariantHasAUnit(content struct{}) EnumHasVoidType {
    return EnumHasVoidType{
        Type: EnumHasVoidTypeTypeVariantHasAUnit,
        content: &content,
    }
}

