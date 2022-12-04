package abstract_factory_dp

type AmazonOosFactory struct {
}

func (a *AmazonOosFactory) UploadImage(bytes []byte) OosImage {
	return &AmazonOosImage{}
}

func (a *AmazonOosFactory) UploadVideo(bytes []byte) OosVideo {
	return &AmazonVideo{}
}
