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

import java.util.Collection;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.function.Function;

import com.google.common.collect.ImmutableList;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public final class StructuredLoggerFactory {
  private final String prefix;
  private final Map<String, StructuredLoggerImpl> loggers = new ConcurrentHashMap<>();
  private final Function<String, Logger> inner;

  public StructuredLoggerFactory(final String prefix) {
    this(prefix, LoggerFactory::getLogger);
  }

  StructuredLoggerFactory(final String prefix, final Function<String, Logger> inner) {
    this.prefix = prefix;
    this.inner = inner;
  }

  public StructuredLogger getLogger(final Class<?> clazz) {
    return getLogger(clazz.getName());
  }

  public StructuredLogger getLogger(final String name) {
    final String loggerName = prefix + name;
    return loggers.computeIfAbsent(
        loggerName,
        key -> new StructuredLoggerImpl(inner.apply(key))
    );
  }

  public Collection<String> getLoggers() {
    return ImmutableList.copyOf(loggers.keySet());
  }
}
