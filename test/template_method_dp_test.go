package test

import (
	"lbf.com/go-dp/template_method_dp"
	"testing"
)

func TestTemplateMethod(t *testing.T) {

	fileTemplate := template_method_dp.FileTemplate{FileTemplateOption: &template_method_dp.LocalFileTemplate{}}

	file, err := fileTemplate.UploadFile("../README.md")
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(file)
}
