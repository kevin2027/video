package mjpeg

import "github.com/kevin2027/video/av"

type CodecData struct {
}

func (d CodecData) Type() av.CodecType {
	return av.MJPEG
}
