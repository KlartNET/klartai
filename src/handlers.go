package src

import (
	"io"
	"net/http"
	"encoding/json"

	"github.com/labstack/echo/v4"
)
import (
	"klartai/src/llama"
	"klartai/src/types"
	"klartai/src/models"
)

type LLamaRequest struct {
	Messages	[]types.Message	`json:"messages"`
	Temperature	float64			`json:"temperature"`
	TopK		int				`json:"top_k"`
	TopP		float64			`json:"top_p"`
	MinP		float64			`json:"min_p"`
}

func handleGetModels(ctx echo.Context) error {
	type Response struct {
		Code		string		`json:"code"`
		DisplayName string		`json:"displayName"`
		Source		string		`json:"source"`
		ContextSize string		`json:"contextSize"`
		Parameters  string		`json:"parameters"`
		Tags		[]string	`json:"tags,omitempty"`
	}

	modelList := make([]Response, 0, len(models.Registry))

	for code, m := range models.Registry {
		modelList = append(modelList, Response{
			Code:		 code,
			DisplayName: m.Metadata.DisplayName,
			ContextSize: m.RunOption.ContextSize,
			Source:		 m.Metadata.Source,
			Parameters:  m.Metadata.Parameters,
			Tags: 		 m.Metadata.Tags,
		})
	}

	return ctx.JSON(http.StatusOK, modelList)
}

func handleChat(ctx echo.Context) error {
	type Request struct {
		Model		string		`json:"model"`
		Messages	[]types.Message	`json:"messages"`
	}

	request := new(Request)
	if err := ctx.Bind(request); err != nil {
		return ctx.String(
			http.StatusBadRequest,
			"Bad request",
		)
	}

	err := llama.Manager.SwitchModel(request.Model)
	if err != nil {
		return ctx.String(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	model := models.Registry[request.Model]

	finalMessages := append(model.RunOption.System, request.Messages...)
	openAIReq := LLamaRequest{
		Messages: finalMessages,
		Temperature: model.Tuning.Temperature,
		TopK: model.Tuning.TopK,
		TopP: model.Tuning.TopP,
		MinP: model.Tuning.MinP,
	}

	jsonData, err := json.Marshal(openAIReq)
	if err != nil {
		return ctx.String(
			http.StatusInternalServerError,
			"Payload failure",
		)
	}

	resp, err := llama.Post(jsonData)
	if err != nil {
		return ctx.String(
			http.StatusInternalServerError,
			"No respond",
		)
	}
	defer resp.Body.Close()

	
	writer := ctx.Response().Writer
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().WriteHeader(resp.StatusCode)
	
	if _, err := io.Copy(writer, resp.Body); err != nil {
		return err
	}
	return nil
}
