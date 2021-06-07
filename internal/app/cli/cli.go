package cli

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

func CliStart(c Config) error {

	// слушатель консоли
	ent := make(chan string)
	go func() {
		for {
			var cmd string
			_, err := fmt.Fscan(os.Stdin, &cmd)
			if err == nil {
				ent <- cmd
			}
		}
	}()

	// читаем из канала и обрабатываем сообщения
	for {
		select {
		case cmd := <-ent:
			switch cmd {
			case "exit":
				log.Println("Завершение работы")
				c.Cancel()
				// case ...
			}
		case <-c.Ctx.Done():
			return nil
		}
	}
}
