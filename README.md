# GoGradio
This package is a simple sdk to connect your Gradio api hosted on HuggingFace Spaces in go.
Only support simple chat completion without session (history memory).

## Getting started
### Installation
1. Require [Golang 1.23](https://go.dev/doc/install).
2. To use this package run command below.
```sh
go get github.com/xncs120/go-gradio@latest
```

### Initial code
1. Use code below to initiate gradio client and call to your Gradio api hosted on HuggingFace Spaces.
2. Replace the correct url link to your HuggingFace Spaces and token of HuggingFace.
```sh
import (
    "fmt"

    "github.com/xncs120/go-gradio"
)

gr := gradio.NewClient(
    "https://username-space.hf.space/gradio_api/call/chat",
    gradio.WithHfToken("HF_TOKEN"),
)
response := gr.ChatCompletion("hello")
fmt.Println(response)
```

3. Able to write in oop style
```sh
gr := gradio.NewClient("https://username-space.hf.space/gradio_api/call/chat").
    SetHfToken("HF_TOKEN").
    SetMaxToken("2048")
response := gr.ChatCompletion("hello")
fmt.Println(response)
```

## Reference and external source
- [Gradio](https://github.com/gradio-app/gradio)
