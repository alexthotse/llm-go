# Implementation Plan: Rewriting llm in Go

This document provides a detailed implementation plan for rewriting the `llm` Python library and CLI in Go. It builds upon the high-level plan outlined in `ROADMAP.md`.

## 1. Package Structure

The Go project will be organized into the following packages:

```
llm-go/
├── cmd/llm/
│   └── main.go
├── pkg/
│   ├── llm/
│   │   ├── llm.go
│   │   ├── conversation.go
│   │   ├── response.go
│   │   └── ...
│   ├── provider/
│   │   ├── openai/
│   │   │   └── openai.go
│   │   └── ...
│   ├── store/
│   │   ├── keys.go
│   │   └── logs.go
│   └── ...
└── ...
```

*   **`cmd/llm`**: The main application package, responsible for parsing command-line arguments and orchestrating the execution of commands.
*   **`pkg/llm`**: The core library package, containing the main `LLM` struct, as well as the `Model`, `Conversation`, and `Response` interfaces and implementations.
*   **`pkg/provider`**: A package containing sub-packages for each supported language model provider (e.g., `openai`, `anthropic`, etc.). Each provider package will implement the `llm.Model` interface.
*   **`pkg/store`**: A package for managing data persistence, including API keys and conversation logs.

## 2. Data Models and Interfaces

The following Go interfaces and structs will be defined to model the core concepts of the `llm` library.

### 2.1. `llm.Model`

The `Model` interface will be the central abstraction for interacting with language models.

```go
package llm

type Model interface {
    Name() string
    Execute(prompt *Prompt, options *Options) (*Response, error)
}

type EmbeddingModel interface {
    Model
    Embed(text string) ([]float64, error)
}

type Prompt struct {
    System string
    Prompt string
    // ... and other fields
}

type Options struct {
    // ... model-specific options
}
```

### 2.2. `llm.Conversation`

The `Conversation` struct will manage the history of prompts and responses.

```go
package llm

type Conversation struct {
    ID        string
    Model     Model
    History   []*Response
    // ... and other fields
}
```

### 2.3. `llm.Response`

The `Response` struct will represent the output from a model.

```go
package llm

type Response struct {
    ID           string
    Prompt       *Prompt
    Text         string
    Model        Model
    Conversation *Conversation
    // ... and other fields
}
```

## 3. Key Features

### 3.1. CLI

The CLI will be implemented using the `cobra` library, which provides a powerful and flexible framework for building command-line applications in Go.

### 3.2. Plugin System

The plugin system will be a key feature of the Go implementation. Since Go does not have a direct equivalent to Python's entry points, we will need to explore alternative approaches. Possible options include:

*   **Shared Libraries**: Plugins could be distributed as shared libraries (`.so` files) that are loaded at runtime. This would require using `cgo` and would introduce some complexity.
*   **RPC**: Plugins could be implemented as separate processes that communicate with the main `llm` application via RPC. This would provide a high degree of isolation, but would also introduce performance overhead.
*   **Custom Plugin Protocol**: We could define a custom plugin protocol based on standard input/output. The `llm` application would spawn plugin processes and communicate with them over `stdin` and `stdout`.

The custom plugin protocol approach is the most promising, as it would be the most portable and easiest to implement.

### 3.3. Logging

Logging will be implemented using a SQLite database, similar to the Python version. The `gorm` library will be used for database interactions, as it provides a convenient and type-safe way to work with SQLite.

## 4. Phase 1 Implementation Details

### 4.1. Project Scaffolding

**Status: DONE**

*   Initialize a new Go module: `go mod init github.com/user/llm-go`
*   Create the initial directory structure.
*   Set up a basic `main.go` file with a `cobra` root command.

### 4.2. Model Abstraction

**Status: DONE**

*   Define the `Model`, `EmbeddingModel`, `Prompt`, and `Options` interfaces and structs in the `pkg/llm` package.

### 4.3. OpenAI Integration

**Status: IN PROGRESS**

*   Create a new `pkg/provider/openai` package.
*   Implement an `OpenAIModel` struct that implements the `llm.Model` interface.
*   Use the official OpenAI Go client library for interacting with the API.

#### 4.3.1. `OpenAIModel` Struct

The `OpenAIModel` struct will have the following fields:

```go
package openai

import (
    "github.com/sashabaranov/go-openai"
)

type OpenAIModel struct {
    name   string
    client *openai.Client
}
```

#### 4.3.2. `Execute` Method

The `Execute` method will be responsible for the following:
1.  Creating a chat completion request from the `llm.Prompt`.
2.  Calling the OpenAI API.
3.  Parsing the response and returning an `llm.Response`.

#### 4.3.3. API Key Management

The OpenAI API key will be read from the `OPENAI_API_KEY` environment variable. In a later step, we will implement a more robust key management system.

### 4.4. Key Management

**Status: NOT STARTED**

*   Create a `pkg/store` package.
*   Implement a `KeyStore` struct for managing API keys.
*   Store keys in a JSON file at `~/.llm/keys.json`.

### 4.5. Logging

**Status: NOT STARTED**

*   Implement a `LogStore` struct in the `pkg/store` package.
*   Use `gorm` to create and manage a SQLite database at `~/.llm/logs.db`.
*   Define a `LogEntry` struct for storing prompts and responses.

### 4.6. Basic CLI Commands

**Status: IN PROGRESS**

*   Implement the `prompt` command in `cmd/llm/prompt.go`. **DONE**
*   Implement the `chat` command in `cmd/llm/chat.go`.
*   Implement the `keys` command in `cmd/llm/keys.go`.
*   Implement the `logs` command in `cmd/llm/logs.go`.
