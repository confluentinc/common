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
import org.apache.kafka.connect.data.Struct;

public final class LogRecordStructBuilder implements LogRecordBuilder<Struct> {
  private String level = null;
  private Long timeMs = null;
  private String loggerName = null;
  private SchemaAndValue messageWithSchema = null;

  public LogRecordStructBuilder() {
  }

  @Override
  public LogRecordStructBuilder withLevel(final String level) {
    this.level = level;
    return this;
  }

  @Override
  public LogRecordStructBuilder withTimeMs(final long timeMs) {
    this.timeMs = timeMs;
    return this;
  }

  @Override
  public LogRecordStructBuilder withLoggerName(final String loggerName) {
    this.loggerName = loggerName;
    return this;
  }

  @Override
  public LogRecordStructBuilder withMessageSchemaAndValue(final SchemaAndValue messageWithSchema) {
    this.messageWithSchema = messageWithSchema;
    return this;
  }

  @Override
  public LogRecordBuilder<Struct> withMessageJson(String message) {
    throw new RuntimeException("not implemented");
  }

  public Struct build() {
    final Schema logRecordSchema = LogRecordBuilder.baseSchemaBuilder()
        .field(FIELD_MESSAGE, messageWithSchema.schema())
        .build();
    final Struct logRecord = new Struct(logRecordSchema);
    logRecord.put(FIELD_LOGGER, loggerName);
    logRecord.put(FIELD_LEVEL, level);
    logRecord.put(FIELD_TIME, timeMs);
    logRecord.put(FIELD_MESSAGE, messageWithSchema.value());
    return logRecord;
  }
}
