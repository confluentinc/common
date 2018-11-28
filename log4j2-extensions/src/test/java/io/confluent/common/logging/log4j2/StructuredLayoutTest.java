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

import io.confluent.common.logging.StructuredLogMessage;
import org.apache.kafka.connect.data.Schema;
import org.apache.kafka.connect.data.SchemaAndValue;
import org.apache.kafka.connect.data.SchemaBuilder;
import org.apache.kafka.connect.data.Struct;
import org.apache.kafka.connect.storage.Converter;
import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.core.LogEvent;
import org.apache.logging.log4j.message.Message;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.ExpectedException;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnit;
import org.mockito.junit.MockitoRule;

import static org.hamcrest.CoreMatchers.equalTo;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

public class StructuredLayoutTest {
  private final String TOPIC = "topic";
  private final Level LOG_LEVEL = Level.INFO;
  private final String LOGGER_NAME = "foo.bar";
  private final long LOG_TIME_MS = 123456L;
  private final byte[] SERIALIZED_MSG = "serialized".getBytes();

  @Mock
  private Converter converter;
  @Mock
  private LogEvent logEvent;
  @Mock
  private Message log4jMessage;
  @Mock
  private StructuredLogMessage logMessage;
  private final Schema schema = SchemaBuilder.struct()
      .field("field", Schema.STRING_SCHEMA)
      .build();
  private final Struct msgStruct = new Struct(schema).put("field", "val");

  private StructuredLayout layout;

  @Rule
  public MockitoRule mockitoRule = MockitoJUnit.rule();
  @Rule
  public ExpectedException exceptionRule = ExpectedException.none();

  @Before
  public void setUp() {
    layout = new StructuredLayout(TOPIC, converter);
    when(logEvent.getMessage()).thenReturn(log4jMessage);
    when(logEvent.getLevel()).thenReturn(LOG_LEVEL);
    when(logEvent.getLoggerName()).thenReturn(LOGGER_NAME);
    when(logEvent.getTimeMillis()).thenReturn(LOG_TIME_MS);
    when(converter.fromConnectData(any(String.class), any(Schema.class), any(Struct.class)))
        .thenReturn(SERIALIZED_MSG);
    when(logMessage.getMessage()).thenReturn(new SchemaAndValue(schema, msgStruct));
  }

  @Test
  public void shouldSerializeMessageCorrectly() {
    // Given:
    when(log4jMessage.getParameters()).thenReturn(new Object[]{logMessage});

    // When:
    final byte[] serialized = layout.toByteArray(logEvent);

    // Then:
    assertThat(serialized, equalTo(SERIALIZED_MSG));
    final ArgumentCaptor<Schema> schemaCaptor = ArgumentCaptor.forClass(Schema.class);
    final ArgumentCaptor<Struct> structCaptor = ArgumentCaptor.forClass(Struct.class);
    verify(converter, times(1)).fromConnectData(
        eq(TOPIC),
        schemaCaptor.capture(),
        structCaptor.capture());
    final Struct struct = structCaptor.getValue();
    assertThat(struct.schema(), equalTo(schemaCaptor.getValue()));
    assertThat(struct.get(LogRecordStructBuilder.FIELD_LEVEL), equalTo(LOG_LEVEL.intLevel()));
    assertThat(struct.get(LogRecordStructBuilder.FIELD_LOGGER), equalTo(LOGGER_NAME));
    assertThat(struct.get(LogRecordStructBuilder.FIELD_TIME), equalTo(LOG_TIME_MS));
    assertThat(struct.get(LogRecordStructBuilder.FIELD_MESSAGE), is(msgStruct));
  }

  @Test
  public void shouldThrowIfInvalidNumberParameters() {
    expectInvalidParameters();
    layout.toByteArray(logEvent);
  }

  @Test
  public void shouldThrowIfInvalidParameterType() {
    expectInvalidParameters(123);
    layout.toByteArray(logEvent);
  }

  private void expectInvalidParameters(Object... params) {
    exceptionRule.expect(IllegalArgumentException.class);
    exceptionRule.expectMessage("LogEvent must contain a single parameter of type StructuredLogMessage");
    when(log4jMessage.getParameters()).thenReturn(params);
  }
}