package doppler_test

import (
	"fmt"
	"log"
	"testing"
)

func BenchmarkDopplerLatencyV1ToV1_1000(b *testing.B) {
	log.Println("Starting Doppler latency benchmark (V1ToV1_1000)...")
	defer log.Println("Done with Doppler latency benchmark (V1ToV1_1000).")
	producer := newV1Producer(grpcConfig)
	consumer := newV1Consumer(grpcConfig)
	defer producer.closeSend()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			msg := fmt.Sprintf("message-%d", i)
			payload := buildPayload([]byte(msg))
			producer.send(payload)
		}
		// Assume Doppler drops some envelopes.
		consumer.observe(900)
	}
	b.StopTimer()
}

func BenchmarkDopplerLatencyV1ToV1_1(b *testing.B) {
	log.Println("Starting Doppler latency benchmark (V1ToV1_1)...")
	defer log.Println("Done with Doppler latency benchmark (V1ToV1_1).")
	producer := newV1Producer(grpcConfig)
	consumer := newV1Consumer(grpcConfig)
	defer producer.closeSend()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msg := fmt.Sprintf("message-%d", i)
		payload := buildPayload([]byte(msg))
		producer.send(payload)
		consumer.observe(1)
	}
	b.StopTimer()
}