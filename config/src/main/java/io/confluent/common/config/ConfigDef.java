/**
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

/**
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

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;

import io.confluent.common.config.types.Password;

/**
 * This class is used for specifying the set of expected configurations, their type, their defaults,
 * their documentation, and any special validation logic used for checking the correctness of the
 * values the user provides. <p> Usage of this class looks something like this:
 *
 * <pre>
 * ConfigDef defs = new ConfigDef();
 * defs.define(&quot;config_name&quot;, Type.STRING, &quot;default string value&quot;,
 *             &quot;This configuration is used for blah blah blah.&quot;);
 * defs.define(&quot;another_config_name&quot;, Type.INT, 42, Range.atLeast(0),
 *             &quot;More documentation on this config&quot;);
 *
 * Properties props = new Properties();
 * props.setProperty(&quot;config_name&quot;, &quot;some value&quot;);
 * Map&lt;String, Object&gt; configs = defs.parse(props);
 *
 * // will return &quot;some value&quot;
 * String someConfig = (String) configs.get(&quot;config_name&quot;);
 * // will return default value of 42
 * int anotherConfig = (Integer) configs.get(&quot;another_config_name&quot;);
 * </pre>
 *
 * This class can be used stand-alone or in combination with {@link AbstractConfig} which provides
 * some additional functionality for accessing configs.
 */
public class ConfigDef {

  protected static final Object NO_DEFAULT_VALUE = new Object();

  private final Map<String, ConfigKey> configKeys;

  public ConfigDef() {
    configKeys = new HashMap<>();
  }

  public ConfigDef(ConfigDef configDef) {
    configKeys = new HashMap<>(configDef.configKeys);
  }

  /**
   * Returns unmodifiable set of properties names defined in this {@linkplain ConfigDef}
   *
   * @return new unmodifiable {@link Set} instance containing the keys
   */
  public Set<String> names() {
    return Collections.unmodifiableSet(configKeys.keySet());
  }

  /**
   * Define a new configuration
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param defaultValue  The default value to use if this config isn't present
   * @param validator     A validator to use in checking the correctness of the config
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef define(String name, Type type, Object defaultValue, Validator validator,
                          Importance importance, String documentation) {
    if (configKeys.containsKey(name)) {
      throw new ConfigException("Configuration " + name + " is defined twice.");
    }
    Object parsedDefault =
        NO_DEFAULT_VALUE.equals(defaultValue) ? NO_DEFAULT_VALUE : parseType(name, defaultValue, type);
    configKeys
        .put(name, new ConfigKey(name, type, parsedDefault, validator, importance, documentation));
    return this;
  }

  /**
   * Override a configuration
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param defaultValue  The default value to use if this config isn't present
   * @param validator     A validator to use in checking the correctness of the config
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef defineOverride(String name, Type type, Object defaultValue, Validator validator,
                                  Importance importance, String documentation) {
    if (!configKeys.containsKey(name)) {
      throw new ConfigException("Configuration " + name +
                                " is defined as an override but does not exist.");
    }
    Object parsedDefault =
            NO_DEFAULT_VALUE.equals(defaultValue) ? NO_DEFAULT_VALUE : parseType(name, defaultValue, type);
    configKeys
        .put(name, new ConfigKey(name, type, parsedDefault, validator, importance, documentation));
    return this;
  }

  /**
   * Define a new configuration with no special validation logic
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param defaultValue  The default value to use if this config isn't present
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef define(String name, Type type, Object defaultValue, Importance importance,
                          String documentation) {
    return define(name, type, defaultValue, null, importance, documentation);
  }

  /**
   * Override a configuration with no special validation logic
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param defaultValue  The default value to use if this config isn't present
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef defineOverride(String name, Type type, Object defaultValue, Importance importance,
                          String documentation) {
    return defineOverride(name, type, defaultValue, null, importance, documentation);
  }

  /**
   * Define a required parameter with no default value
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param validator     A validator to use in checking the correctness of the config
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef define(String name, Type type, Validator validator, Importance importance,
                          String documentation) {
    return define(name, type, NO_DEFAULT_VALUE, validator, importance, documentation);
  }

  /**
   * Override a required parameter with no default value
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param validator     A validator to use in checking the correctness of the config
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef defineOverride(String name, Type type, Validator validator, Importance importance,
                          String documentation) {
    return defineOverride(name, type, NO_DEFAULT_VALUE, validator, importance, documentation);
  }

  /**
   * Define a required parameter with no default value and no special validation logic
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef define(String name, Type type, Importance importance, String documentation) {
    return define(name, type, NO_DEFAULT_VALUE, null, importance, documentation);
  }

  /**
   * Define a required parameter with no default value and no special validation logic
   *
   * @param name          The name of the config parameter
   * @param type          The type of the config
   * @param importance    The importance of this config: is this something you will likely need to
   *                      change.
   * @param documentation The documentation string for the config
   * @return This ConfigDef so you can chain calls
   */
  public ConfigDef defineOverride(String name, Type type, Importance importance, String documentation) {
    return defineOverride(name, type, NO_DEFAULT_VALUE, null, importance, documentation);
  }

