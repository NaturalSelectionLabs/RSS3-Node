// Code generated by "enumer --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql"; DO NOT EDIT.

package decentralized

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _WorkerName = "aaveaavegotchicorecrossbellcurveenshighlightiqwikikiwistandlenslidolooksraremattersmirrormomoka1inchopenseaoptimismparagraphparaswaprss3savmstargateuniswapvsl"

var _WorkerIndex = [...]uint8{0, 4, 14, 18, 27, 32, 35, 44, 50, 59, 63, 67, 76, 83, 89, 95, 100, 107, 115, 124, 132, 136, 140, 148, 155, 158}

const _WorkerLowerName = "aaveaavegotchicorecrossbellcurveenshighlightiqwikikiwistandlenslidolooksraremattersmirrormomoka1inchopenseaoptimismparagraphparaswaprss3savmstargateuniswapvsl"

func (i Worker) String() string {
	i -= 1
	if i < 0 || i >= Worker(len(_WorkerIndex)-1) {
		return fmt.Sprintf("Worker(%d)", i+1)
	}
	return _WorkerName[_WorkerIndex[i]:_WorkerIndex[i+1]]
}

func (Worker) Values() []string {
	return WorkerStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _WorkerNoOp() {
	var x [1]struct{}
	_ = x[Aave-(1)]
	_ = x[Aavegotchi-(2)]
	_ = x[Core-(3)]
	_ = x[Crossbell-(4)]
	_ = x[Curve-(5)]
	_ = x[ENS-(6)]
	_ = x[Highlight-(7)]
	_ = x[IQWiki-(8)]
	_ = x[KiwiStand-(9)]
	_ = x[Lens-(10)]
	_ = x[Lido-(11)]
	_ = x[Looksrare-(12)]
	_ = x[Matters-(13)]
	_ = x[Mirror-(14)]
	_ = x[Momoka-(15)]
	_ = x[Oneinch-(16)]
	_ = x[OpenSea-(17)]
	_ = x[Optimism-(18)]
	_ = x[Paragraph-(19)]
	_ = x[Paraswap-(20)]
	_ = x[RSS3-(21)]
	_ = x[SAVM-(22)]
	_ = x[Stargate-(23)]
	_ = x[Uniswap-(24)]
	_ = x[VSL-(25)]
}

var _WorkerValues = []Worker{Aave, Aavegotchi, Core, Crossbell, Curve, ENS, Highlight, IQWiki, KiwiStand, Lens, Lido, Looksrare, Matters, Mirror, Momoka, Oneinch, OpenSea, Optimism, Paragraph, Paraswap, RSS3, SAVM, Stargate, Uniswap, VSL}

var _WorkerNameToValueMap = map[string]Worker{
	_WorkerName[0:4]:          Aave,
	_WorkerLowerName[0:4]:     Aave,
	_WorkerName[4:14]:         Aavegotchi,
	_WorkerLowerName[4:14]:    Aavegotchi,
	_WorkerName[14:18]:        Core,
	_WorkerLowerName[14:18]:   Core,
	_WorkerName[18:27]:        Crossbell,
	_WorkerLowerName[18:27]:   Crossbell,
	_WorkerName[27:32]:        Curve,
	_WorkerLowerName[27:32]:   Curve,
	_WorkerName[32:35]:        ENS,
	_WorkerLowerName[32:35]:   ENS,
	_WorkerName[35:44]:        Highlight,
	_WorkerLowerName[35:44]:   Highlight,
	_WorkerName[44:50]:        IQWiki,
	_WorkerLowerName[44:50]:   IQWiki,
	_WorkerName[50:59]:        KiwiStand,
	_WorkerLowerName[50:59]:   KiwiStand,
	_WorkerName[59:63]:        Lens,
	_WorkerLowerName[59:63]:   Lens,
	_WorkerName[63:67]:        Lido,
	_WorkerLowerName[63:67]:   Lido,
	_WorkerName[67:76]:        Looksrare,
	_WorkerLowerName[67:76]:   Looksrare,
	_WorkerName[76:83]:        Matters,
	_WorkerLowerName[76:83]:   Matters,
	_WorkerName[83:89]:        Mirror,
	_WorkerLowerName[83:89]:   Mirror,
	_WorkerName[89:95]:        Momoka,
	_WorkerLowerName[89:95]:   Momoka,
	_WorkerName[95:100]:       Oneinch,
	_WorkerLowerName[95:100]:  Oneinch,
	_WorkerName[100:107]:      OpenSea,
	_WorkerLowerName[100:107]: OpenSea,
	_WorkerName[107:115]:      Optimism,
	_WorkerLowerName[107:115]: Optimism,
	_WorkerName[115:124]:      Paragraph,
	_WorkerLowerName[115:124]: Paragraph,
	_WorkerName[124:132]:      Paraswap,
	_WorkerLowerName[124:132]: Paraswap,
	_WorkerName[132:136]:      RSS3,
	_WorkerLowerName[132:136]: RSS3,
	_WorkerName[136:140]:      SAVM,
	_WorkerLowerName[136:140]: SAVM,
	_WorkerName[140:148]:      Stargate,
	_WorkerLowerName[140:148]: Stargate,
	_WorkerName[148:155]:      Uniswap,
	_WorkerLowerName[148:155]: Uniswap,
	_WorkerName[155:158]:      VSL,
	_WorkerLowerName[155:158]: VSL,
}

var _WorkerNames = []string{
	_WorkerName[0:4],
	_WorkerName[4:14],
	_WorkerName[14:18],
	_WorkerName[18:27],
	_WorkerName[27:32],
	_WorkerName[32:35],
	_WorkerName[35:44],
	_WorkerName[44:50],
	_WorkerName[50:59],
	_WorkerName[59:63],
	_WorkerName[63:67],
	_WorkerName[67:76],
	_WorkerName[76:83],
	_WorkerName[83:89],
	_WorkerName[89:95],
	_WorkerName[95:100],
	_WorkerName[100:107],
	_WorkerName[107:115],
	_WorkerName[115:124],
	_WorkerName[124:132],
	_WorkerName[132:136],
	_WorkerName[136:140],
	_WorkerName[140:148],
	_WorkerName[148:155],
	_WorkerName[155:158],
}

// WorkerString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func WorkerString(s string) (Worker, error) {
	if val, ok := _WorkerNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _WorkerNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Worker values", s)
}

// WorkerValues returns all values of the enum
func WorkerValues() []Worker {
	return _WorkerValues
}

// WorkerStrings returns a slice of all String values of the enum
func WorkerStrings() []string {
	strs := make([]string, len(_WorkerNames))
	copy(strs, _WorkerNames)
	return strs
}

// IsAWorker returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Worker) IsAWorker() bool {
	for _, v := range _WorkerValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Worker
func (i Worker) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Worker
func (i *Worker) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Worker should be a string, got %s", data)
	}

	var err error
	*i, err = WorkerString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Worker
func (i Worker) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Worker
func (i *Worker) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = WorkerString(s)
	return err
}

func (i Worker) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Worker) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of Worker: %[1]T(%[1]v)", value)
	}

	val, err := WorkerString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
