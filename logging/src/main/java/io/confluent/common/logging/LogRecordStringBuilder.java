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

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonRawValue;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.kafka.connect.data.SchemaAndValue;

public final class LogRecordStringBuilder implements LogRecordBuilder<String> {
  private static final ObjectMapper mapper = new ObjectMapper();

  private String level = null;
  private Long timeMs = null;
  private String loggerName = null;
  private String message = null;

  private static class LogRecord {
    @JsonProperty(FIELD_LEVEL) String level;
    @JsonProperty(FIELD_LOGGER) String logger;
    @JsonProperty(FIELD_TIME) Long timeMs;

    @JsonProperty(FIELD_MESSAGE)
    @JsonRawValue
    String message;

    public LogRecord(
        final String level,
        final String logger,
        final Long timeMs,
        final String message) {
      this.level = level;
      this.logger = logger;
      this.timeMs = timeMs;
      this.message = message;
    }
  }

  @Override
  public LogRecordStringBuilder withLevel(final String level) {
    this.level = level;
    return this;
  }

  @Override
  public LogRecordStringBuilder withTimeMs(final long timeMs) {
    this.timeMs = timeMs;
    return this;
  }

  @Override
  public LogRecordStringBuilder withLoggerName(final String loggerName) {
    this.loggerName = loggerName;
    return this;
  }

  @Override
  public LogRecordStringBuilder withMessageSchemaAndValue(final SchemaAndValue messageWithSchema) {
    throw new RuntimeException("not implemented");
  }

  @Override
  public LogRecordBuilder<String> withMessageJson(final String message) {
    this.message = message;
    return this;
  }

  public String build() {
    try {
      return mapper.writeValueAsString(
          new LogRecord(level, loggerName, timeMs, message));
    } catch (final JsonProcessingException e) {
      throw new RuntimeException(e);
    }
  }
}