  /**
   * Parse and validate configs against this configuration definition. The input is a map of
   * configs. It is expected that the keys of the map are strings, but the values can either be
   * strings or they may already be of the appropriate type (int, string, etc). This will work
   * equally well with either java.util.Properties instances or a programmatically constructed map.
   *
   * @param props The configs to parse and validate
   * @return Parsed and validated configs. The key will be the config name and the value will be the
   * value parsed into the appropriate type (int, string, etc)
   */
  public Map<String, Object> parse(Map<?, ?> props) {
    /* parse all known keys */
    Map<String, Object> values = new HashMap<>();
    for (ConfigKey key : configKeys.values()) {
      Object value;
      if (props.containsKey(key.name)) {
        value = parseType(key.name, props.get(key.name), key.type);
      } else if (NO_DEFAULT_VALUE.equals(key.defaultValue)) {
        throw new ConfigException(
            "Missing required configuration \"" + key.name + "\" which has no default value.");
      } else {
        value = key.defaultValue;
      }
      if (key.validator != null) {
        key.validator.ensureValid(key.name, value);
      }
      values.put(key.name, value);
    }
    return values;
  }

  /**
   * Parse a value according to its expected type.
   *
   * @param name  The config name
   * @param value The config value
   * @param type  The expected type
   * @return The parsed object
   */
  private Object parseType(String name, Object value, Type type) {
    try {
      String trimmed = null;
      if (value instanceof String) {
        trimmed = ((String) value).trim();
      }
      switch (type) {
        case BOOLEAN:
          if (value instanceof String) {
            if ("true".equalsIgnoreCase(trimmed)) {
              return true;
            } else if ("false".equalsIgnoreCase(trimmed)) {
              return false;
            } else {
              throw new ConfigException(name, value, "Expected value to be either true or false");
            }
          } else if (value instanceof Boolean) {
            return value;
          } else {
            throw new ConfigException(name, value, "Expected value to be either true or false");
          }
        case STRING:
          if (value instanceof String) {
            return trimmed;
          } else {
            throw new ConfigException(name, value,
                                      "Expected value to be a string, but it was a " + value
                                          .getClass().getName());
          }
        case PASSWORD:
          if (value instanceof Password)
            return value;
          else if (value instanceof String)
            return new Password(trimmed);
          else
            throw new ConfigException(name, value, "Expected value to be a string, but it was a " + value.getClass().getName());
        case INT:
          if (value instanceof Integer) {
            return value;
          } else if (value instanceof String) {
            return Integer.parseInt(trimmed);
          } else {
            throw new ConfigException(name, value, "Expected value to be an number.");
          }
        case LONG:
          if (value instanceof Integer) {
            return ((Integer) value).longValue();
          }
          if (value instanceof Long) {
            return value;
          } else if (value instanceof String) {
            return Long.parseLong(trimmed);
          } else {
            throw new ConfigException(name, value, "Expected value to be an number.");
          }
        case DOUBLE:
          if (value instanceof Number) {
            return ((Number) value).doubleValue();
          } else if (value instanceof String) {
            return Double.parseDouble(trimmed);
          } else {
            throw new ConfigException(name, value, "Expected value to be an number.");
          }
        case LIST:
          if (value instanceof List) {
            return value;
          } else if (value instanceof String) {
            if (trimmed.isEmpty()) {
              return Collections.emptyList();
            } else {
              return Arrays.asList(trimmed.split("\\s*,\\s*", -1));
            }
          } else {
            throw new ConfigException(name, value, "Expected a comma separated list.");
          }
        case MAP:
          if (value instanceof Map) {
            return value;
          } else if (value instanceof String) {
            if (trimmed.isEmpty()) {
              return Collections.emptyMap();
            } else {
              Map<String, String> map = new HashMap<>();
              List<String> entries = Arrays.asList(trimmed.split("\\s*,\\s*", -1));
              for (String entry : entries) {
                String[] keyValue = entry.split("\\s*:\\s*", -1);
                if (keyValue.length != 2) {
                  throw new ConfigException("Map entry should be of form <key>:<value");
                }
                map.put(keyValue[0], keyValue[1]);
              }
              return map;
            }
          } else {
            throw new ConfigException(name, value, "Expected a comma separated Map entries.");
          }
        case CLASS:
          if (value instanceof Class) {
            return value;
          } else if (value instanceof String) {
            if (Thread.currentThread().getContextClassLoader() != null) {
              return Class.forName(trimmed, true, Thread.currentThread().getContextClassLoader());
            } else {
              return Class.forName(trimmed);
            }
          } else {
            throw new ConfigException(name, value, "Expected a Class instance or class name.");
          }
        default:
          throw new IllegalStateException("Unknown type.");
      }
    } catch (NumberFormatException e) {
      throw new ConfigException(name, value, "Not a number of type " + type);
    } catch (ClassNotFoundException e) {
      throw new ConfigException(name, value, "Class " + value + " could not be found.");
    }
  }

