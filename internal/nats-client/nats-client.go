package natsclient

import (
	"encoding/json"
	"fmt"
	"klambri-nats-publisher/internal/models"

	"github.com/nats-io/nats.go"
)

func Run() error {
	nc, err := connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	defer nc.Close()

	playbookConfig := models.PlaybookConfig{
		PlaybookName: "../playbooks/hello-world.yaml",
		Hosts:        []string{"127.0.0.1", "127.0.0.1"},
	}

	data, err := json.Marshal(playbookConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal playbook config: %w", err)
	}

	publish(nc, "playbookQueue", data)

	fmt.Println("publish!")

	return nil
}
