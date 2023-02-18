package m3u8_custom_tag

import (
	"testing"
)

func TestNilTagName(t *testing.T) {
	t.Run("TestNilTagName", func(t *testing.T) {
		var (
			mediaTotalSizeTag  *MediaTotalSizeTag
			mediaIdentifierTag *MediaIdentifierTag
			mediaSourceTag     *MediaSourceTag
			mediaSocialMetaTag *MediaSocialMetaTag
		)

		t.Logf("mediaTotalSizeTag tagName of nil poiter: %s", mediaTotalSizeTag.TagName())
		t.Logf("mediaIdentifierTag tagName of nil poiter: %s", mediaIdentifierTag.TagName())
		t.Logf("mediaSourceTag tagName of nil poiter: %s", mediaSourceTag.TagName())
		t.Logf("mediaSocialMetaTag tagName of nil poiter: %s", mediaSocialMetaTag.TagName())

	})
}
