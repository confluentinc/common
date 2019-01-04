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

package io.confluent.common.logging.log4j2;

import io.confluent.common.logging.LogRecordBuilder;
import io.confluent.common.logging.LogRecordStructBuilder;
import org.apache.kafka.connect.data.Struct;
import org.apache.logging.log4j.core.LogEvent;
import org.apache.logging.log4j.core.layout.AbstractLayout;

import io.confluent.common.logging.StructuredLogMessage;

import java.util.function.Function;
import java.util.function.Supplier;

final class StructuredLayout extends AbstractLayout<byte[]> {
  private static final byte[] EMPTY_BYTES = new byte[0];

  private final Function<Struct, byte[]> serializer;
  private final Supplier<LogRecordBuilder<Struct>> logRecordStructBuilderFactory;

  public byte[] toByteArray(final LogEvent event) {
    if (event.getMessage().getParameters().length != 1
        || !(event.getMessage().getParameters()[0] instanceof StructuredLogMessage)) {
      throw new IllegalArgumentException(
          "LogEvent must contain a single parameter of type StructuredLogMessage");
    }
    final StructuredLogMessage schemaAndValue
        = (StructuredLogMessage) event.getMessage().getParameters()[0];
    final Struct logRecord = logRecordStructBuilderFactory.get()
        .withLoggerName(event.getLoggerName())
        .withLevel(event.getLevel().name())
        .withTimeMs(event.getTimeMillis())
        .withMessageSchemaAndValue(schemaAndValue.getMessage())
        .build();
    return serializer.apply(logRecord);
  }

  public String getContentType() {
    return "bytes";
  }

  public byte[] toSerializable(final LogEvent event) {
    return toByteArray(event);
  }

  StructuredLayout(final Function<Struct, byte[]> serializer) {
    this(serializer, LogRecordStructBuilder::new);
  }

  StructuredLayout(
      final Function<Struct, byte[]> serializer,
      final Supplier<LogRecordBuilder<Struct>> logRecordBuilderFactory) {
    super(null, EMPTY_BYTES, EMPTY_BYTES);
    this.serializer = serializer;
    this.logRecordStructBuilderFactory = logRecordBuilderFactory;
  }
}
