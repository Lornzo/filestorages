package filestoragetypes

func GetExtensionFromMineType(m MineType) Extension {

	switch m {
	case MINE_TYPE_IMAGE_JPEG:
		return EXTENSION_IMAGE_JPEG
	case MINE_TYPE_IMAGE_JPG:
		return EXTENSION_IMAGE_JPG
	case MINE_TYPE_IMAGE_GIF:
		return EXTENSION_IMAGE_GIF
	case MINE_TYPE_IMAGE_PNG:
		return EXTENSION_IMAGE_PNG
	}

	return ""
}
