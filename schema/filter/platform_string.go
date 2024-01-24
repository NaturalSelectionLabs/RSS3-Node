// Code generated by "enumer --values --type=Platform --linecomment --output platform_string.go --json --sql"; DO NOT EDIT.

package filter

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _PlatformName = "RSS3MirrorFarcasterParagraphOpenSeaUniswapOptimismAavegotchiLensLooksRareMattersMomokaHighlightLido"

var _PlatformIndex = [...]uint8{0, 4, 10, 19, 28, 35, 42, 50, 60, 64, 73, 80, 86, 95, 99}

const _PlatformLowerName = "rss3mirrorfarcasterparagraphopenseauniswapoptimismaavegotchilenslooksraremattersmomokahighlightlido"

func (i Platform) String() string {
	i -= 1
	if i < 0 || i >= Platform(len(_PlatformIndex)-1) {
		return fmt.Sprintf("Platform(%d)", i+1)
	}
	return _PlatformName[_PlatformIndex[i]:_PlatformIndex[i+1]]
}

func (Platform) Values() []string {
	return PlatformStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PlatformNoOp() {
	var x [1]struct{}
	_ = x[PlatformRSS3-(1)]
	_ = x[PlatformMirror-(2)]
	_ = x[PlatformFarcaster-(3)]
	_ = x[PlatformParagraph-(4)]
	_ = x[PlatformOpenSea-(5)]
	_ = x[PlatformUniswap-(6)]
	_ = x[PlatformOptimism-(7)]
	_ = x[PlatformAavegotchi-(8)]
	_ = x[PlatformLens-(9)]
	_ = x[PlatformLooksRare-(10)]
	_ = x[PlatformMatters-(11)]
	_ = x[PlatformMomoka-(12)]
	_ = x[PlatformHighlight-(13)]
	_ = x[PlatformLido-(14)]
}

var _PlatformValues = []Platform{PlatformRSS3, PlatformMirror, PlatformFarcaster, PlatformParagraph, PlatformOpenSea, PlatformUniswap, PlatformOptimism, PlatformAavegotchi, PlatformLens, PlatformLooksRare, PlatformMatters, PlatformMomoka, PlatformHighlight, PlatformLido}

var _PlatformNameToValueMap = map[string]Platform{
	_PlatformName[0:4]:        PlatformRSS3,
	_PlatformLowerName[0:4]:   PlatformRSS3,
	_PlatformName[4:10]:       PlatformMirror,
	_PlatformLowerName[4:10]:  PlatformMirror,
	_PlatformName[10:19]:      PlatformFarcaster,
	_PlatformLowerName[10:19]: PlatformFarcaster,
	_PlatformName[19:28]:      PlatformParagraph,
	_PlatformLowerName[19:28]: PlatformParagraph,
	_PlatformName[28:35]:      PlatformOpenSea,
	_PlatformLowerName[28:35]: PlatformOpenSea,
	_PlatformName[35:42]:      PlatformUniswap,
	_PlatformLowerName[35:42]: PlatformUniswap,
	_PlatformName[42:50]:      PlatformOptimism,
	_PlatformLowerName[42:50]: PlatformOptimism,
	_PlatformName[50:60]:      PlatformAavegotchi,
	_PlatformLowerName[50:60]: PlatformAavegotchi,
	_PlatformName[60:64]:      PlatformLens,
	_PlatformLowerName[60:64]: PlatformLens,
	_PlatformName[64:73]:      PlatformLooksRare,
	_PlatformLowerName[64:73]: PlatformLooksRare,
	_PlatformName[73:80]:      PlatformMatters,
	_PlatformLowerName[73:80]: PlatformMatters,
	_PlatformName[80:86]:      PlatformMomoka,
	_PlatformLowerName[80:86]: PlatformMomoka,
	_PlatformName[86:95]:      PlatformHighlight,
	_PlatformLowerName[86:95]: PlatformHighlight,
	_PlatformName[95:99]:      PlatformLido,
	_PlatformLowerName[95:99]: PlatformLido,
}

var _PlatformNames = []string{
	_PlatformName[0:4],
	_PlatformName[4:10],
	_PlatformName[10:19],
	_PlatformName[19:28],
	_PlatformName[28:35],
	_PlatformName[35:42],
	_PlatformName[42:50],
	_PlatformName[50:60],
	_PlatformName[60:64],
	_PlatformName[64:73],
	_PlatformName[73:80],
	_PlatformName[80:86],
	_PlatformName[86:95],
	_PlatformName[95:99],
}

// PlatformString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PlatformString(s string) (Platform, error) {
	if val, ok := _PlatformNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PlatformNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Platform values", s)
}

// PlatformValues returns all values of the enum
func PlatformValues() []Platform {
	return _PlatformValues
}

// PlatformStrings returns a slice of all String values of the enum
func PlatformStrings() []string {
	strs := make([]string, len(_PlatformNames))
	copy(strs, _PlatformNames)
	return strs
}

// IsAPlatform returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Platform) IsAPlatform() bool {
	for _, v := range _PlatformValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Platform
func (i Platform) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Platform
func (i *Platform) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Platform should be a string, got %s", data)
	}

	var err error
	*i, err = PlatformString(s)
	return err
}

func (i Platform) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Platform) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of Platform: %[1]T(%[1]v)", value)
	}

	val, err := PlatformString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
