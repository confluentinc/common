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
import io.confluent.common.logging.StructuredLogMessage;
import org.apache.kafka.connect.data.Schema;
import org.apache.kafka.connect.data.SchemaAndValue;
import org.apache.kafka.connect.data.Struct;
import org.apache.kafka.connect.storage.Converter;
import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.core.LogEvent;
import org.apache.logging.log4j.message.Message;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.ExpectedException;
import org.mockito.InOrder;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.MockitoJUnit;
import org.mockito.junit.MockitoRule;

import java.util.function.Consumer;
import java.util.function.Function;

import static org.hamcrest.CoreMatchers.equalTo;
import static org.junit.Assert.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

public class StructuredLayoutTest {
  private static final String TOPIC = "topic";
  private static final Level LOG_LEVEL = Level.INFO;
  private static final String LOGGER_NAME = "foo.bar";
  private static final long LOG_TIME_MS = 123456L;
  private static final byte[] SERIALIZED_MSG = "serialized".getBytes();

  @Mock
  private Function<Struct, byte[]> converter;
  @Mock
  private LogEvent logEvent;
  @Mock
  private Message log4jMessage;
  @Mock
  private LogRecordBuilder<Struct> builder;
  @Mock
  private StructuredLogMessage logMessage;
  @Mock
  private Schema schema;
  @Mock
  private Struct struct;
  private SchemaAndValue schemaAndValue;

  private StructuredLayout layout;

  @Rule
  public MockitoRule mockitoRule = MockitoJUnit.rule();
  @Rule
  public ExpectedException exceptionRule = ExpectedException.none();

  @Before
  public void setup() {
    layout = new StructuredLayout(converter, () -> builder);
    when(logEvent.getMessage()).thenReturn(log4jMessage);
    when(logEvent.getLevel()).thenReturn(LOG_LEVEL);
    when(logEvent.getLoggerName()).thenReturn(LOGGER_NAME);
    when(logEvent.getTimeMillis()).thenReturn(LOG_TIME_MS);
    when(converter.apply(any(Struct.class)))
        .thenReturn(SERIALIZED_MSG);
    schemaAndValue = new SchemaAndValue(schema, struct);
    when(logMessage.getMessage()).thenReturn(schemaAndValue);
    when(builder.withLevel(anyString())).thenReturn(builder);
    when(builder.withLoggerName(anyString())).thenReturn(builder);
    when(builder.withMessageSchemaAndValue(any(SchemaAndValue.class)))
        .thenReturn(builder);
    when(builder.withTimeMs(anyLong())).thenReturn(builder);
    when(struct.schema()).thenReturn(schema);
  }

  private void verifyBeforeBuild(
      final LogRecordBuilder<Struct> builder,
      final Consumer<InOrder> action) {
    final InOrder inOrder = Mockito.inOrder(builder);
    action.accept(inOrder);
    inOrder.verify(builder).build();
  }

  @Test
  public void shouldSerializeMessageCorrectly() {
    // Given:
    when(log4jMessage.getParameters()).thenReturn(new Object[]{logMessage});
    final Schema logSchema = mock(Schema.class);
    final Struct logRecord = mock(Struct.class);
    when(logRecord.schema()).thenReturn(logSchema);
    when(builder.build()).thenReturn(logRecord);

    // When:
    final byte[] serialized = layout.toByteArray(logEvent);

    // Then:
    verifyBeforeBuild(builder, io -> io.verify(builder).withLoggerName(LOGGER_NAME));
    verifyBeforeBuild(builder, io -> io.verify(builder).withTimeMs(LOG_TIME_MS));
    verifyBeforeBuild(builder, io -> io.verify(builder).withLevel(LOG_LEVEL.name()));
    verifyBeforeBuild(
        builder,
        io -> io.verify(builder).withMessageSchemaAndValue(schemaAndValue));
    verify(converter, times(1))
        .apply(logRecord);
    assertThat(serialized, equalTo(SERIALIZED_MSG));
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