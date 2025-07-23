package utils

var EventChan chan []byte

func InitEventChannel() {
	EventChan = make(chan []byte)
}

func AddEventToChannel(eventJson []byte) {
	EventChan <- eventJson
}
