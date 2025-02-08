package gradio

type Gradio struct {
	url          string
	hfToken      string
	systemPrompt string
	maxTokens    string
	temperature  string
	topP         string
	apiName      string
}

type ClientOption func(*Gradio)

func NewClient(url string, opts ...ClientOption) *Gradio {
	client := &Gradio{
		url:         url,
		maxTokens:   "1024",
		temperature: "0.8",
		topP:        "0.95",
		apiName:     "/chat",
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithHfToken(token string) ClientOption {
	return func(gr *Gradio) {
		gr.hfToken = token
	}
}

func (gr *Gradio) SetHfToken(token string) *Gradio {
	gr.hfToken = token
	return gr
}

func (gr *Gradio) SetSystemPrompt(prompt string) *Gradio {
	gr.systemPrompt = prompt
	return gr
}

func (gr *Gradio) SetMaxToken(value string) *Gradio {
	gr.maxTokens = value
	return gr
}

func (gr *Gradio) SetTemperature(value string) *Gradio {
	gr.temperature = value
	return gr
}

func (gr *Gradio) SetTopP(value string) *Gradio {
	gr.topP = value
	return gr
}

func (gr *Gradio) SetApiName(path string) *Gradio {
	gr.apiName = path
	return gr
}
