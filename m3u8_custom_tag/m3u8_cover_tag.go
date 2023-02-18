package m3u8_custom_tag

import (
	"bytes"
	"fmt"
	"github.com/grafov/m3u8"
	"net/url"
)

type MediaCoverTag struct {
	CoverURI *url.URL
}

func (tag *MediaCoverTag) TagName() string {
	return "#MEDIA-COVER:"
}

// line will be the entire matched line, including the identifier
func (tag *MediaCoverTag) Decode(line string) (t m3u8.CustomTag, err error) {
	tagLen := len(tag.TagName())
	// not allow empty value
	if len(line) <= tagLen {
		return nil, fmt.Errorf("invaild tag, expect length large than %d", tagLen)
	}
	coverURI := line[tagLen:]
	u, err := url.Parse(coverURI)
	if err != nil {
		return nil, fmt.Errorf("invaild tag,parse cover uri failed %w", err)
	}
	tag.CoverURI = u
	return tag, nil
}

func (tag *MediaCoverTag) SegmentTag() bool {
	return false
}

func (tag *MediaCoverTag) Encode() *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.WriteString(tag.TagName())
	buf.WriteString(tag.CoverURI.String())
	return buf
}

func (tag *MediaCoverTag) String() string {
	return tag.Encode().String()
}
