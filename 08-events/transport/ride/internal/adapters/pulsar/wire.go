// (pseudo) in internal/adapters/pulsar/wire.go
// sendRaw := func(ctx context.Context, topic, key string, payload []byte, props map[string]string) error { ... }
// pull := func(ctx context.Context) (rawMsg, error) { ... }

// Producer for main topic:
prod := pulsar.NewProducer[avro.AssignmentCreated](pulsar.ProducerConfig[avro.AssignmentCreated]{
	Topic:  "assignments",
	Encode: avroadapter.EncodeAssignment,
	Send:   sendRaw,
})

// Producer for DLQ (same generic type!):
dlq := pulsar.NewProducer[avro.AssignmentCreated](pulsar.ProducerConfig[avro.AssignmentCreated]{
	Topic:  "assignments.DLQ",
	Encode: avroadapter.EncodeAssignment,
	Send:   sendRaw,
})

// Consumer (same reusable logic, different params):
cons := pulsar.NewConsumer[avro.AssignmentCreated](pulsar.ConsumerConfig[avro.AssignmentCreated]{
	Topic:        "assignments",
	Subscription: "assignments-service",
	Decode:       avroadapter.DecodeAssignment,
	Pull:         pull,
	Processor:    application.AssignmentCreatedProcessor{},
	DLQ:          prodToDLQ, // use the generic producer
	MaxRedeliver: 5,
	Parallelism:  4,
})
