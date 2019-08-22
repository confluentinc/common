/**
 * Copyright 2017 Confluent Inc.
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
 **/

package io.confluent.common.config;

import org.junit.Test;

import java.util.Properties;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertNull;

public class ConfigUtilsTest {
  @Test
  public void testTranslateDeprecated() throws Exception {
    Properties props = new Properties();
    props.setProperty("foo.bar", "baz");
    props.setProperty("foo.bar.deprecated", "quux");
    props.setProperty("chicken", "1");
    props.setProperty("rooster", "2");
    props.setProperty("hen", "3");
    props.setProperty("heifer", "moo");
    props.setProperty("blah", "blah");
    props.put("unexpected.non.string.object", new Integer(42));
    Properties newProps = ConfigUtils.translateDeprecated(props, new String[][]{
        {"foo.bar", "foo.bar.deprecated"},
        {"chicken", "rooster", "hen"},
        {"cow", "beef", "heifer", "steer"}
    });
    assertEquals("baz", newProps.getProperty("foo.bar"));
    assertEquals(null, newProps.getProperty("foo.bar.deprecated"));
    assertEquals("1", newProps.getProperty("chicken"));
    assertEquals(null, newProps.getProperty("rooster"));
    assertEquals(null, newProps.getProperty("hen"));
    assertEquals("moo", newProps.getProperty("cow"));
    assertEquals(null, newProps.getProperty("beef"));
    assertEquals(null, newProps.getProperty("heifer"));
    assertEquals(null, newProps.getProperty("steer"));
    assertEquals(null, props.getProperty("cow"));
    assertEquals("blah", props.getProperty("blah"));
    assertEquals("blah", newProps.getProperty("blah"));

    // The java.util.Properties class was intended to store only String values.
    // However, because of a design mistake, it can actually store arbitrary Objects.
    // They are not returned when Properties#getProperty is invoked, but they are
    // returned when Properties#get is invoked.
    // Here, we test that ConfigUtils passes through these objects unchanged.
    assertEquals(null, newProps.getProperty("unexpected.non.string.object"));
    assertEquals(new Integer(42), newProps.get("unexpected.non.string.object"));
    assertEquals(null, props.getProperty("unexpected.non.string.object"));
    assertEquals(new Integer(42), props.get("unexpected.non.string.object"));

  }

  @Test
  public void testAllowsNewKey() throws Exception {
    Properties props = new Properties();
    props.setProperty("foo.bar", "baz");
    Properties newProps = ConfigUtils.translateDeprecated(props, new String[][]{
        {"foo.bar", "foo.bar.deprecated"},
        {"chicken", "rooster", "hen"},
        {"cow", "beef", "heifer", "steer"}
    });
    assertNotNull(newProps);
    assertEquals("baz", newProps.getProperty("foo.bar"));
    assertNull(newProps.getProperty("foo.bar.deprecated"));
  }

  @Test
  public void testDuplicateSynonyms() throws Exception {
    Properties props = new Properties();
    props.setProperty("foo.bar", "baz");
    props.setProperty("foo.bar.deprecated", "derp");
    Properties newProps = ConfigUtils.translateDeprecated(props, new String[][]{
        {"foo.bar", "foo.bar.deprecated"},
        {"chicken", "foo.bar.deprecated"}
    });
    assertNotNull(newProps);
    assertEquals("baz", newProps.getProperty("foo.bar"));
    assertEquals("derp", newProps.getProperty("chicken"));
    assertNull(newProps.getProperty("foo.bar.deprecated"));
  }

  @Test
  public void testMultipleDeprecations() throws Exception {
    Properties props = new Properties();
    props.setProperty("foo.bar.deprecated", "derp");
    props.setProperty("foo.bar.even.more.deprecated", "very old configuration");
    Properties newProps = ConfigUtils.translateDeprecated(props, new String[][]{
        {"foo.bar", "foo.bar.deprecated", "foo.bar.even.more.deprecated"}
    });
    assertNotNull(newProps);
    assertEquals("derp", newProps.getProperty("foo.bar"));
    assertNull(newProps.getProperty("foo.bar.deprecated"));
    assertNull(newProps.getProperty("foo.bar.even.more.deprecated"));
  }
}
