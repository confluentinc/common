/*
 * Copyright 2015 Confluent Inc.
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

/*
 * Original license:
 *
 * Licensed to the Apache Software Foundation (ASF) under one or more contributor license agreements. See the NOTICE
 * file distributed with this work for additional information regarding copyright ownership. The ASF licenses this file
 * to You under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the
 * License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */
package io.confluent.common.config;

import org.junit.Assert;
import org.junit.Test;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;

import io.confluent.common.config.types.Password;

import javax.xml.bind.DatatypeConverter;

import static io.confluent.common.config.ConfigDef.Type;
import static io.confluent.common.config.ConfigDef.Type.BOOLEAN;
import static io.confluent.common.config.ConfigDef.Type.CLASS;
import static io.confluent.common.config.ConfigDef.Type.DOUBLE;
import static io.confluent.common.config.ConfigDef.Type.INT;
import static io.confluent.common.config.ConfigDef.Type.LIST;
import static io.confluent.common.config.ConfigDef.Type.LONG;
import static io.confluent.common.config.ConfigDef.Type.STRING;
import static java.util.Arrays.asList;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertNotEquals;
import static org.junit.Assert.fail;

public class ConfigDefTest {

  @Test
  public void testCopyConstructor() {
    ConfigDef original = new ConfigDef().define("a", INT, ConfigDef.Importance.LOW, "docs")
        .define("b", STRING, ConfigDef.Importance.HIGH, "docs");

    ConfigDef copy = new ConfigDef(original);

    Properties props = new Properties();
    props.put("a", "5");
    props.put("b", "foo");
    props.put("c", "true");

    // Ensure the copied ConfigDef contains all the definitions of the original
    Map<String, Object> originalParsedProps = original.parse(props);
    Map<String, Object> copyParsedProps = copy.parse(props);
    assertEquals(originalParsedProps, copyParsedProps);
    assertFalse(copyParsedProps.containsKey("c"));

    copy.define("c", BOOLEAN, ConfigDef.Importance.MEDIUM, "docs");

    // Ensure that mutating the copied ConfigDef doesn't also mutate the original
    originalParsedProps = original.parse(props);
    copyParsedProps = copy.parse(props);
    assertNotEquals(originalParsedProps, copyParsedProps);
    assertEquals(true, copyParsedProps.get("c"));
  }

  @Test
  public void testBasicTypes() {
    ConfigDef def = new ConfigDef().define("a", INT, 5, ConfigDef.Range
        .between(0, 14), ConfigDef.Importance.HIGH, "docs")
        .define("b", LONG, ConfigDef.Importance.HIGH, "docs")
        .define("c", STRING, "hello", ConfigDef.Importance.HIGH, "docs")
        .define("d", LIST, ConfigDef.Importance.HIGH, "docs")
        .define("e", DOUBLE, ConfigDef.Importance.HIGH, "docs")
        .define("f", CLASS, ConfigDef.Importance.HIGH, "docs")
        .define("g", BOOLEAN, ConfigDef.Importance.HIGH, "docs")
        .define("h", Type.BOOLEAN, ConfigDef.Importance.HIGH, "docs")
        .define("i", Type.BOOLEAN, ConfigDef.Importance.HIGH, "docs")
        .define("j", Type.PASSWORD, ConfigDef.Importance.HIGH, "docs")
        .define("k", Type.MAP, ConfigDef.Importance.HIGH, "docs")
        .define("l", Type.MAP, ConfigDef.Importance.HIGH, "docs");

    Properties props = new Properties();
    props.put("a", "1   ");
    props.put("b", 2);
    props.put("d", " a , b, c");
    props.put("e", 42.5d);
    props.put("f", String.class.getName());
    props.put("g", "true");
    props.put("h", "FalSE");
    props.put("i", "TRUE");
    props.put("j", "password");
    props.put("k", "k1:v1,k2:v2");
    props.put("l", "k1:v1");

    Map<String, Object> vals = def.parse(props);
    assertEquals(1, vals.get("a"));
    assertEquals(2L, vals.get("b"));
    assertEquals("hello", vals.get("c"));
    assertEquals(asList("a", "b", "c"), vals.get("d"));
    assertEquals(42.5d, vals.get("e"));
    assertEquals(String.class, vals.get("f"));
    assertEquals(true, vals.get("g"));
    assertEquals(false, vals.get("h"));
    assertEquals(true, vals.get("i"));
    assertEquals(new Password("password"), vals.get("j"));
    Map expectedMap = new HashMap(){
      {
        put("k1","v1");
        put("k2","v2");
      }
    };
    assertEquals(expectedMap, vals.get("k"));
    expectedMap = new HashMap(){
      {
        put("k1","v1");
      }
    };
    assertEquals(expectedMap, vals.get("l"));
  }

