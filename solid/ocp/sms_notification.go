package main

import (
	"errors"
	"fmt"
)

type SMSNotificationService struct {
}

func (ns SMSNotificationService) SendNotification(message string) error {
	if message == "" {
		return errors.New("message is required")
	}

	fmt.Printf("Send a notification via SMS: '%s'\n", message)
	return nil
}
