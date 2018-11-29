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
import org.apache.log4j.Category;
import org.apache.log4j.Level;
import org.apache.log4j.spi.LoggingEvent;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.mockito.InOrder;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.MockitoJUnit;
import org.mockito.junit.MockitoRule;

import java.util.function.Consumer;

import static org.hamcrest.CoreMatchers.equalTo;
import static org.junit.Assert.assertThat;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.when;

public class StructuredJsonLayoutTest {
  private static final Level LOG_LEVEL = Level.INFO;
  private static final String LOGGER_NAME = "foo.bar";
  private static final long LOG_TIME_MS = 123456L;
  private static final String STRUCTURED_MSG = "msg";
  private static final String SERIALIZED_MSG = "serialized";

  @Mock
  private LogRecordBuilder<String> builder;
  private LoggingEvent loggingEvent;
  private StructuredJsonLayout layout;

  @Rule
  public MockitoRule mockitoRule = MockitoJUnit.rule();

  @Before
  public void setUp() {
    when(builder.withTimeMs(anyLong())).thenReturn(builder);
    when(builder.withLoggerName(anyString())).thenReturn(builder);
    when(builder.withLevel(anyString())).thenReturn(builder);
    when(builder.withMessageJson(anyString())).thenReturn(builder);
    when(builder.build()).thenReturn(SERIALIZED_MSG);
    loggingEvent = new LoggingEvent(
        "fcqn",
        new FakeCategory(LOGGER_NAME),
        LOG_TIME_MS,
        LOG_LEVEL,
        STRUCTURED_MSG,
        null
    );
    layout = new StructuredJsonLayout(() -> builder);
  }

  private void verifyBeforeBuild(
      final LogRecordBuilder<String> builder,
      final Consumer<InOrder> action) {
    final InOrder inOrder = Mockito.inOrder(builder);
    action.accept(inOrder);
    inOrder.verify(builder).build();
  }

  @Test
  public void shouldFormatLogRecordCorrectly() {
    // When:
    final String msg = layout.format(loggingEvent);

    // Then:
    verifyBeforeBuild(builder, io -> io.verify(builder).withLevel(LOG_LEVEL.toString()));
    verifyBeforeBuild(builder, io -> io.verify(builder).withTimeMs(LOG_TIME_MS));
    verifyBeforeBuild(builder, io -> io.verify(builder).withLoggerName(LOGGER_NAME));
    verifyBeforeBuild(builder, io -> io.verify(builder).withMessageJson(STRUCTURED_MSG));
    assertThat(msg, equalTo(SERIALIZED_MSG));
  }

  private static class FakeCategory extends Category {
    private FakeCategory(final String name) {
      super(name);
    }
  }
}