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

import java.util.function.Supplier;
import org.apache.kafka.connect.data.SchemaAndValue;
import org.slf4j.Logger;

final class StructuredLoggerImpl implements StructuredLogger {
  private static final String LOG_MSG = "{}";

  private final Logger inner;

  StructuredLoggerImpl(final Logger inner) {
    this.inner = inner;
  }

  public String getName() {
    return inner.getName();
  }

  public void error(final Supplier<SchemaAndValue> msgSupplier) {
    if (!inner.isErrorEnabled()) {
      return;
    }
    error(msgSupplier.get());
  }

  public void error(final SchemaAndValue msg) {
    inner.error(LOG_MSG, new SerializableSchemaAndValue(msg));
  }

  public void info(final Supplier<SchemaAndValue> msgSupplier) {
    if (!inner.isInfoEnabled()) {
      return;
    }
    info(msgSupplier.get());
  }

  public void info(final SchemaAndValue msg) {
    inner.info(LOG_MSG, new SerializableSchemaAndValue(msg));
  }

  public void debug(final Supplier<SchemaAndValue> msgSupplier) {
    if (!inner.isDebugEnabled()) {
      return;
    }
    debug(msgSupplier.get());
  }

  public void debug(final SchemaAndValue msg) {
    inner.debug(LOG_MSG, new SerializableSchemaAndValue(msg));
  }
}
