package broker

import (
	"encoding/json"
	"github/SLANGERES/CQRS/Read/database"
	"github/SLANGERES/CQRS/Read/internal/models"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp091.Connection
	db   *database.Storage // your ES wrapper
}

func NewConsumer(url string, db *database.Storage) (*Consumer, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		slog.Error("Unable to connect with MessageBroker", "error", err)
		return nil, err
	}
	return &Consumer{
		conn: conn,
		db:   db,
	}, nil
}

func (c *Consumer) Consume(queueName string, handler func(models.Blog) error) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}

	// Declare queue (idempotent, same as publisher)
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	// Fair dispatch (don’t give more than one unacked message at a time)
	err = ch.Qos(1, 0, false)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queueName,
		"",    // consumer tag
		false, // auto-ack = false
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		slog.Info("Rabbit mq is start consuming")
		for d := range msgs {
			var blog models.Blog
			if err := json.Unmarshal(d.Body, &blog); err != nil {
				slog.Error("Failed to unmarshal blog", "error", err)
				_ = d.Nack(false, false) // reject and don't requeue
				continue
			}

			// Process message
			if err := handler(blog); err != nil {
				slog.Error("Failed to handle blog", "error", err)
				_ = d.Nack(false, true) // requeue for retry
				continue
			}
			slog.Info("blog recived", "blog id", blog.ID)
			if err := c.db.InsertInDb(blog); err != nil {
				slog.Error("unable to process the queue in your db", "error", err)
			}

			// ✅ Ack after success
			if err := d.Ack(false); err != nil {
				slog.Error("Failed to ack message", "error", err)
			}
			slog.Info("blog sync sucessfully")

		}
	}()

	slog.Info("Consumer started", "queue", queueName)
	return nil
}

func (c *Consumer) Close() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