  /**
   * The config types
   */
  public enum Type {
    BOOLEAN, STRING, INT, LONG, DOUBLE, LIST, CLASS, PASSWORD, MAP
  }

  public enum Importance {
    HIGH, MEDIUM, LOW
  }

  /**
   * Validation logic the user may provide
   */
  public interface Validator {

    public void ensureValid(String name, Object o);
  }

  /**
   * Validation logic for numeric ranges
   */
  public static class Range implements Validator {

    private final Number min;
    private final Number max;

    private Range(Number min, Number max) {
      this.min = min;
      this.max = max;
    }

    /**
     * A numeric range that checks only the lower bound
     *
     * @param min The minimum acceptable value
     */
    public static Range atLeast(Number min) {
      return new Range(min, null);
    }

    /**
     * A numeric range that checks both the upper and lower bound
     */
    public static Range between(Number min, Number max) {
      return new Range(min, max);
    }

    @Override
    public void ensureValid(String name, Object o) {
      Number n = (Number) o;
      if (min != null && n.doubleValue() < min.doubleValue()) {
        throw new ConfigException(name, o, "Value must be at least " + min);
      }
      if (max != null && n.doubleValue() > max.doubleValue()) {
        throw new ConfigException(name, o, "Value must be no more than " + max);
      }
    }

    @Override
    public String toString() {
      if (min == null) {
        return "[...," + max + "]";
      } else if (max == null) {
        return "[" + min + ",...]";
      } else {
        return "[" + min + ",...," + max + "]";
      }
    }
  }

  public static class ValidString implements Validator {

    List<String> validStrings;

    private ValidString(List<String> validStrings) {
      this.validStrings = validStrings;
    }

    public static ValidString in(List<String> validStrings) {
      return new ValidString(validStrings);
    }

    @Override
    public void ensureValid(String name, Object o) {

      String s = (String) o;

      if (!validStrings.contains(s)) {
        throw new ConfigException(name, o, "String must be one of:" + join(validStrings));
      }

    }

    @Override
    public String toString() {
      return "[" + join(validStrings) + "]";
    }

