package abstract_factory_dp

type AmazonOosImage struct {
}

func (a *AmazonOosImage) GetLink() string {
	return "http://amazon.com/1.png"
}

func (a *AmazonOosImage) GetThumb() string {
	return "http://amazon.com/1_thumb.png"
}

func (a *AmazonOosImage) GetWatermark() string {
	return "http://amazon.com/1_watermark.png"

}

type AmazonVideo struct {
}

func (a *AmazonVideo) Get720p() string {
	return "http://amazon.com/720p.mp4"
}

func (a *AmazonVideo) Get1080p() string {
	return "http://amazon.com/1080p.mp4"
}

func (a *AmazonVideo) GetCoverImgLink() string {
	return "http://amazon.com/cover_img.png"

}
