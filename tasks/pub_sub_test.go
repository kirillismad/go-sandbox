package tasks

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func expectMessage(t *testing.T, subscriber <-chan string, want string) {
	t.Helper()
	select {
	case got := <-subscriber:
		if got != want {
			t.Fatalf("received %q, want %q", got, want)
		}
	default:
		t.Fatalf("no message available, want %q", want)
	}
}

func expectEmpty(t *testing.T, subscriber <-chan string) {
	t.Helper()
	select {
	case got := <-subscriber:
		t.Fatalf("unexpected message %q", got)
	default:
	}
}

func TestPubSub(t *testing.T) {
	for _, shards := range []int{1, 8, 16} {
		t.Run(fmt.Sprintf("shards=%d", shards), func(t *testing.T) {
			t.Run("SubscriptionsAndTopicIsolation", func(t *testing.T) {
				ps := NewPubSub(shards)
				first := ps.Subscribe("BTC")
				second := ps.Subscribe("BTC")
				other := ps.Subscribe("USD")
				if first == second {
					t.Fatal("subscriptions must return distinct channels")
				}
				for _, subscriber := range []<-chan string{first, second, other} {
					if cap(subscriber) != 10 {
						t.Fatalf("channel capacity = %d, want 10", cap(subscriber))
					}
				}
				ps.Publish("BTC", "update")
				expectMessage(t, first, "update")
				expectMessage(t, second, "update")
				expectEmpty(t, other)
				ps.Publish("USD", "dollar")
				expectMessage(t, other, "dollar")
				expectEmpty(t, first)
				expectEmpty(t, second)
				ps.Publish("missing", "ignored")
				expectEmpty(t, first)
				expectEmpty(t, second)
				expectEmpty(t, other)
			})

			t.Run("FullSubscriberDoesNotBlockOthers", func(t *testing.T) {
				ps := NewPubSub(shards)
				slow := ps.Subscribe("BTC")
				for index := 0; index < 10; index++ {
					ps.Publish("BTC", fmt.Sprint(index))
				}
				fast := ps.Subscribe("BTC")
				published := make(chan struct{})
				go func() {
					ps.Publish("BTC", "overflow")
					close(published)
				}()
				select {
				case <-published:
				case <-time.After(2 * time.Second):
					t.Fatal("Publish blocked on a full subscriber")
				}
				expectMessage(t, fast, "overflow")
				for index := 0; index < 10; index++ {
					expectMessage(t, slow, fmt.Sprint(index))
				}
				expectEmpty(t, slow)
				expectEmpty(t, fast)
				ps.Publish("BTC", "resumed")
				expectMessage(t, slow, "resumed")
				expectMessage(t, fast, "resumed")
			})

			t.Run("EmptyTopicAndMessage", func(t *testing.T) {
				ps := NewPubSub(shards)
				subscriber := ps.Subscribe("")
				ps.Publish("", "")
				expectMessage(t, subscriber, "")
				expectEmpty(t, subscriber)
			})
		})
	}
}

func TestPubSubConcurrentSubscribeAndPublish(t *testing.T) {
	for _, shards := range []int{1, 16} {
		t.Run(fmt.Sprintf("shards=%d", shards), func(t *testing.T) {
			const workers = 100
			ps := NewPubSub(shards)
			topics := []string{"BTC", "USD", "ETH"}
			subscribers := make([][]<-chan string, workers)
			start := make(chan struct{})
			var work sync.WaitGroup
			for worker := 0; worker < workers; worker++ {
				work.Add(2)
				go func() {
					defer work.Done()
					<-start
					local := make([]<-chan string, len(topics))
					for index, topic := range topics {
						local[index] = ps.Subscribe(topic)
					}
					subscribers[worker] = local
				}()
				go func() {
					defer work.Done()
					<-start
					for iteration := 0; iteration < 20; iteration++ {
						for _, topic := range topics {
							ps.Publish(topic, "burst:"+topic)
						}
					}
				}()
			}
			close(start)
			finished := make(chan struct{})
			go func() {
				work.Wait()
				close(finished)
			}()
			select {
			case <-finished:
			case <-time.After(10 * time.Second):
				t.Fatal("concurrent subscriptions and publishing did not finish")
			}
			for _, group := range subscribers {
				for index, subscriber := range group {
					for len(subscriber) > 0 {
						expectMessage(t, subscriber, "burst:"+topics[index])
					}
				}
			}
			for _, topic := range topics {
				ps.Publish(topic, "final:"+topic)
			}
			for _, group := range subscribers {
				for index, subscriber := range group {
					if cap(subscriber) != 10 {
						t.Fatalf("channel capacity = %d, want 10", cap(subscriber))
					}
					expectMessage(t, subscriber, "final:"+topics[index])
					expectEmpty(t, subscriber)
				}
			}
		})
	}
}

func TestHashFNV1a(t *testing.T) {
	sharded := NewShardedSyncMap[int](8)
	for _, test := range []struct {
		input string
		want  uint64
	}{
		{input: "", want: 0xcbf29ce484222325},
		{input: "a", want: 0xaf63dc4c8601ec8c},
		{input: "hello", want: 0xa430d84680aabd0b},
	} {
		t.Run(test.input, func(t *testing.T) {
			if got := sharded.hash(test.input); got != test.want {
				t.Fatalf("hash(%q) = %#x, want %#x", test.input, got, test.want)
			}
		})
	}
}
