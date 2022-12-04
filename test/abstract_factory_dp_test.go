package test

import (
	"lbf.com/go-dp/abstract_factory_dp"
	"testing"
)

func TestAbsFactory(t *testing.T) {

	var fa abstract_factory_dp.OosFactory
	fa = &abstract_factory_dp.AmazonOosFactory{}

	Print(t, fa)

	fa = &abstract_factory_dp.LocalS3OosFactory{}

	Print(t, fa)

}

func Print(t *testing.T, fa abstract_factory_dp.OosFactory) {
	image := fa.UploadImage(nil)
	t.Log(image.GetLink())
	t.Log(image.GetThumb())
	t.Log(image.GetWatermark())
	video := fa.UploadVideo(nil)
	t.Log(video.Get720p())
	t.Log(video.Get1080p())
	t.Log(video.GetCoverImgLink())
}
