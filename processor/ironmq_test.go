package processor_test

import (
	"testing"

	"gopkg.in/queue.v1"
	"gopkg.in/queue.v1/ironmq"

	"github.com/iron-io/iron_go3/mq"
)

func TestIronmqProcessor(t *testing.T) {
	testProcessor(t, ironmq.NewQueue(mq.New("test-ironmq-processor"), &queue.Options{}))
}

func TestIronmqDelay(t *testing.T) {
	testDelay(t, ironmq.NewQueue(mq.New("test-ironmq-delay"), &queue.Options{}))
}

func TestIronmqRetry(t *testing.T) {
	testRetry(t, ironmq.NewQueue(mq.New("test-ironmq-retry"), &queue.Options{}))
}

func TestIronmqNamedMessage(t *testing.T) {
	testNamedMessage(t, ironmq.NewQueue(mq.New("test-ironmq-named-message"), &queue.Options{
		Redis: redisRing(),
	}))
}

func TestIronmqCallOnce(t *testing.T) {
	testCallOnce(t, ironmq.NewQueue(mq.New("test-ironmq-call-once"), &queue.Options{
		Redis: redisRing(),
	}))
}

func TestIronmqRateLimit(t *testing.T) {
	testRateLimit(t, ironmq.NewQueue(mq.New("test-ironmq-rate-limit"), &queue.Options{}))
}

func TestIronmqDelayer(t *testing.T) {
	testDelayer(t, ironmq.NewQueue(mq.New("test-ironmq-delayer"), &queue.Options{}))
}
