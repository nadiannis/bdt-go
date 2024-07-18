package main

import "fmt"

type NotificationSender struct {
	notificationService NotificationService
}

func (s *NotificationSender) Notify(message string) {
	err := s.notificationService.SendNotification(message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
