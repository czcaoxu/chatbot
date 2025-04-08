package lib

type AIModel string

const (
	ModelLLama3 AIModel = "llama3"
	ModelGPT4   AIModel = "gtp4"
	ModelQwen   AIModel = "Qwen"
	ModelOthers AIModel = "others"
)
