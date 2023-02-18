package m3u8_custom_tag

import "github.com/grafov/m3u8"

func GetMasterTagSet() []m3u8.CustomDecoder {
	return []m3u8.CustomDecoder{
		&MediaTotalSizeTag{},
		&MediaIdentifierTag{},
		&MediaSourceTag{},
		&MediaSocialMetaTag{},
		&MediaCoverTag{},
	}
}

func GetPlaylistTagSet() []m3u8.CustomDecoder {
	return []m3u8.CustomDecoder{
		&MediaTotalSizeTag{},
		&MediaIdentifierTag{},
	}
}

func GetFullTagSet() []m3u8.CustomDecoder {
	return []m3u8.CustomDecoder{
		&MediaTotalSizeTag{},
		&MediaIdentifierTag{},
		&MediaSourceTag{},
		&MediaSocialMetaTag{},
		&MediaCoverTag{},
	}
}
