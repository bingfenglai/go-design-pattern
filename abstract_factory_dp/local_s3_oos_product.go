package abstract_factory_dp

type LocalS3Image struct {
}

func (l *LocalS3Image) GetLink() string {
	return "http://lbf.com/1.png"
}

func (l *LocalS3Image) GetThumb() string {
	return "http://lbf.com/1_thumb.png"

}

func (l *LocalS3Image) GetWatermark() string {
	return "http://lbf.com/1_watermark.png"
}

type LocalS3Video struct {
}

func (l *LocalS3Video) Get720p() string {
	return "http://lbf.com/1_720p.mp4"
}

func (l *LocalS3Video) Get1080p() string {
	return "http://lbf.com/1_1080p.mp4"
}

func (l *LocalS3Video) GetCoverImgLink() string {
	return "http://lbf.com/1_coverimg.png"
}
