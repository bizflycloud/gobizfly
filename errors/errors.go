package gobizflyErr

import (
	"bytes"
	"html/template"
)

var tmpl = template.New("GobizflyTemplate")

type GobizflyErr struct {
	Message  string
	Code     string
	Metadata map[string]interface{}
}

func (err GobizflyErr) Error() string {
	message := err.GetMessage()
	return message
}

func (err GobizflyErr) String() string {
	message := err.GetMessage()
	return message
}

func (err GobizflyErr) GetMessage() string {
	newTmpl, tmplErr := tmpl.Parse(err.Message)
	if tmplErr != nil {
		return err.Message
	}
	var result bytes.Buffer
	tmplErr = newTmpl.Execute(&result, err.Metadata)
	if tmplErr != nil {
		return err.Message
	}
	return result.String()
}

func (err GobizflyErr) SetMetadata(metadata map[string]interface{}) GobizflyErr {
	err.Metadata = metadata
	return err
}

var (
	InvalidRegion = GobizflyErr{
		Message: "Invalid region {{.Region}}",
		Code:    "InvalidRegion",
	}
)
