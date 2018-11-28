package io.confluent.common.logging;

import org.apache.kafka.connect.data.SchemaAndValue;

public interface StructuredLogMessage {
  SchemaAndValue getMessage();
}
