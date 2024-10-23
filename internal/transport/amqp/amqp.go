package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"teams_service/internal/core/configuration"
	"teams_service/internal/core/dto"

	"github.com/streadway/amqp"
)

type useCase interface {
	UpdateMember(c context.Context, dto *dto.UpdateMember) error
}

type server struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	useCase useCase
}

func New(cfg *configuration.AMQPConfig, useCase useCase) (*server, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:5672/", cfg.User, cfg.Pass, cfg.Host)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return &server{}, err
	}

	ch, err := conn.Channel()
	return &server{
		ch:      ch,
		useCase: useCase,
	}, err
}

func (s *server) Setup() error {
	if err := s.setupUserUpdateEvent(); err != nil {
		return err
	}
	return nil
}

func (s *server) setupUserUpdateEvent() error {
	exchangeName := "user_update"
	err := s.ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	queue, err := s.ch.QueueDeclare("", true, false, false, false, nil)
	if err != nil {
		return err
	}

	if err := s.ch.QueueBind(queue.Name, "", exchangeName, false, amqp.Table{}); err != nil {
		return err
	}

	messages, err := s.ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for message := range messages {
			input := dto.UpdateMember{}

			if err := json.Unmarshal(message.Body, &input); err != nil {
				log.Println(err)
				return
			}

			if err := s.useCase.UpdateMember(context.Background(), &input); err != nil {
				log.Println(err)
				return
			}

			message.Ack(false)
		}
	}()

	return nil
}

func (s *server) Close() {
	s.ch.Close()
	s.conn.Close()
}
