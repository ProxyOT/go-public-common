package m3u8_custom_tag

import (
	"bytes"
	"fmt"
	"github.com/grafov/m3u8"
)

type MediaIdentifierTag struct {
	Identifier string
}

func (tag *MediaIdentifierTag) TagName() string {
	return "#MEDIA-IDENTIFIER:"
}

// line will be the entire matched line, including the identifier
func (tag *MediaIdentifierTag) Decode(line string) (t m3u8.CustomTag, err error) {
	tagLen := len(tag.TagName())
	// not allow empty value
	if len(line) <= tagLen {
		return nil, fmt.Errorf("invaild tag, expect length large than %d", tagLen)
	}
	tag.Identifier = line[tagLen:]
	return tag, nil
}

func (tag *MediaIdentifierTag) SegmentTag() bool {
	return false
}

func (tag *MediaIdentifierTag) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString(tag.TagName())
	buf.WriteString(tag.Identifier)
	return buf
}

func (tag *MediaIdentifierTag) String() string {
	return tag.Encode().String()
}
