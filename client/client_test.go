package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestNewClient1(t *testing.T) {
	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)

	if err := c.Set(context.TODO(), "foo ", 1); err != nil {
		log.Fatal(err)
	}

	// val, err := c.Get(context.TODO(), fmt.Sprintf("foo _%d", i))
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func TestNewClient(t *testing.T) {
	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		fmt.Println("SET this =>", fmt.Sprintf("bar _%d", i))

		if err := c.Set(context.TODO(), fmt.Sprintf("foo _%d", i), fmt.Sprintf("bar _%d", i)); err != nil {
			log.Fatal(err)
		}

		val, err := c.Get(context.TODO(), fmt.Sprintf("foo _%d", i))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("GET this =>", val)
	}
}
