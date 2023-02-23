package main

import "fmt"

/*
	Nehmen wir das Beispiel der One Time Password (OTP)-Funktionalität.

	Es gibt verschiedene Möglichkeiten, wie das OTP an einen Nutzer übermittelt werden kann (SMS, E-Mail usw.).
	Aber unabhängig davon, ob es sich um ein SMS- oder E-Mail-OTP handelt, ist der gesamte OTP-Prozess derselbe:

	* Generierung einer n-stelligen Zufallszahl
	* Speichern dieser Zahl im Zwischenspeicher zur späteren Überprüfung
	* Bereitet den Inhalt vor
	* Senden der Benachrichtigung
	* Alle neuen OTP-Typen, die in Zukunft eingeführt werden, durchlaufen höchstwahrscheinlich weiterhin die oben genannten Schritte

	Ich habe also ein Szenario, in dem die Schritte eines bestimmten Vorgangs die gleichen sind, aber die
	Implementierungsschritte unterschiedlich sein können.

	Dies ist eine geeignete Situation, um das Template-Methodenmuster zu verwenden.

	Zunächst definiere ich einen Basis-Template-Algorithmus, der aus einer festen Anzahl von Methoden besteht.
	Das wird meine Template-Methode sein.

	Dann implementiere ich jede der Schrittmethoden, lasse aber die Vorlagenmethode unverändert.

*/

// Template method

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

// Concrete Implementation

type Sms struct {
	Otp
}

func (s *Sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type Email struct {
	Otp
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "4321"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}

// Client code

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
