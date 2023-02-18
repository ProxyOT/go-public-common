package m3u8_custom_tag

import (
	"bytes"
	"fmt"
	"github.com/grafov/m3u8"
	"strconv"
)

type MediaTotalSizeTag struct {
	Size int64
}

func (tag *MediaTotalSizeTag) TagName() string {
	return "#MEDIA-TOTAL-SIZE:"
}

// line will be the entire matched line, including the identifier
func (tag *MediaTotalSizeTag) Decode(line string) (t m3u8.CustomTag, err error) {
	tagLen := len(tag.TagName())
	if len(line) <= tagLen {
		return nil, fmt.Errorf("invaild tag, expect length large than %d", tagLen)
	}
	tag.Size, err = strconv.ParseInt(line[tagLen:], 10, 64)
	return tag, err
}

func (tag *MediaTotalSizeTag) SegmentTag() bool {
	return false
}

func (tag *MediaTotalSizeTag) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)

	buf.WriteString(tag.TagName())
	buf.WriteString(strconv.FormatInt(tag.Size, 10))

	return buf
}

func (tag *MediaTotalSizeTag) String() string {
	return tag.Encode().String()
}
