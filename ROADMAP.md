# Roadmap: Rewriting llm in Go

This document outlines the high-level roadmap for rewriting the `llm` Python library and CLI in Go. The rewrite will be conducted in several phases to ensure a smooth transition and allow for incremental progress.

## Status

**Current Phase:** Phase 1: Core Functionality

**Completed:**
*   Project Scaffolding
*   Model Abstraction
*   Basic CLI Commands (prompt)

**In Progress:**
*   OpenAI Integration

## Phase 1: Core Functionality

The goal of Phase 1 is to replicate the core functionality of the `llm` tool in Go. This includes the ability to run prompts against language models, manage API keys, and log conversations.

*   **1.1: Project Scaffolding**: **DONE**
    *   Set up a new Go project with a clear directory structure.
    *   Define the initial package layout (e.g., `cmd/llm`, `pkg/llm`, `pkg/provider`).
    *   Implement a basic CLI using a library like `cobra`.

*   **1.2: Model Abstraction**: **DONE**
    *   Define a set of Go interfaces for interacting with language models (e.g., `Model`, `EmbeddingModel`).
    *   Implement a `Conversation` struct to manage the history of prompts and responses.
    *   Create a `Response` struct to represent the output from a model.

*   **1.3: OpenAI Integration**: **IN PROGRESS**
    *   **1.3.1: Implement a Go client for the OpenAI API.**
        *   Use the official OpenAI Go client library.
        *   Add the dependency to `go.mod`.
    *   **1.3.2: Update the `OpenAIModel` to use the client.**
        *   The `Execute` method should now call the OpenAI API.
        *   Handle API errors gracefully.
    *   **1.3.3: Support for streaming and non-streaming responses.**
        *   Implement a streaming version of the `Execute` method.
        *   Update the CLI to handle streaming responses.
    *   **1.3.4: Add tests for the `OpenAIModel`.**
        *   Write unit tests for the `Execute` method.
        *   Use a mock HTTP client to avoid making real API calls in tests.

*   **1.4: Key Management**:
    *   Implement a secure way to store and manage API keys, similar to the `llm keys` command.
    *   Store keys in a JSON file in the user's home directory.

*   **1.5: Logging**:
    *   Implement a logging system that stores prompts and responses in a SQLite database.
    *   Use a library like `gorm` or `sqlx` for database interactions.

*   **1.6: Basic CLI Commands**:
    *   Implement the `llm prompt` command for running prompts. **DONE**
    *   Implement the `llm chat` command for interactive conversations.
    *   Implement the `llm keys` command for managing API keys.
    *   Implement the `llm logs` command for viewing conversation logs.

## Phase 2: Extensibility and Plugins

The goal of Phase 2 is to implement a plugin system that allows for extending the `llm` tool with new models, commands, and other features.

*   **2.1: Plugin Architecture**:
    *   Design and implement a plugin system for Go. This could involve using shared libraries, RPC, or a custom plugin protocol.
    *   Define a clear interface for plugins to implement.

*   **2.2: Plugin Hooks**:
    *   Implement a set of hooks that plugins can use to extend the functionality of the `llm` tool, similar to the `pluggy` hooks in the Python version.
    *   Hooks should be available for registering new models, commands, and other features.

*   **2.3: Default Plugins**:
    *   Re-implement the default plugins from the Python version in Go (e.g., `openai_models`, `default_tools`).

## Phase 3: Advanced Features

The goal of Phase 3 is to implement the remaining advanced features of the `llm` tool.

*   **3.1: Embeddings**:
    *   Implement the `llm embed` and `llm similar` commands.
    *   Create a `Collection` struct for managing embeddings in a SQLite database.

*   **3.2: Templates**:
    *   Implement support for prompt templates.
    *   Implement the `llm templates` command for managing templates.

*   **3.3: Tools**:
    *   Implement support for tools that can be executed by the language models.
    *   Implement the `llm tools` command for managing tools.

*   **3.4: Schemas**:
    *   Implement support for JSON schemas for structured output.
    *   Implement the `llm schemas` command for managing schemas.

## Phase 4: Documentation and Release

The goal of Phase 4 is to prepare the Go version of `llm` for its initial release.

*   **4.1: Documentation**:
    *   Write comprehensive documentation for the Go library and CLI.
    *   Include examples and tutorials for all major features.

*   **4.2: Testing**:
    *   Write a comprehensive suite of unit and integration tests.
    *   Ensure that the Go version is as reliable and robust as the Python version.

*   **4.3: Release**:
    *   Package the Go application for distribution on multiple platforms (e.g., binaries, Homebrew, etc.).
    *   Publish the initial release of `llm-go`.
