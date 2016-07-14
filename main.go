package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/sebgl/chatbox/box"
	"github.com/sebgl/chatbox/chat"
)

var (
	room      = flag.String("room", "", "Room to relax and chat in")
	username  = flag.String("username", "Bernard Minet", "User name")
	brokerURL = flag.String("broker", "", "Kafka broker URL")
	clientID  = flag.String("client-id", "", "Kafka client ID")
)

func main() {
	flag.Parse()
	//log.SetLevel(log.DebugLevel)

	config := &chat.Config{
		Room:      *room,
		Username:  *username,
		BrokerURL: *brokerURL,
		ClientID:  *clientID,
	}
	chatter, err := chat.New(config)
	if err != nil {
		log.WithError(err).Error("Fail to start chatbox")
		os.Exit(1)
	}
	msgIn, textOut := chatter.Start()

	layout := box.NewLayout(msgIn, textOut)
	layout.Init()
}
