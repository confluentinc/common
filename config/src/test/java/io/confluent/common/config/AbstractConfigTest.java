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

import java.util.HashMap;
import java.util.Map;
import java.util.Properties;

import static org.junit.Assert.*;

/**
 * Unit tests for class {@link AbstractConfig}.
 *
 * @see AbstractConfig
 */
public class AbstractConfigTest {

  @Test
  public void testGetPropsFromFileReturningEmptyProperties() {
    Properties properties = AbstractConfig.getPropsFromFile(null);

    assertTrue(properties.isEmpty());
  }

  @Test
  public void testWithPrefixWithNonEmptyStringReturnsEmptyMap() {
    ConfigDef configDef = new ConfigDef();
    Map<Integer, Object> hashMap = new HashMap<>();
    AbstractConfig abstractConfig = new AbstractConfig(configDef, hashMap);
    Map<String, String> hashMapTwo = new HashMap<>();
    hashMapTwo.put(",..", ",..");
    Map<String, Object> map = abstractConfig.withPrefix(",..", hashMapTwo);

    assertEquals(0, map.size());
  }

  @Test
  public void testWithPrefixWithNonEmptyStrings() {
    ConfigDef configDef = new ConfigDef();
    Map<String, Object> hashMap = new HashMap<>();
    hashMap.put("- asd1", "-asd");
    hashMap.put(" asd1", "-asd");
    AbstractConfig abstractConfig = new AbstractConfig(configDef, hashMap);
    Map<String, Object> map = abstractConfig.withPrefix("-", hashMap);

    assertEquals(1, map.size());
    assertEquals(" asd1", map.keySet().iterator().next());
    assertEquals(" asd1=-asd", map.entrySet().iterator().next().toString());
  }

  @Test
  public void testFailsToCreateAbstractConfigThrowsConfigException() {
    ConfigDef configDef = new ConfigDef();
    Map<Object, Long> hashMap = new HashMap<>();
    hashMap.put(configDef, new Long((-1995L)));

    try {
      new AbstractConfig(configDef, hashMap);
      fail("Expecting exception: ConfigException");
    } catch (ConfigException e) {
      assertEquals(AbstractConfig.class.getName(), e.getStackTrace()[0].getClassName());
    }
  }

  @Test
  public void testGetStringThrowsConfigException() {
    ConfigDef configDef = new ConfigDef();
    Map<Integer, String> hashMap = new HashMap<>();
    AbstractConfig abstractConfig = new AbstractConfig(configDef, hashMap);

    try {
      abstractConfig.getString(" could not be found.");
      fail("Expecting exception: ConfigException");
    } catch (ConfigException e) {
      assertEquals(AbstractConfig.class.getName(), e.getStackTrace()[0].getClassName());
    }
  }

  @Test
  public void testValuesWithPrefix() {
    ConfigDef configDef = new ConfigDef();
    Map<Object, Boolean> hashMap = new HashMap<>();
    AbstractConfig abstractConfig = new AbstractConfig(configDef, hashMap);
    Map<String, Object> map = abstractConfig.valuesWithPrefix("");

    assertEquals(0, map.size());
  }

  @Test
  public void testOriginalsWithPrefix() {
    ConfigDef configDef = new ConfigDef();
    Map<Integer, String> hashMap = new HashMap<>();
    AbstractConfig abstractConfig = new AbstractConfig(configDef, hashMap);
    Map<String, Object> map = abstractConfig.originalsWithPrefix("~cOo+qt{k]k}#%y");

    assertEquals(0, map.size());
  }

}