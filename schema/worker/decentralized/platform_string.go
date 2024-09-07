// Code generated by "enumer --values --type=Platform --linecomment --output platform_string.go --json --yaml --sql"; DO NOT EDIT.

package decentralized

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _PlatformName = "Unknown1inchAAVEAavegotchiCowArbitrumBendDAOCrossbellCurveENSFarcasterHighlightIQWikiKiwiStandLensLidoLooksRareMattersMirrorOpenSeaOptimismParagraphParaswapRSS3SAVMStargateUniswapVSL"

var _PlatformIndex = [...]uint8{0, 7, 12, 16, 26, 29, 37, 44, 53, 58, 61, 70, 79, 85, 94, 98, 102, 111, 118, 124, 131, 139, 148, 156, 160, 164, 172, 179, 182}

const _PlatformLowerName = "unknown1inchaaveaavegotchicowarbitrumbenddaocrossbellcurveensfarcasterhighlightiqwikikiwistandlenslidolooksraremattersmirroropenseaoptimismparagraphparaswaprss3savmstargateuniswapvsl"

func (i Platform) String() string {
	if i >= Platform(len(_PlatformIndex)-1) {
		return fmt.Sprintf("Platform(%d)", i)
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
	_ = x[PlatformUnknown-(0)]
	_ = x[Platform1Inch-(1)]
	_ = x[PlatformAAVE-(2)]
	_ = x[PlatformAavegotchi-(3)]
	_ = x[PlatformCow-(4)]
	_ = x[PlatformArbitrum-(5)]
	_ = x[PlatformBendDAO-(6)]
	_ = x[PlatformCrossbell-(7)]
	_ = x[PlatformCurve-(8)]
	_ = x[PlatformENS-(9)]
	_ = x[PlatformFarcaster-(10)]
	_ = x[PlatformHighlight-(11)]
	_ = x[PlatformIQWiki-(12)]
	_ = x[PlatformKiwiStand-(13)]
	_ = x[PlatformLens-(14)]
	_ = x[PlatformLido-(15)]
	_ = x[PlatformLooksRare-(16)]
	_ = x[PlatformMatters-(17)]
	_ = x[PlatformMirror-(18)]
	_ = x[PlatformOpenSea-(19)]
	_ = x[PlatformOptimism-(20)]
	_ = x[PlatformParagraph-(21)]
	_ = x[PlatformParaswap-(22)]
	_ = x[PlatformRSS3-(23)]
	_ = x[PlatformSAVM-(24)]
	_ = x[PlatformStargate-(25)]
	_ = x[PlatformUniswap-(26)]
	_ = x[PlatformVSL-(27)]
}

var _PlatformValues = []Platform{PlatformUnknown, Platform1Inch, PlatformAAVE, PlatformAavegotchi, PlatformCow, PlatformArbitrum, PlatformBendDAO, PlatformCrossbell, PlatformCurve, PlatformENS, PlatformFarcaster, PlatformHighlight, PlatformIQWiki, PlatformKiwiStand, PlatformLens, PlatformLido, PlatformLooksRare, PlatformMatters, PlatformMirror, PlatformOpenSea, PlatformOptimism, PlatformParagraph, PlatformParaswap, PlatformRSS3, PlatformSAVM, PlatformStargate, PlatformUniswap, PlatformVSL}

var _PlatformNameToValueMap = map[string]Platform{
	_PlatformName[0:7]:          PlatformUnknown,
	_PlatformLowerName[0:7]:     PlatformUnknown,
	_PlatformName[7:12]:         Platform1Inch,
	_PlatformLowerName[7:12]:    Platform1Inch,
	_PlatformName[12:16]:        PlatformAAVE,
	_PlatformLowerName[12:16]:   PlatformAAVE,
	_PlatformName[16:26]:        PlatformAavegotchi,
	_PlatformLowerName[16:26]:   PlatformAavegotchi,
	_PlatformName[26:29]:        PlatformCow,
	_PlatformLowerName[26:29]:   PlatformCow,
	_PlatformName[29:37]:        PlatformArbitrum,
	_PlatformLowerName[29:37]:   PlatformArbitrum,
	_PlatformName[37:44]:        PlatformBendDAO,
	_PlatformLowerName[37:44]:   PlatformBendDAO,
	_PlatformName[44:53]:        PlatformCrossbell,
	_PlatformLowerName[44:53]:   PlatformCrossbell,
	_PlatformName[53:58]:        PlatformCurve,
	_PlatformLowerName[53:58]:   PlatformCurve,
	_PlatformName[58:61]:        PlatformENS,
	_PlatformLowerName[58:61]:   PlatformENS,
	_PlatformName[61:70]:        PlatformFarcaster,
	_PlatformLowerName[61:70]:   PlatformFarcaster,
	_PlatformName[70:79]:        PlatformHighlight,
	_PlatformLowerName[70:79]:   PlatformHighlight,
	_PlatformName[79:85]:        PlatformIQWiki,
	_PlatformLowerName[79:85]:   PlatformIQWiki,
	_PlatformName[85:94]:        PlatformKiwiStand,
	_PlatformLowerName[85:94]:   PlatformKiwiStand,
	_PlatformName[94:98]:        PlatformLens,
	_PlatformLowerName[94:98]:   PlatformLens,
	_PlatformName[98:102]:       PlatformLido,
	_PlatformLowerName[98:102]:  PlatformLido,
	_PlatformName[102:111]:      PlatformLooksRare,
	_PlatformLowerName[102:111]: PlatformLooksRare,
	_PlatformName[111:118]:      PlatformMatters,
	_PlatformLowerName[111:118]: PlatformMatters,
	_PlatformName[118:124]:      PlatformMirror,
	_PlatformLowerName[118:124]: PlatformMirror,
	_PlatformName[124:131]:      PlatformOpenSea,
	_PlatformLowerName[124:131]: PlatformOpenSea,
	_PlatformName[131:139]:      PlatformOptimism,
	_PlatformLowerName[131:139]: PlatformOptimism,
	_PlatformName[139:148]:      PlatformParagraph,
	_PlatformLowerName[139:148]: PlatformParagraph,
	_PlatformName[148:156]:      PlatformParaswap,
	_PlatformLowerName[148:156]: PlatformParaswap,
	_PlatformName[156:160]:      PlatformRSS3,
	_PlatformLowerName[156:160]: PlatformRSS3,
	_PlatformName[160:164]:      PlatformSAVM,
	_PlatformLowerName[160:164]: PlatformSAVM,
	_PlatformName[164:172]:      PlatformStargate,
	_PlatformLowerName[164:172]: PlatformStargate,
	_PlatformName[172:179]:      PlatformUniswap,
	_PlatformLowerName[172:179]: PlatformUniswap,
	_PlatformName[179:182]:      PlatformVSL,
	_PlatformLowerName[179:182]: PlatformVSL,
}

var _PlatformNames = []string{
	_PlatformName[0:7],
	_PlatformName[7:12],
	_PlatformName[12:16],
	_PlatformName[16:26],
	_PlatformName[26:29],
	_PlatformName[29:37],
	_PlatformName[37:44],
	_PlatformName[44:53],
	_PlatformName[53:58],
	_PlatformName[58:61],
	_PlatformName[61:70],
	_PlatformName[70:79],
	_PlatformName[79:85],
	_PlatformName[85:94],
	_PlatformName[94:98],
	_PlatformName[98:102],
	_PlatformName[102:111],
	_PlatformName[111:118],
	_PlatformName[118:124],
	_PlatformName[124:131],
	_PlatformName[131:139],
	_PlatformName[139:148],
	_PlatformName[148:156],
	_PlatformName[156:160],
	_PlatformName[160:164],
	_PlatformName[164:172],
	_PlatformName[172:179],
	_PlatformName[179:182],
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

// MarshalYAML implements a YAML Marshaler for Platform
func (i Platform) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Platform
func (i *Platform) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
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
