package connections

import (
	"github.com/nats-io/stan.go"
	"os"
)

type StanComposite struct {
	SC stan.Conn
}

func ConnectToStan() (*StanComposite, error) {
	sc, err := stan.Connect(
		os.Getenv("NATS_CLUSTER_ID"),
		os.Getenv("NATS_CLIENT_ID"),
		stan.NatsURL(os.Getenv("NATS_URL")),
	)
	if err != nil {
		return nil, err
	}
	return &StanComposite{SC: sc}, nil
}
