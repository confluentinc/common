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

import org.apache.kafka.connect.data.Schema;
import org.apache.kafka.connect.data.SchemaAndValue;
import org.apache.kafka.connect.data.SchemaBuilder;

public interface LogRecordBuilder<T> {
  String FIELD_LOGGER = "logger";
  String FIELD_LEVEL = "level";
  String FIELD_TIME = "time";
  String FIELD_MESSAGE = "message";

  String NAMESPACE = "io.confluent.common.logging.";

  static SchemaBuilder baseSchemaBuilder() {
    return SchemaBuilder.struct()
        .name(NAMESPACE + "StructuredLogRecord")
        .field(FIELD_LOGGER, Schema.OPTIONAL_STRING_SCHEMA)
        .field(FIELD_LEVEL, Schema.OPTIONAL_STRING_SCHEMA)
        .field(FIELD_TIME, Schema.OPTIONAL_INT64_SCHEMA);
  }

  LogRecordBuilder<T> withLevel(final String level);

  LogRecordBuilder<T> withTimeMs(final long timeMs);

  LogRecordBuilder<T> withLoggerName(final String loggerName);

  LogRecordBuilder<T> withMessageSchemaAndValue(final SchemaAndValue message);

  LogRecordBuilder<T> withMessageJson(final String message);

  public T build();
}
