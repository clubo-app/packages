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

type Stream struct {
	nc *nats.EncodedConn
}

func new(nc *nats.EncodedConn) Stream {
	return Stream{nc: nc}
}

func (s Stream) Close() {
	s.nc.Close()
}

func Connect(cluster string, opts []nats.Option) (Stream, error) {
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(cluster, opts...)
	if err != nil {
		return Stream{}, err
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return Stream{}, err
	}
	log.Println("Connected to Nats Server at ", c.Conn.ConnectedUrl())
	return new(c), nil
}

func (s Stream) PublishEvent(event any) error {
	sub := eventToSubject(event)
	return s.nc.Publish(sub, event)
}

func (s Stream) SubscribeToEvent(queue string, event any, handler nats.Handler) (*nats.Subscription, error) {
	sub := eventToSubject(event)

	return s.nc.QueueSubscribe(sub, queue, handler)
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
