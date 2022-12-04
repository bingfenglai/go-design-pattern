package abstract_factory_dp

type OosFactory interface {
	UploadImage([]byte) OosImage
	UploadVideo([]byte) OosVideo
}
