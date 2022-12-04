package abstract_factory_dp

type OosImage interface {
	GetLink() string
	GetThumb() string
	GetWatermark() string
}

type OosVideo interface {
	Get720p() string
	Get1080p() string
	GetCoverImgLink() string
}
