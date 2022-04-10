package ws

import (
	"bytes"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/vlpolak/swtgo/logger"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	chat *Chat
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump() error {
	defer func() {
		c.chat.unregister <- c
		err := c.conn.Close()
		if err != nil {
			logger.ErrorLogger("Closing connection failed", err)
		}
	}()
	c.connectionConfig()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return err
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.chat.broadcast <- message
	}
}

func (c *Client) connectionConfig() {
	c.conn.SetReadLimit(maxMessageSize)
	errdl := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if errdl != nil {
		logger.ErrorLogger("Writing dead line failed", errdl)
	}
	c.conn.SetPongHandler(func(string) error {
		errdl := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if errdl != nil {
			logger.ErrorLogger("Writing dead line failed", errdl)
			return errdl
		}
		return nil
	})
	errwdl := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if errwdl != nil {
		logger.ErrorLogger("Writing dead line failed", errwdl)
	}
}

func (c *Client) writePump() error {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.connectionConfig()
			if !ok {
				errwm := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if errwm != nil {
					logger.ErrorLogger("Writing message failed", errwm)
					return errwm
				}
				err := errors.New("Message sending failed")
				logger.ErrorLogger("Message sending failed", err)
				return err
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.ErrorLogger("Writing message failed", err)
				return err
			}
			_, errm := w.Write(message)
			if errm != nil {
				logger.ErrorLogger("Writing message failed", errm)
				return errm
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				_, errnl := w.Write(newline)
				if errnl != nil {
					logger.ErrorLogger("Writing new line failed", errnl)
					return errnl
				}
				_, errb := w.Write(<-c.send)
				if errb != nil {
					logger.ErrorLogger("Writing channel failed", errb)
					return errb
				}
			}

			if err := w.Close(); err != nil {
				logger.ErrorLogger("Closing connection failed", err)
				return err
			}
		case <-ticker.C:
			err := c.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				logger.ErrorLogger("Writing message failed", err)
				return err
			}
		}
	}
}

func serveWs(chat *Chat, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{chat: chat, conn: conn, send: make(chan []byte, 256)}
	client.chat.register <- client
	go func() {
		err := client.writePump()
		if err != nil {
			logger.ErrorLogger("Writing socket message failed", err)
		}
	}()
	go func() {
		err := client.readPump()
		if err != nil {
			logger.ErrorLogger("Reading socket message failed", err)
		}
	}()
}
