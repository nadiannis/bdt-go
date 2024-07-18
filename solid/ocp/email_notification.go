package main

import (
	"errors"
	"fmt"
)

type EmailNotificationService struct {
}

func (ns EmailNotificationService) SendNotification(message string) error {
	if message == "" {
		return errors.New("message is required")
	}

	fmt.Printf("Send a notification via email: '%s'\n", message)
	return nil
}
