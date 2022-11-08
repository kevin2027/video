package format

import (
	"github.com/kevin2027/video/av/avutil"
	"github.com/kevin2027/video/format/aac"
	"github.com/kevin2027/video/format/flv"
	"github.com/kevin2027/video/format/mp4"
	"github.com/kevin2027/video/format/rtmp"
	"github.com/kevin2027/video/format/rtsp"
	"github.com/kevin2027/video/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
