package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Tag --transform=snake --trimprefix=Tag --output tag_string.go --json --sql
type Tag uint64

const (
	TagUnknown Tag = iota
	TagTransaction
	TagCollectible
	TagSocial
)

func TagAndTypeString(tagValue string, typeValue string) (Tag, Type, error) {
	tag, err := TagString(tagValue)
	if err != nil {
		return TagUnknown, nil, err
	}

	_type, err := TypeString(tag, typeValue)
	if err != nil {
		return TagUnknown, nil, err
	}

	return tag, _type, err
}
