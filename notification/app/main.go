package main

import (
	// "net"
	// "strconv"

	"log"
	// "time"
	"fmt"
	"context"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9094"},
		Topic:     "notification-message-group",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	// r.SetOffset(2)
	
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		log.Println("RUNNN")
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
	
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

	// topic := "my-topic"
	// partition := 0
	
	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9094", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }
	
	// conn.SetReadDeadline(time.Now().Add(10*time.Second))
	// batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	
	// b := make([]byte, 10e3) // 10KB max per message
	// for {
	// 	n, err := batch.Read(b)
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(string(b[:n]))
	// }
	
	// if err := batch.Close(); err != nil {
	// 	log.Fatal("failed to close batch:", err)
	// }
	
	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close connection:", err)
	// }
}