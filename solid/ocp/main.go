package main

func main() {
	emailNotificationService := EmailNotificationService{}
	smsNotificationService := SMSNotificationService{}

	// Initially, the notification is sent via email by using EmailNotificationService
	sender := NotificationSender{
		notificationService: emailNotificationService,
	}

	sender.Notify("Hello, world!")
	sender.Notify("")

	// Extending to use SMSNotificationService without modifying the existing code
	sender.notificationService = smsNotificationService

	sender.Notify("Hello, user!")
	sender.Notify("")
}
