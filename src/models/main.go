package models

import (
	"fmt"
	"klartai/src/types"
)

type RunOption struct {
	Path string
	ContextSize string
	System []types.Message
}
type Tuning struct {
	Temperature float64
	TopK int
	TopP float64
	MinP float64
}
type Metadata struct {
	DisplayName string
	Source string
	Parameters string
	Tags []string
}
type ModelInfo struct {
	RunOption RunOption
	Tuning Tuning
	Metadata Metadata
}


var defaultSystem = "You are a helpful AI assistant. Provide brief, concise responses."

var Registry = map[string]ModelInfo{
	"gemma-4-e2b": {
		RunOption: RunOption{
			Path: "./models/gemma4-E2B-Q5_K_M.gguf",
			ContextSize: fmt.Sprint(1024 * 8),
			System: []types.Message{
				{
					Role: "system",
					Content: defaultSystem,
				},
			},
		},
		Tuning: Tuning{
			Temperature: 0.6,
			TopK: 40,
			TopP: 0.95,
			MinP: 0.05,
		},
		Metadata: Metadata{
			DisplayName: "Gemma 4 E2B",
			Source: "Google DeepMind",
			Parameters: "2.3B",
		},
	},
	
	"gemma-4-e4b": {
		RunOption: RunOption{
			Path: "./models/gemma4-E4B-Q4_K_M.gguf",
			ContextSize: fmt.Sprint(1024 * 2),
			System: []types.Message{
				{
					Role: "system",
					Content: defaultSystem,
				},
			},
		},
		Tuning: Tuning{
			Temperature: 0.6,
			TopK: 40,
			TopP: 0.95,
			MinP: 0.05,
		},
		Metadata: Metadata{
			DisplayName: "Gemma 4 E4B",
			Source: "Google DeepMind",
			Parameters: "4.5B",
		},
	},

	"gemma-4-e2b-uncensored": {
		RunOption: RunOption{
			Path: "./models/gemma4-E2B-uncensored-HauhauCS-Q4_K_P.gguf",
			ContextSize: fmt.Sprint(1024 * 8),
			System: []types.Message{
				{
					Role: "system",
					Content: defaultSystem,
				},
			},
		},
		Tuning: Tuning{
			Temperature: 0.6,
			TopK: 40,
			TopP: 0.95,
			MinP: 0.05,
		},
		Metadata: Metadata{
			DisplayName: "Gemma 4 E2B",
			Source: "Google DeepMind",
			Parameters: "2.3B",
			Tags: []string{"Uncensored", "Aggressive Tuning"},
		},
	},

	"gemma-4-e4b-uncensored": {
		RunOption: RunOption{
			Path: "./models/gemma4-E4B-uncensored-HauhauCS-Q4_K_M.gguf",
			ContextSize: fmt.Sprint(1024 * 2),
			System: []types.Message{
				{
					Role: "system",
					Content: defaultSystem,
				},
			},
		},
		Tuning: Tuning{
			Temperature: 0.6,
			TopK: 40,
			TopP: 0.95,
			MinP: 0.05,
		},
		Metadata: Metadata{
			DisplayName: "Gemma 4 E4B",
			Source: "Google DeepMind",
			Parameters: "4.5B",
			Tags: []string{"Uncensored", "Aggressive Tuning"},
		},
	},

	"exaone-3.5-q4": {
		RunOption: RunOption{
			Path: "./models/EXAONE3.5-2.4B-Q4_K_M.gguf",
			ContextSize: fmt.Sprint(1024 * 4),
			System: []types.Message{
				{
					Role: "system",
					Content: fmt.Sprintf("%s Never use emojis. Must answer in Korean", defaultSystem),
				},
			},
		},
		Tuning: Tuning{
			Temperature: 1.8,
			TopK: 40,
			TopP: 1.0,
			MinP: 0.0,
		},
		Metadata: Metadata{
			DisplayName: "EXAONE 3.5",
			Source: "EXAONE",
			Parameters: "2.4B",
			Tags: []string{"Hallucination Tuning"},
		},
	},
}