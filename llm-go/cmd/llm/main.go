package main

import (
	"fmt"
	"os"

	"github.com/simonw/llm-go/pkg/llm"
	"github.com/simonw/llm-go/pkg/provider/openai"
	"github.com/spf13/cobra"
)

var model string

var rootCmd = &cobra.Command{
	Use:   "llm",
	Short: "A Go implementation of llm",
	Long:  `A Go implementation of the llm command-line tool.`,
}

var promptCmd = &cobra.Command{
	Use:   "prompt [prompt]",
	Short: "Execute a prompt",
	Long:  `Execute a prompt against a language model.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		promptText := args[0]

		var m llm.Model
		if model != "" {
			m = openai.NewOpenAIModel(model)
		} else {
			m = openai.NewOpenAIModel("gpt-3.5-turbo")
		}

		prompt := &llm.Prompt{
			Prompt: promptText,
		}

		response, err := m.Execute(prompt, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(response.Text)
	},
}

func init() {
	promptCmd.Flags().StringVarP(&model, "model", "m", "", "Model to use")
	rootCmd.AddCommand(promptCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