  @Test
  public void testOverride() {
    ConfigDef def = new ConfigDef()
        .define("a", INT, 0, ConfigDef.Range.between(0, 0),
                        ConfigDef.Importance.HIGH, "doc")
        .define("b", STRING, "original", ConfigDef.Importance.HIGH, "doc")
        .define("c", STRING, ConfigDef.ValidString.in(Arrays.asList("", "foo")),
                ConfigDef.Importance.HIGH, "doc")
        .define("d", STRING, ConfigDef.Importance.HIGH, "doc")
        .defineOverride("a", INT, 4, ConfigDef.Range.between(0, 14),
                    ConfigDef.Importance.HIGH, "doc")
        .defineOverride("b", STRING, "override", ConfigDef.Importance.HIGH, "doc")
        .defineOverride("c", STRING, ConfigDef.ValidString.in(Arrays.asList("", "bar")),
                        ConfigDef.Importance.HIGH, "doc")
        .defineOverride("d", STRING, ConfigDef.Importance.HIGH, "doc");

    Properties props = new Properties();
    props.put("c", "bar");
    props.put("d", "foo");

    Map<String,Object> vals = def.parse(props);
    assertEquals(4, vals.get("a"));
    assertEquals("override", vals.get("b"));
    assertEquals("bar", vals.get("c"));
    assertEquals("foo", vals.get("d"));
  }

  @Test(expected = ConfigException.class)
  public void testInvalidDefault() {
    new ConfigDef().define("a", Type.INT, "hello", ConfigDef.Importance.HIGH, "docs");
  }

  @Test(expected = ConfigException.class)
  public void testNullDefault() {
    new ConfigDef().define("a", Type.INT, null, null, null, "docs");
  }

  @Test(expected = ConfigException.class)
  public void testMissingRequired() {
    new ConfigDef().define("a", Type.INT, ConfigDef.Importance.HIGH, "docs")
        .parse(new HashMap<String, Object>());
  }

  @Test
  public void testParsingEmptyDefaultValueForStringFieldShouldSucceed() {
    new ConfigDef().define("a", Type.STRING, "", ConfigDef.Importance.HIGH, "docs")
            .parse(new HashMap<String, Object>());
  }

  @Test(expected = ConfigException.class)
  public void testDefinedTwice() {
    new ConfigDef().define("a", Type.STRING, ConfigDef.Importance.HIGH, "docs")
        .define("a", Type.INT, ConfigDef.Importance.HIGH, "docs");
  }

  @Test(expected = ConfigException.class)
  public void testOverrideNotPreviouslyDefined() {
    new ConfigDef().defineOverride("a", Type.STRING, ConfigDef.Importance.HIGH, "docs");
  }

  @Test
  public void testBadInputs() {
    testBadInputs(Type.INT, "hello", null, "42.5", 42.5, Long.MAX_VALUE,
                  Long.toString(Long.MAX_VALUE), new Object());
    testBadInputs(Type.LONG, "hello", null, "42.5", Long.toString(Long.MAX_VALUE) + "00",
                  new Object());
    testBadInputs(Type.DOUBLE, "hello", null, new Object());
    testBadInputs(Type.STRING, new Object());
    testBadInputs(Type.LIST, 53, new Object());
    testBadInputs(Type.BOOLEAN, "hello", "truee", "fals");
    testBadInputs(Type.MAP, "test", new Object());
  }

  private void testBadInputs(Type type, Object... values) {
    for (Object value : values) {
      Map<String, Object> m = new HashMap<String, Object>();
      m.put("name", value);
      ConfigDef def = new ConfigDef().define("name", type, ConfigDef.Importance.HIGH, "docs");
      try {
        def.parse(m);
        fail("Expected a config exception on bad input for value " + value);
      } catch (ConfigException e) {
        // this is good
      }
    }
  }

  @Test(expected = ConfigException.class)
  public void testInvalidDefaultRange() {
    ConfigDef
        def =
        new ConfigDef()
            .define("name", Type.INT, -1, ConfigDef.Range.between(0, 10), ConfigDef.Importance.HIGH,
                    "docs");
  }

