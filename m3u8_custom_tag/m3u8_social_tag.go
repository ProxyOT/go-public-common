package m3u8_custom_tag

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/grafov/m3u8"
)

type MediaSocialTagData struct {
	Title   string
	Summary string
}

type MediaSocialMetaTag struct {
	Data MediaSocialTagData
}

func (tag *MediaSocialMetaTag) TagName() string {
	return "#MEDIA-SOCIAL:"
}

// line will be the entire matched line, including the identifier
func (tag *MediaSocialMetaTag) Decode(line string) (t m3u8.CustomTag, err error) {
	tagLen := len(tag.TagName())
	data := line[tagLen:]
	// allow empty value
	if len(data) > 0 {
		var (
			titleBytes []byte
		)
		titleBytes, err = base64.StdEncoding.DecodeString(data)
		if err != nil {
			return nil, fmt.Errorf("invaild tag, base64 decode failed: %w", err)
		}
		err = json.Unmarshal(titleBytes, &tag.Data)
		if err != nil {
			return nil, fmt.Errorf("invaild tag, json unmarshal failed: %w", err)
		}
	}

	return tag, nil
}

func (tag *MediaSocialMetaTag) SegmentTag() bool {
	return false
}

func (tag *MediaSocialMetaTag) Encode() *bytes.Buffer {
	titleBytes, err := json.Marshal(tag.Data)
	if err != nil {
		panic(err)
	}
	titleBase64 := base64.StdEncoding.EncodeToString(titleBytes)
	buf := new(bytes.Buffer)
	buf.WriteString(tag.TagName())
	buf.WriteString(titleBase64)
	return buf
}

func (tag *MediaSocialMetaTag) String() string {
	return tag.Encode().String()
}
