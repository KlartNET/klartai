package llama

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
	"sync"
	"bytes"
)
import (
	"klartai/src/models"
)

type LLamaManager struct {
	mu sync.Mutex
	activeCmd *exec.Cmd
	currentCode string
}



var Manager = &LLamaManager{}

func (em *LLamaManager) SwitchModel(modelCode string) error {
	em.mu.Lock()
	defer em.mu.Unlock()

	if em.currentCode == modelCode && em.activeCmd != nil {
		return nil
	}

	if em.activeCmd != nil && em.activeCmd.Process != nil {
		fmt.Printf("Stopping current model: %s\n", em.currentCode)

		em.activeCmd.Process.Kill()
		em.activeCmd.Wait()
		em.activeCmd = nil
	}

	model, exists := models.Registry[modelCode]
	if !exists {
		return fmt.Errorf("Model not found: %s\n", modelCode)
	}

	em.activeCmd = exec.Command(
		"./llama-cpp/llama-server",
		"--model", model.RunOption.Path,
		"--threads", "12",
		"--ctx-size", model.RunOption.ContextSize,
		"--n-gpu-layers", "999",
		"--log-verbosity", "2",
		"--no-webui",
	)
	em.activeCmd.Stdout = os.Stdout
	em.activeCmd.Stderr = os.Stderr

	err := em.activeCmd.Start()
	if err != nil {
		return fmt.Errorf("Failed to start model: %s\n", modelCode)
	}

	em.currentCode = modelCode

	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)

		resp, err := http.Get("http://localhost:8080/health")
		if err == nil {
			resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				fmt.Printf("Model '%s' is ready.\n", modelCode)

				return nil
			}
		}
	}

	return fmt.Errorf("Failed to start model: %s\n", modelCode)
}


func Post(body []byte) (*http.Response, error) {
	return http.Post(
		"http://localhost:8080/v1/chat/completions",
		"application/json",
		bytes.NewBuffer(body),
	)
}