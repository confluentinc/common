/*
 * Copyright 2018 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.confluent.common.logging;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.common.collect.ImmutableMap;
import org.apache.kafka.connect.data.SchemaAndValue;
import org.apache.kafka.connect.json.JsonConverter;
import org.apache.kafka.connect.json.JsonConverterConfig;

import java.io.IOException;

// Wrapper around SchemaAndValue that implements toString by serializing to json
// so that string-based loggers (e.g. log4j1) can log useful messages for structured
// events.
final class SerializableSchemaAndValue implements StructuredLogMessage {
  private static final JsonConverter converter = new JsonConverter();
  private static final ObjectMapper objectMapper = new ObjectMapper();

  static {
    converter.configure(
        ImmutableMap.of(
            JsonConverterConfig.SCHEMAS_ENABLE_CONFIG, false
        ),
        false
    );
  }

  private final SchemaAndValue schemaAndValue;
  private volatile String asString = null;

  SerializableSchemaAndValue(final SchemaAndValue schemaAndValue) {
    this.schemaAndValue = schemaAndValue;
  }

  public SchemaAndValue getMessage() {
    return schemaAndValue;
  }

  private String serializeToString() {
    final byte[] bytes =
        converter.fromConnectData("", schemaAndValue.schema(), schemaAndValue.value());
    try {
      return objectMapper.writeValueAsString(
          objectMapper.readValue(bytes, Object.class));
    } catch (final IOException e) {
      throw new RuntimeException(e);
    }
  }

  @Override
  public String toString() {
    if (asString == null) {
      asString = serializeToString();
    }
    return asString;
  }
}
