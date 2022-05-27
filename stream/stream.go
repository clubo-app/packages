package stream

import (
	"log"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
)

type stream struct {
	nats *nats.EncodedConn
}

type Stream interface {
	PublishEvent(event any) error
	SubscribeToEvent(queue string, event any, handler nats.Handler) (*nats.Subscription, error)
}

func New(nats *nats.EncodedConn) Stream {
	return &stream{nats: nats}
}

func Connect(cluster string, opts []nats.Option) (*nats.EncodedConn, error) {
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(cluster, opts...)
	if err != nil {
		return nil, err
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to Nats Server at ", c.Conn.ConnectedUrl())
	return c, nil
}

func (s stream) PublishEvent(event any) error {
	sub := eventToSubject(event)
	return s.nats.Publish(sub, event)
}

func (s stream) SubscribeToEvent(queue string, event any, handler nats.Handler) (*nats.Subscription, error) {
	sub := eventToSubject(event)

	return s.nats.QueueSubscribe(sub, queue, handler)
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}

func eventToSubject(event any) string {
	t := reflect.TypeOf(event)

	str := strings.ReplaceAll(t.String(), "*", "")

	s := strings.Split(str, ".")

	// if type is events.ImportantType, remove events prefix from string
	if len(s) == 2 && s[0] == "events" {
		return camelcaseStringToDotString(s[1])
	}

	return camelcaseStringToDotString(t.String())
}

func camelcaseStringToDotString(camelcase string) string {
	var b strings.Builder

	for i, c := range camelcase {
		if unicode.IsUpper(c) {
			if i != 0 {
				b.WriteString(".")
			}
			b.WriteRune(unicode.ToLower(c))
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}
