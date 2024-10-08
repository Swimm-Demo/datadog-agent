// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux_bpf

package protocols

// #cgo CFLAGS: -I ../../ebpf/c  -I ../ebpf/c
// #include "../ebpf/c/protocols/classification/defs.h"
import "C"

import "github.com/DataDog/datadog-agent/pkg/util/log"

const (
	layerAPIBit         = C.LAYER_API_BIT
	layerApplicationBit = C.LAYER_APPLICATION_BIT
	layerEncryptionBit  = C.LAYER_ENCRYPTION_BIT
)

// DispatcherProgramType is a C type to represent the eBPF programs used for tail calls.
type DispatcherProgramType C.dispatcher_prog_t

const (
	// DispatcherKafkaProg is the Golang representation of the C.DISPATCHER_KAFKA_PROG enum.
	DispatcherKafkaProg DispatcherProgramType = C.DISPATCHER_KAFKA_PROG
)

// ProgramType is a C type to represent the eBPF programs used for tail calls.
type ProgramType C.protocol_prog_t

const (
	// ProgramHTTP is the Golang representation of the C.PROG_HTTP enum
	ProgramHTTP ProgramType = C.PROG_HTTP
	// ProgramHTTPTermination is tail call to process http termination.
	ProgramHTTPTermination ProgramType = C.PROG_HTTP_TERMINATION
	// ProgramHTTP2HandleFirstFrame is the Golang representation of the C.PROG_HTTP2_HANDLE_FIRST_FRAME enum
	ProgramHTTP2HandleFirstFrame ProgramType = C.PROG_HTTP2_HANDLE_FIRST_FRAME
	// ProgramHTTP2FrameFilter is the Golang representation of the C.PROG_HTTP2_HANDLE_FRAME enum
	ProgramHTTP2FrameFilter ProgramType = C.PROG_HTTP2_FRAME_FILTER
	// ProgramHTTP2HeadersParser is the Golang representation of the C.PROG_HTTP2_HEADERS_PARSER enum
	ProgramHTTP2HeadersParser ProgramType = C.PROG_HTTP2_HEADERS_PARSER
	// ProgramHTTP2DynamicTableCleaner is the Golang representation of the C.PROG_HTTP2_DYNAMIC_TABLE_CLEANER enum
	ProgramHTTP2DynamicTableCleaner ProgramType = C.PROG_HTTP2_DYNAMIC_TABLE_CLEANER
	// ProgramHTTP2EOSParser is the Golang representation of the C.PROG_HTTP2_EOS_PARSER enum
	ProgramHTTP2EOSParser ProgramType = C.PROG_HTTP2_EOS_PARSER
	// ProgramHTTP2Termination is tail call to process HTTP2 termination.
	ProgramHTTP2Termination ProgramType = C.PROG_HTTP2_TERMINATION
	// ProgramKafka is the Golang representation of the C.PROG_KAFKA enum
	ProgramKafka ProgramType = C.PROG_KAFKA
	// ProgramKafkaResponsePartitionParserV0 is the Golang representation of the C.PROG_KAFKA_RESPONSE_PARTITION_PARSER_v0 enum
	ProgramKafkaResponsePartitionParserV0 ProgramType = C.PROG_KAFKA_RESPONSE_PARTITION_PARSER_V0
	// ProgramKafkaResponsePartitionParserV12 is the Golang representation of the C.PROG_KAFKA_RESPONSE_PARTITION_PARSER_v0 enum
	ProgramKafkaResponsePartitionParserV12 ProgramType = C.PROG_KAFKA_RESPONSE_PARTITION_PARSER_V12
	// ProgramKafkaResponseRecordBatchParserV0 is the Golang representation of the C.PROG_KAFKA_RESPONSE_RECORD_BATCH_PARSER_v0 enum
	ProgramKafkaResponseRecordBatchParserV0 ProgramType = C.PROG_KAFKA_RESPONSE_RECORD_BATCH_PARSER_V0
	// ProgramKafkaResponseRecordBatchParserV12 is the Golang representation of the C.PROG_KAFKA_RESPONSE_RECORD_BATCH_PARSER_v0 enum
	ProgramKafkaResponseRecordBatchParserV12 ProgramType = C.PROG_KAFKA_RESPONSE_RECORD_BATCH_PARSER_V12
	// ProgramKafkaTermination is tail call to process Kafka termination.
	ProgramKafkaTermination ProgramType = C.PROG_KAFKA_TERMINATION
	// ProgramPostgres is the Golang representation of the C.PROG_POSTGRES enum
	ProgramPostgres ProgramType = C.PROG_POSTGRES
	// ProgramPostgresParseMessage is the Golang representation of the C.PROG_POSTGRES_PROCESS_PARSE_MESSAGE enum
	ProgramPostgresParseMessage ProgramType = C.PROG_POSTGRES_PROCESS_PARSE_MESSAGE
	// ProgramPostgresTermination is tail call to process Postgres termination.
	ProgramPostgresTermination ProgramType = C.PROG_POSTGRES_TERMINATION
	// ProgramRedis is the Golang representation of the C.PROG_REDIS enum
	ProgramRedis ProgramType = C.PROG_REDIS
	// ProgramRedisTermination is the Golang representation of the C.PROG_REDIS_TERMINATION enum
	ProgramRedisTermination ProgramType = C.PROG_REDIS_TERMINATION
)

// Application layer of the protocol stack.
func Application(protoNum uint8) ProtocolType {
	return toProtocolType(protoNum, layerApplicationBit)
}

// API layer of the protocol stack.
func API(protoNum uint8) ProtocolType {
	return toProtocolType(protoNum, layerAPIBit)
}

// Encryption layer of the protocol stack.
func Encryption(protoNum uint8) ProtocolType {
	return toProtocolType(protoNum, layerEncryptionBit)
}

func toProtocolType(protoNum uint8, layerBit uint16) ProtocolType {
	if protoNum == 0 {
		return Unknown
	}

	protocol := uint16(protoNum) | uint16(layerBit)
	switch protocol {
	case C.PROTOCOL_UNKNOWN:
		return Unknown
	case C.PROTOCOL_GRPC:
		return GRPC
	case C.PROTOCOL_HTTP:
		return HTTP
	case C.PROTOCOL_HTTP2:
		return HTTP2
	case C.PROTOCOL_KAFKA:
		return Kafka
	case C.PROTOCOL_TLS:
		return TLS
	case C.PROTOCOL_MONGO:
		return Mongo
	case C.PROTOCOL_POSTGRES:
		return Postgres
	case C.PROTOCOL_AMQP:
		return AMQP
	case C.PROTOCOL_REDIS:
		return Redis
	case C.PROTOCOL_MYSQL:
		return MySQL
	default:
		log.Errorf("unknown eBPF protocol type: %x", protocol)
		return Unknown
	}
}
