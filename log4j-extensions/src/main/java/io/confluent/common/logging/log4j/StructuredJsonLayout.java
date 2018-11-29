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

package io.confluent.common.logging.log4j;

import io.confluent.common.logging.LogRecordBuilder;
import io.confluent.common.logging.LogRecordStringBuilder;
import org.apache.log4j.Layout;
import org.apache.log4j.spi.LoggingEvent;

import java.util.function.Supplier;

public class StructuredJsonLayout extends Layout {
  final Supplier<LogRecordBuilder<String>> logRecordBuilderFactory;

  @Override
  public String format(LoggingEvent loggingEvent) {
    final LogRecordBuilder<String> recordBuilder = logRecordBuilderFactory.get();
    return recordBuilder.withLevel(loggingEvent.getLevel().toString())
        .withLoggerName(loggingEvent.getLoggerName())
        .withTimeMs(loggingEvent.getTimeStamp())
        .withMessageJson(loggingEvent.getRenderedMessage())
        .build();
  }

  @Override
  public boolean ignoresThrowable() {
    return true;
  }

  @Override
  public void activateOptions() {
  }

  public StructuredJsonLayout() {
    this(LogRecordStringBuilder::new);
  }

  StructuredJsonLayout(
      final Supplier<LogRecordBuilder<String>> logRecordBuilderFactory) {
    this.logRecordBuilderFactory = logRecordBuilderFactory;
  }
}