  @Test(expected = ConfigException.class)
  public void testInvalidDefaultString() {
    ConfigDef def = new ConfigDef().define("name", Type.STRING, "bad", ConfigDef.ValidString
        .in(Arrays.asList("valid", "values")), ConfigDef.Importance.HIGH, "docs");
  }

  @Test
  public void testValidators() {
    testValidators(Type.INT, ConfigDef.Range.between(0, 10), 5, new Object[]{1, 5, 9},
                   new Object[]{-1, 11});
    testValidators(Type.STRING, ConfigDef.ValidString
                       .in(Arrays.asList("good", "values", "default")), "default",
                   new Object[]{"good", "values", "default"}, new Object[]{"bad", "inputs"});
  }

  private void testValidators(Type type, ConfigDef.Validator validator, Object defaultVal,
                              Object[] okValues, Object[] badValues) {
    ConfigDef
        def =
        new ConfigDef()
            .define("name", type, defaultVal, validator, ConfigDef.Importance.HIGH, "docs");

    for (Object value : okValues) {
      Map<String, Object> m = new HashMap<String, Object>();
      m.put("name", value);
      def.parse(m);
    }

    for (Object value : badValues) {
      Map<String, Object> m = new HashMap<String, Object>();
      m.put("name", value);
      try {
        def.parse(m);
        fail("Expected a config exception due to invalid value " + value);
      } catch (ConfigException e) {
        // this is good
      }
    }
  }

  @Test
  public void attemptFindStringClass() {
    ConfigDef def = new ConfigDef().define("f", CLASS, ConfigDef.Importance.HIGH, "docs");
    Properties props = new Properties();
    Thread.currentThread().setContextClassLoader(null);
    props.put("f", "java.lang.String");
    Map<String, Object> vals = def.parse(props);
    Assert.assertNotNull(vals.get("f"));
  }

  @Test
  public void attemptFindFoo() {
    ConfigDef def = new ConfigDef().define("f", CLASS, ConfigDef.Importance.HIGH, "docs");
    Properties props = new Properties();
    Thread.currentThread().setContextClassLoader(new FooLoader(ClassLoader.getSystemClassLoader()));
    props.put("f", "io.confluent.common.config.Foo");
    Map<String, Object> vals = def.parse(props);
    Assert.assertNotNull(vals.get("f"));
  }

  @Test(expected = ConfigException.class)
  public void attemptLoadFooWithDefaultClassLoader() throws ClassNotFoundException {
    ConfigDef def = new ConfigDef().define("f", CLASS, ConfigDef.Importance.HIGH, "docs");
    Properties props = new Properties();
    Thread.currentThread().setContextClassLoader(null);
    props.put("f", "io.confluent.common.config.Foo");
    def.parse(props);
    throw new ClassNotFoundException("Expected to raise a ClassNotFoundException.");
  }

  /**
   * A classloader which can load a io.confluent.common.config.Foo class.
   */
  private static class FooLoader extends ClassLoader {

    // A Base64 encoding of the bytecode of a class with name io.confluent.common.config.Foo.
    static String BASE64 =
            "yv66vgAAADMAEAoAAwANBwAOBwAPAQAGPGluaXQ+AQADKClWAQAEQ29kZQEAD0xpbmVOdW1iZXJUYWJsZQEAEkxvY2FsVmFyaWFi" +
                    "bGVUYWJsZQEABHRoaXMBACBMaW8vY29uZmx1ZW50L2NvbW1vbi9jb25maWcvRm9vOwEAClNvdXJjZUZpbGUBAAhGb28ua" +
                    "mF2YQwABAAFAQAeaW8vY29uZmx1ZW50L2NvbW1vbi9jb25maWcvRm9vAQAQamF2YS9sYW5nL09iamVjdAAhAAIAAwAAAA" +
                    "AAAQABAAQABQABAAYAAAAvAAEAAQAAAAUqtwABsQAAAAIABwAAAAYAAQAAAAMACAAAAAwAAQAAAAUACQAKAAAAAQALAAA" +
                    "AAgAM";

    FooLoader(ClassLoader parent) {
      super(parent);
    }

    @Override
    public Class<?> findClass(String name) throws ClassNotFoundException {
      try {
        return getParent().loadClass("io.confluent.common.config.Foo");
      } catch (ClassNotFoundException e) {
        if (name.equals("io.confluent.common.config.Foo")) {
          byte[] bytes = DatatypeConverter.parseBase64Binary(BASE64);
          return defineClass("io.confluent.common.config.Foo", bytes, 0, bytes.length);
        } else {
          throw new ClassNotFoundException("Could not find class " + name);
        }
      }
    }
  }
}
