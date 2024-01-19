// Code generated by "enumer --values --type=Network --linecomment --output network_string.go --json --yaml --sql"; DO NOT EDIT.

package filter

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _NetworkName = "unknownethereumoptimismpolygonarbitrumfantomrssarweavefarcaster"

var _NetworkIndex = [...]uint8{0, 7, 15, 23, 30, 38, 44, 47, 54, 63}

const _NetworkLowerName = "unknownethereumoptimismpolygonarbitrumfantomrssarweavefarcaster"

func (i Network) String() string {
	if i >= Network(len(_NetworkIndex)-1) {
		return fmt.Sprintf("Network(%d)", i)
	}
	return _NetworkName[_NetworkIndex[i]:_NetworkIndex[i+1]]
}

func (Network) Values() []string {
	return NetworkStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NetworkNoOp() {
	var x [1]struct{}
	_ = x[NetworkUnknown-(0)]
	_ = x[NetworkEthereum-(1)]
	_ = x[NetworkOptimism-(2)]
	_ = x[NetworkPolygon-(3)]
	_ = x[NetworkArbitrum-(4)]
	_ = x[NetworkFantom-(5)]
	_ = x[NetworkRSS-(6)]
	_ = x[NetworkArweave-(7)]
	_ = x[NetworkFarcaster-(8)]
}

var _NetworkValues = []Network{NetworkUnknown, NetworkEthereum, NetworkOptimism, NetworkPolygon, NetworkArbitrum, NetworkFantom, NetworkRSS, NetworkArweave, NetworkFarcaster}

var _NetworkNameToValueMap = map[string]Network{
	_NetworkName[0:7]:        NetworkUnknown,
	_NetworkLowerName[0:7]:   NetworkUnknown,
	_NetworkName[7:15]:       NetworkEthereum,
	_NetworkLowerName[7:15]:  NetworkEthereum,
	_NetworkName[15:23]:      NetworkOptimism,
	_NetworkLowerName[15:23]: NetworkOptimism,
	_NetworkName[23:30]:      NetworkPolygon,
	_NetworkLowerName[23:30]: NetworkPolygon,
	_NetworkName[30:38]:      NetworkArbitrum,
	_NetworkLowerName[30:38]: NetworkArbitrum,
	_NetworkName[38:44]:      NetworkFantom,
	_NetworkLowerName[38:44]: NetworkFantom,
	_NetworkName[44:47]:      NetworkRSS,
	_NetworkLowerName[44:47]: NetworkRSS,
	_NetworkName[47:54]:      NetworkArweave,
	_NetworkLowerName[47:54]: NetworkArweave,
	_NetworkName[54:63]:      NetworkFarcaster,
	_NetworkLowerName[54:63]: NetworkFarcaster,
}

var _NetworkNames = []string{
	_NetworkName[0:7],
	_NetworkName[7:15],
	_NetworkName[15:23],
	_NetworkName[23:30],
	_NetworkName[30:38],
	_NetworkName[38:44],
	_NetworkName[44:47],
	_NetworkName[47:54],
	_NetworkName[54:63],
}

// NetworkString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NetworkString(s string) (Network, error) {
	if val, ok := _NetworkNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NetworkNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Network values", s)
}

// NetworkValues returns all values of the enum
func NetworkValues() []Network {
	return _NetworkValues
}

// NetworkStrings returns a slice of all String values of the enum
func NetworkStrings() []string {
	strs := make([]string, len(_NetworkNames))
	copy(strs, _NetworkNames)
	return strs
}

// IsANetwork returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Network) IsANetwork() bool {
	for _, v := range _NetworkValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Network
func (i Network) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Network
func (i *Network) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Network should be a string, got %s", data)
	}

	var err error
	*i, err = NetworkString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Network
func (i Network) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Network
func (i *Network) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = NetworkString(s)
	return err
}

func (i Network) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Network) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of Network: %[1]T(%[1]v)", value)
	}

	val, err := NetworkString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
