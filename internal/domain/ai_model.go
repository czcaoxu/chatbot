package domain

type AIModelType string

const (
	ModelLLama3 AIModelType = "llama3"
	ModelGPT4   AIModelType = "gtp4"
	ModelQwen   AIModelType = "Qwen"
	ModelOthers AIModelType = "others"
)

func (a AIModelType) String() string {
	return string(a)
}

func (a AIModelType) IsValid() bool {
	switch a {
	case ModelLLama3, ModelGPT4, ModelQwen, ModelOthers:
		return true
	default:
		return false
	}
}
