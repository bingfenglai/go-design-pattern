package abstract_factory_dp

type LocalS3OosFactory struct {
}

func (l *LocalS3OosFactory) UploadImage(bytes []byte) OosImage {
	return &LocalS3Image{}
}

func (l *LocalS3OosFactory) UploadVideo(bytes []byte) OosVideo {
	return &LocalS3Video{}
}
