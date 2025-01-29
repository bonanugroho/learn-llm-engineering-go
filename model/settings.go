package model

type Settings struct {
	Llm struct {
		Settings struct {
			Openai struct {
				Key   string `json:"key"`
				Url   string `json:"url"`
				Model string `json:"model"`
			} `json:"openai"`
			Ollama struct {
				Key   string `json:"key"`
				Url   string `json:"url"`
				Model string `json:"model"`
			} `json:"ollama"`
			Deepseek struct {
				Key   string `json:"key"`
				Url   string `json:"url"`
				Model string `json:"model"`
			} `json:"deepseek"`
		} `json:"settings"`
		Message struct {
			SystemRole string `json:"system_role"`
			UserRole   string `json:"user_role"`
		} `json:"message"`
	} `json:"llm"`
}
