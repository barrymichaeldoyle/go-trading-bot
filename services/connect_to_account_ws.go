package services

import (
	"encoding/json"
	"fmt"
	"go-trading-bot/config"
	"go-trading-bot/libs"
	"go-trading-bot/tasks"
	"go-trading-bot/types"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/sacOO7/gowebsocket"
)

func sendPingMessage(socket gowebsocket.Socket) {
	out, err := json.Marshal(types.GenericaMessage{Type: "PING"})
	if err != nil {
		panic(err)
	}
	socket.SendText(string(out))
}

func pingServer(socket gowebsocket.Socket) {
	for {
		<-time.After(30 * time.Second)
		go sendPingMessage(socket)
	}
}

func ConnectToAccountWS() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	timestamp := libs.GetCurrentTimestampString()
	socket := gowebsocket.New("wss://api.valr.com/ws/account")
	socket.RequestHeader.Set("X-VALR-API-KEY", config.API_KEY)
	socket.RequestHeader.Set("X-VALR-SIGNATURE", libs.SignRequest(timestamp, "GET", "/ws/account", ""))
	socket.RequestHeader.Set("X-VALR-TIMESTAMP", timestamp)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		fmt.Println("Connected to account ws")
		go pingServer(socket)
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Received account ws connect error ", err)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		fmt.Println("Received account ws binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		fmt.Println("Received account ws ping " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		fmt.Println("Received account ws pong " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		fmt.Println("Disconnected from accounts ws ")
		return
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		genericMessage := &types.GenericaMessage{}
		if err := json.Unmarshal([]byte(message), genericMessage); err != nil {
			panic(err)
		}
		switch genericMessage.Type {
		case "ORDER_STATUS_UPDATE":
			{
				orderStatusUpdate := &types.OrderStatusUpdate{}
				if err := json.Unmarshal([]byte(message), orderStatusUpdate); err != nil {
					panic(err)
				}
				if orderStatusUpdate.Data.OrderStatusType == "Filled" {
					if orderStatusUpdate.Data.OrderSide == "sell" {
						go tasks.RespondToSell(orderStatusUpdate.Data)
					} else if orderStatusUpdate.Data.OrderSide == "buy" {
						go tasks.RespondToBuy(orderStatusUpdate.Data)
					}
				}
			}
		case "PONG":
			fmt.Println("Received PONG")
		default:
		}
	}

	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}