    private String join(List<String> list) {
      StringBuilder sb = new StringBuilder();
      boolean first = true;
      for (String item : list) {
        if (first) {
          first = false;
        } else {
          sb.append(",");
        }
        sb.append(item);
      }
      return sb.toString();
    }
  }

  private static class ConfigKey {

    public final String name;
    public final Type type;
    public final String documentation;
    public final Object defaultValue;
    public final Validator validator;
    public final Importance importance;

    public ConfigKey(String name, Type type, Object defaultValue, Validator validator,
                     Importance importance, String documentation) {
      super();
      this.name = name;
      this.type = type;
      this.defaultValue = defaultValue;
      this.validator = validator;
      this.importance = importance;
      if (this.validator != null && !NO_DEFAULT_VALUE.equals(defaultValue)) {
        this.validator.ensureValid(name, defaultValue);
      }
      this.documentation = documentation;
    }

    public boolean hasDefault() {
      return !NO_DEFAULT_VALUE.equals(this.defaultValue);
    }

  }

  public String toHtmlTable() {
    List<ConfigDef.ConfigKey> configs = sortedConfigs();
    StringBuilder b = new StringBuilder();
    b.append("<table>\n");
    b.append("<tr>\n");
    b.append("<th>Name</th>\n");
    b.append("<th>Type</th>\n");
    b.append("<th>Default</th>\n");
    b.append("<th>Importance</th>\n");
    b.append("<th>Description</th>\n");
    b.append("</tr>\n");
    for (ConfigKey def : configs) {
      b.append("<tr>\n");
      b.append("<td>");
      b.append(def.name);
      b.append("</td>");
      b.append("<td>");
      b.append(def.type.toString().toLowerCase());
      b.append("</td>");
      b.append("<td>");
      b.append(def.defaultValue == null ? "" : def.defaultValue);
      b.append("</td>");
      b.append("<td>");
      b.append(def.importance.toString().toLowerCase());
      b.append("</td>");
      b.append("<td>");
      b.append(def.documentation);
      b.append("</td>");
      b.append("</tr>\n");
    }
    b.append("</table>");
    return b.toString();
  }

  /**
   * Get the configs formatted with reStructuredText, suitable for embedding in Sphinx
   * documentation.
   */
  public String toRst() {
    List<ConfigDef.ConfigKey> configs = sortedConfigs();
    StringBuilder b = new StringBuilder();

    for (ConfigKey def : configs) {
      b.append("``");
      b.append(def.name);
      b.append("``\n");
      for(String docLine : def.documentation.split("\n")) {
        if (docLine.length() == 0) {
          continue;
        }
        b.append("  ");
        b.append(docLine);
        b.append("\n\n");
      }
      b.append("  * Type: ");
      b.append(def.type.toString().toLowerCase());
      b.append("\n");
      if (def.defaultValue != null) {
        b.append("  * Default: ");
        if (def.type == Type.STRING) {
          b.append("\"");
          b.append(def.defaultValue);
          b.append("\"");
        } else {
          b.append(def.defaultValue);
        }
        b.append("\n");
      }
      b.append("  * Importance: ");
      b.append(def.importance.toString().toLowerCase());
      b.append("\n\n");
    }
    return b.toString();
  }

  /**
   * Get a list of configs sorted into "natural" order: listing required fields first, then
   * ordering by importance, and finally by name.
   */
  private List<ConfigDef.ConfigKey> sortedConfigs() {
    // sort first required fields, then by importance, then name
    List<ConfigDef.ConfigKey> configs = new ArrayList<>(this.configKeys.values());
    Collections.sort(configs, new Comparator<ConfigDef.ConfigKey>() {

      @Override
      public int compare(ConfigDef.ConfigKey k1, ConfigDef.ConfigKey k2) {
        // first take anything with no default value
        if (!k1.hasDefault() && k2.hasDefault()) {
          return -1;
        } else if (!k2.hasDefault() && k1.hasDefault()) {
          return 1;
        }

        // then sort by importance
        int cmp = k1.importance.compareTo(k2.importance);
        if (cmp == 0)
        // then sort in alphabetical order
        {
          return k1.name.compareTo(k2.name);
        } else {
          return cmp;
        }
      }
    });
    return configs;
  }
}
