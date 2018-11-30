package io.confluent.common.logging;

import org.apache.kafka.connect.data.SchemaAndValue;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnit;
import org.mockito.junit.MockitoRule;
import org.slf4j.Logger;

import java.util.Map;
import java.util.function.Function;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

public class StructuredLoggerFactoryTest {
  private final static String PREFIX = "prefix.";

  @Mock
  private Function<String, Logger> innerFactory;
  @Mock
  private Logger innerLogger;

  private StructuredLoggerFactory loggerFactory;

  @Rule
  public MockitoRule mockitoRule = MockitoJUnit.rule();

  @Before
  public void setup() {
    loggerFactory = new StructuredLoggerFactory(PREFIX, innerFactory);
  }

  @Test
  public void shouldCreateLoggerCorrectly() {
    // Given:
    when(innerFactory.apply("prefix.foo")).thenReturn(innerLogger);

    // When:
    final StructuredLogger logger = loggerFactory.getLogger("foo");
    logger.info(mock(SchemaAndValue.class));

    // Then:
    verify(innerFactory, times(1)).apply("prefix.foo");
    verify(innerLogger).info(any(String.class), any(SerializableSchemaAndValue.class));
  }

  @Test
  public void shouldAddPrefixToClassName() {
    // Given:
    when(innerFactory.apply("prefix.java.util.Map")).thenReturn(innerLogger);

    // When:
    loggerFactory.getLogger(Map.class);

    // Then:
    verify(innerFactory, times(1)).apply("prefix.java.util.Map");
  }
}