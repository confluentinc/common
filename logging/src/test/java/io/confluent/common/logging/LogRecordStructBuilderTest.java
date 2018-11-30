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

import org.apache.kafka.connect.data.Field;
import org.apache.kafka.connect.data.Schema;
import org.apache.kafka.connect.data.SchemaAndValue;
import org.apache.kafka.connect.data.SchemaBuilder;
import org.apache.kafka.connect.data.Struct;
import org.junit.Test;

import java.util.stream.Collectors;

import static org.hamcrest.Matchers.contains;
import static org.hamcrest.Matchers.equalTo;
import static org.junit.Assert.assertThat;
import static io.confluent.common.logging.LogRecordBuilder.FIELD_LEVEL;
import static io.confluent.common.logging.LogRecordBuilder.FIELD_LOGGER;
import static io.confluent.common.logging.LogRecordBuilder.FIELD_MESSAGE;
import static io.confluent.common.logging.LogRecordBuilder.FIELD_TIME;

public class LogRecordStructBuilderTest {
  private final LogRecordStructBuilder builder = new LogRecordStructBuilder();

  @Test
  public void shouldBuildStructCorrectly() {
    // Given:
    final Schema msgSchema = SchemaBuilder.struct()
        .field("field1", Schema.STRING_SCHEMA)
        .field("field2", Schema.INT32_SCHEMA)
        .build();
    final Struct msg = new Struct(msgSchema)
        .put("field1", "bizbaz")
        .put("field2", 1129);

    // When:
    final Struct logRecord = builder
        .withLevel("INFO")
        .withLoggerName("foo.bar")
        .withTimeMs(123L)
        .withMessageSchemaAndValue(new SchemaAndValue(msgSchema, msg))
        .build();

    // Then:
    assertThat(
        logRecord.schema().fields().stream().map(Field::name).collect(Collectors.toList()),
        contains(FIELD_LOGGER, FIELD_LEVEL, FIELD_TIME, FIELD_MESSAGE));
    assertThat(logRecord.get(FIELD_LEVEL), equalTo("INFO"));
    assertThat(logRecord.get(FIELD_LOGGER), equalTo("foo.bar"));
    assertThat(logRecord.get(FIELD_TIME), equalTo(123L));
    assertThat(logRecord.get(FIELD_MESSAGE), equalTo(msg));
  }
}