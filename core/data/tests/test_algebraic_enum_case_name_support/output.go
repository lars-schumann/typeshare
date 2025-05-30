package proto

import "encoding/json"

type ItemDetailsFieldValue struct {
}
type AdvancedColorsTypes string
const (
	AdvancedColorsTypeVariantString AdvancedColorsTypes = "string"
	AdvancedColorsTypeVariantNumber AdvancedColorsTypes = "number"
	AdvancedColorsTypeVariantNumberArray AdvancedColorsTypes = "number-array"
	AdvancedColorsTypeVariantReallyCoolType AdvancedColorsTypes = "reallyCoolType"
)
type AdvancedColors struct{ 
	Type AdvancedColorsTypes `json:"type"`
	content interface{}
}

func (a *AdvancedColors) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    AdvancedColorsTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	a.Type = enum.Tag
	switch a.Type {
	case AdvancedColorsTypeVariantString:
		var res string
		a.content = &res
	case AdvancedColorsTypeVariantNumber:
		var res int
		a.content = &res
	case AdvancedColorsTypeVariantNumberArray:
		var res []int
		a.content = &res
	case AdvancedColorsTypeVariantReallyCoolType:
		var res ItemDetailsFieldValue
		a.content = &res

	}
	if err := json.Unmarshal(enum.Content, &a.content); err != nil {
		return err
	}

	return nil
}

func (a AdvancedColors) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    AdvancedColorsTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = a.Type
    enum.Content = a.content
    return json.Marshal(enum)
}

func (a AdvancedColors) String() string {
	res, _ := a.content.(*string)
	return *res
}
func (a AdvancedColors) Number() int {
	res, _ := a.content.(*int)
	return *res
}
func (a AdvancedColors) NumberArray() []int {
	res, _ := a.content.(*[]int)
	return *res
}
func (a AdvancedColors) ReallyCoolType() *ItemDetailsFieldValue {
	res, _ := a.content.(*ItemDetailsFieldValue)
	return res
}

func NewAdvancedColorsTypeVariantString(content string) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantString,
        content: &content,
    }
}
func NewAdvancedColorsTypeVariantNumber(content int) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantNumber,
        content: &content,
    }
}
func NewAdvancedColorsTypeVariantNumberArray(content []int) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantNumberArray,
        content: &content,
    }
}
func NewAdvancedColorsTypeVariantReallyCoolType(content *ItemDetailsFieldValue) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantReallyCoolType,
        content: content,
    }
}

