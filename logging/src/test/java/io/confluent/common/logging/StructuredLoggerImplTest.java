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
import org.apache.kafka.connect.data.Schema;
import org.apache.kafka.connect.data.SchemaAndValue;
import org.apache.kafka.connect.data.SchemaBuilder;
import org.apache.kafka.connect.data.Struct;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.mockito.ArgumentCaptor;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnit;
import org.mockito.junit.MockitoRule;
import org.slf4j.Logger;

import java.io.IOException;
import java.util.Map;
import java.util.function.Supplier;

import static org.hamcrest.CoreMatchers.equalTo;
import static org.hamcrest.CoreMatchers.instanceOf;
import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.Matchers.contains;
import static org.junit.Assert.assertThat;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoMoreInteractions;
import static org.mockito.Mockito.when;

public class StructuredLoggerImplTest {
  private final static String LOG_MSG = "{}";

  @Mock
  private Logger innerLogger;
  @Mock
  private SchemaAndValue msg;
  @Mock
  private Supplier<SchemaAndValue> msgSupplier;

  private StructuredLoggerImpl logger;

  @Rule
  public MockitoRule rule = MockitoJUnit.rule();

  @Before
  public void setup() {
    when(msgSupplier.get()).thenReturn(msg);
    logger = new StructuredLoggerImpl(innerLogger);
  }

  @Test
  public void shouldLogErrorMessage() {
    // When:
    logger.error(msg);

    // Then:
    final ArgumentCaptor<SerializableSchemaAndValue> captor = createMsgCaptor();
    verify(innerLogger, times(1)).error(eq(LOG_MSG), captor.capture());
    assertThat(captor.getValue().getMessage(), is(msg));
  }

  @Test
  public void shouldLogErrorMessageGivenSupplier() {
    // Given:
    when(innerLogger.isErrorEnabled()).thenReturn(true);

    // When:
    logger.error(msgSupplier);

    // Then:
    final ArgumentCaptor<SerializableSchemaAndValue> captor = createMsgCaptor();
    verify(msgSupplier, times(1)).get();
    verify(innerLogger, times(1)).error(eq(LOG_MSG), captor.capture());
    assertThat(captor.getValue().getMessage(), is(msg));
  }

  @Test
  public void shouldNotGetMessageIfLoggerDisabledForError() {
    // Given:
    when(innerLogger.isErrorEnabled()).thenReturn(false);

    // When:
    logger.error(msgSupplier);

    // Then:
    verify(innerLogger).isErrorEnabled();
    verifyNoMoreInteractions(innerLogger, msgSupplier);
  }

  @Test
  public void shouldLogInfoMessage() {
    // When:
    logger.info(msg);

    // Then:
    final ArgumentCaptor<SerializableSchemaAndValue> captor = createMsgCaptor();
    verify(innerLogger, times(1)).info(eq(LOG_MSG), captor.capture());
    assertThat(captor.getValue().getMessage(), is(msg));
  }

  @Test
  public void shouldLogInfoMessageGivenSupplier() {
    // Given:
    when(innerLogger.isInfoEnabled()).thenReturn(true);

    // When:
    logger.info(msgSupplier);

    // Then:
    final ArgumentCaptor<SerializableSchemaAndValue> captor = createMsgCaptor();
    verify(msgSupplier, times(1)).get();
    verify(innerLogger, times(1)).info(eq(LOG_MSG), captor.capture());
    assertThat(captor.getValue().getMessage(), is(msg));
  }

  @Test
  public void shouldNotGetMessageIfLoggerDisabledForInfo() {
    // Given:
    when(innerLogger.isInfoEnabled()).thenReturn(false);

    // When:
    logger.info(msgSupplier);

    // Then:
    verify(innerLogger).isInfoEnabled();
    verifyNoMoreInteractions(innerLogger, msgSupplier);
  }

  @Test
  public void shouldLogDebugMessage() {
    // When:
    logger.debug(msg);

    // Then:
    final ArgumentCaptor<SerializableSchemaAndValue> captor = createMsgCaptor();
    verify(innerLogger, times(1)).debug(eq(LOG_MSG), captor.capture());
    assertThat(captor.getValue().getMessage(), is(msg));
  }

  @Test
  public void shouldLogDebugMessageGivenSupplier() {
    // Given:
    when(innerLogger.isDebugEnabled()).thenReturn(true);

    // When:
    logger.debug(msgSupplier);

    // Then:
    final ArgumentCaptor<SerializableSchemaAndValue> captor = createMsgCaptor();
    verify(msgSupplier, times(1)).get();
    verify(innerLogger, times(1)).debug(eq(LOG_MSG), captor.capture());
    assertThat(captor.getValue().getMessage(), is(msg));
  }

  @Test
  public void shouldNotGetMessageIfLoggerDisabledForDebug() {
    // Given:
    when(innerLogger.isDebugEnabled()).thenReturn(false);

    // When:
    logger.debug(msgSupplier);

    // Then:
    verify(innerLogger).isDebugEnabled();
    verifyNoMoreInteractions(innerLogger, msgSupplier);
  }

  @Test
  public void shouldSerializeMessageToJsonString() throws IOException {
    // Given:
    final Schema msgSchema = SchemaBuilder.struct()
        .field("field1", Schema.STRING_SCHEMA)
        .field("field2", Schema.INT32_SCHEMA)
        .build();
    final Struct msgStruct = new Struct(msgSchema);
    msgStruct.put("field1", "foobar");
    msgStruct.put("field2", 123);
    final SchemaAndValue schemaAndValue = new SchemaAndValue(msgSchema, msgStruct);

    // When:
    logger.info(schemaAndValue);

    // Then:
    final ArgumentCaptor<SerializableSchemaAndValue> captor = createMsgCaptor();
    verify(innerLogger).info(any(), captor.capture());
    final String asString = captor.getValue().toString();
    final Object deserialized = new ObjectMapper().readValue(asString, Object.class);
    assertThat(deserialized, instanceOf(Map.class));
    assertThat(
        deserialized,
        equalTo(
            ImmutableMap.of(
                "field1", "foobar",
                "field2", 123)));
  }

  private ArgumentCaptor<SerializableSchemaAndValue> createMsgCaptor() {
    return ArgumentCaptor.forClass(SerializableSchemaAndValue.class);
  }
}