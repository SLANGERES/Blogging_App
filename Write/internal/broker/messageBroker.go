package broker

import (
	"context"
	"encoding/json"
	"github/SLANGERES/CQRS/Write/internal/models"
	"log/slog"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type MqBroker struct {
	conn *amqp091.Connection
}

func NewConnection(url string) (*MqBroker, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		slog.Error("Unable to connect with MessageBroker", "error", err)
		return nil, err
	}

	return &MqBroker{
		conn: conn,
	}, nil
}

func (mq *MqBroker) Publish(message models.ElasticBlog) error {

	ch, err := mq.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare queue (idempotent)
	_, err = ch.QueueDeclare(
		"blog-sync-mq",
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonBlog, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx,
		"",             // exchange
		"blog-sync-mq", // routing key
		false,          // mandatory
		false,          // immediate
		amqp091.Publishing{
			ContentType:  "application/json",
			Body:         jsonBlog,
			DeliveryMode: amqp091.Persistent,
		})
	if err != nil {
		return err
	}

	slog.Info("Published message", "msg", message)
	return nil
}

// Gracefully close connection
func (mq *MqBroker) Close() {
	if mq.conn != nil {
		_ = mq.conn.Close()
	}
}
