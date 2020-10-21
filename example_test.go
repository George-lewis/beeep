package beeep

func ExampleBeep() {
	Beep(DefaultFreq, DefaultDuration)
}

func ExampleNotify() {
	SetAppID("App ID")
	Notify("Title", "MessageBody", "assets/information.png")
}

func ExampleAlert() {
	SetAppID("App ID")
	Alert("Title", "MessageBody", "assets/warning.png")
}
