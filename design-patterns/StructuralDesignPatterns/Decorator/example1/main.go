package main

import "fmt"

type notificationResp struct {
	err           error
	destinationId string
}

type notification interface {
	send() notificationResp
}

type smsNotification struct {
	destinationId string
	msg           string
}

func (sn smsNotification) send() notificationResp {
	fmt.Println("Sending SMS notification")
	fmt.Println(sn.msg)
	// This if is just to simulate the usage of destination id,
	// in this case to get the phone number
	if len(sn.destinationId) > 0 {
		fmt.Println("delivered to 111111111111")
	}
	return simulateResponse(sn.destinationId)
}

// defining code to handle email and slack notifications
type slackNotification struct {
	notification notification
}

func (s slackNotification) send() notificationResp {
	resp := s.notification.send()
	if len(resp.destinationId) > 0 {
		slackUser := "slack-user-sq"
		fmt.Println("sending to slack:", slackUser)
	}
	return simulateResponse(resp.destinationId)
}

func main() {
	smS := smsNotification{
		destinationId: "1234",
		msg:           "message for sms",
	}
	slack := slackNotification{notification: smS}
	slack.send()
}

func simulateResponse(dest string) notificationResp {
	return notificationResp{
		nil, dest,
	}
}
