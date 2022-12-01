package template_method_dp

type FileTemplateOption interface {
	LoadFile(filePath string) (string, error)
	SaveFile(filePath string) (string, error)
}

type FileTemplate struct {
	FileTemplateOption
}

func (receiver FileTemplate) UploadFile(filePath string) (string, error) {
	var err error
	if filePath, err = receiver.LoadFile(filePath); err == nil {

		if filePath, err = receiver.SaveFile(filePath); err == nil {
			return filePath, nil
		}
	}

	return "", err
}
