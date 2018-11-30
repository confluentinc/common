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
import org.junit.Test;

import java.io.IOException;
import java.util.Map;

import static org.hamcrest.CoreMatchers.equalTo;
import static org.hamcrest.CoreMatchers.instanceOf;
import static org.junit.Assert.assertThat;

public class LogRecordStringBuilderTest {
  private final ObjectMapper objectMapper = new ObjectMapper();

  private final LogRecordStringBuilder builder = new LogRecordStringBuilder();

  @Test
  public void shouldSerializeRecordCorrectly() throws IOException {
    final String record = builder
        .withLevel("INFO")
        .withLoggerName("foo.bar")
        .withTimeMs(123L)
        .withMessageJson("{\"field1\": 456, \"field2\": \"bizbaz\"}")
        .build();

    final Object deserialized = objectMapper.readValue(record, Object.class);
    assertThat(deserialized, instanceOf(Map.class));
    final Map asMap = (Map) deserialized;
    final Map expected = ImmutableMap.of(
        LogRecordBuilder.FIELD_LEVEL, "INFO",
        LogRecordBuilder.FIELD_LOGGER, "foo.bar",
        LogRecordBuilder.FIELD_TIME, 123,
        LogRecordBuilder.FIELD_MESSAGE,
        ImmutableMap.of(
            "field1", 456,
            "field2", "bizbaz"
        )
    );
    assertThat(asMap, equalTo(expected));
  }
}