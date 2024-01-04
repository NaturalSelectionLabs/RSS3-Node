// Code generated by "enumer --values --type=Name --linecomment --output worker_string.go --json --yaml --sql"; DO NOT EDIT.

package engine

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _NameName = "fallbackmirrorfarcasterrss3paragraphopenseauniswap"

var _NameIndex = [...]uint8{0, 8, 14, 23, 27, 36, 43, 50}

const _NameLowerName = "fallbackmirrorfarcasterrss3paragraphopenseauniswap"

func (i Name) String() string {
	i -= 1
	if i < 0 || i >= Name(len(_NameIndex)-1) {
		return fmt.Sprintf("Name(%d)", i+1)
	}
	return _NameName[_NameIndex[i]:_NameIndex[i+1]]
}

func (Name) Values() []string {
	return NameStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NameNoOp() {
	var x [1]struct{}
	_ = x[Fallback-(1)]
	_ = x[Mirror-(2)]
	_ = x[Farcaster-(3)]
	_ = x[RSS3-(4)]
	_ = x[Paragraph-(5)]
	_ = x[OpenSea-(6)]
	_ = x[Uniswap-(7)]
}

var _NameValues = []Name{Fallback, Mirror, Farcaster, RSS3, Paragraph, OpenSea, Uniswap}

var _NameNameToValueMap = map[string]Name{
	_NameName[0:8]:        Fallback,
	_NameLowerName[0:8]:   Fallback,
	_NameName[8:14]:       Mirror,
	_NameLowerName[8:14]:  Mirror,
	_NameName[14:23]:      Farcaster,
	_NameLowerName[14:23]: Farcaster,
	_NameName[23:27]:      RSS3,
	_NameLowerName[23:27]: RSS3,
	_NameName[27:36]:      Paragraph,
	_NameLowerName[27:36]: Paragraph,
	_NameName[36:43]:      OpenSea,
	_NameLowerName[36:43]: OpenSea,
	_NameName[43:50]:      Uniswap,
	_NameLowerName[43:50]: Uniswap,
}

var _NameNames = []string{
	_NameName[0:8],
	_NameName[8:14],
	_NameName[14:23],
	_NameName[23:27],
	_NameName[27:36],
	_NameName[36:43],
	_NameName[43:50],
}

// NameString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NameString(s string) (Name, error) {
	if val, ok := _NameNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NameNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Name values", s)
}

// NameValues returns all values of the enum
func NameValues() []Name {
	return _NameValues
}

// NameStrings returns a slice of all String values of the enum
func NameStrings() []string {
	strs := make([]string, len(_NameNames))
	copy(strs, _NameNames)
	return strs
}

// IsAName returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Name) IsAName() bool {
	for _, v := range _NameValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Name
func (i Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Name
func (i *Name) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Name should be a string, got %s", data)
	}

	var err error
	*i, err = NameString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Name
func (i Name) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Name
func (i *Name) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = NameString(s)
	return err
}

func (i Name) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Name) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of Name: %[1]T(%[1]v)", value)
	}

	val, err := NameString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
