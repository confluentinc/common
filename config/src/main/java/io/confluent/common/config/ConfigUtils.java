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

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Enumeration;
import java.util.HashSet;
import java.util.List;
import java.util.Properties;

public class ConfigUtils {
  private final static Logger log = LoggerFactory.getLogger(ConfigUtils.class);

  /**
   * Handle deprecated properties by translating them into their non-deprecated
   * equivalents.<p/>
   *
   * Given a Properties object, create a new Properties object
   *
   * @param props           The input Properties object.
   * @param synonymGroups   An array of arrays of synonyms.  Each synonym array begins with the
   *                        non-deprecated synonym.
   *                        For example, new String[][] { { a, b }, { c, d, e} } would declare
   *                        b as a deprecated synonym for a, and d and e as deprecated synonyms
   *                        for c.
   * @return                A new Properties object with deprecated  keys translated to their
   *                        non-deprecated equivalents.
   */
  public static Properties translateDeprecated(Properties props, String[][] synonymGroups) {
    // Copy the Properties which are not part of a synonym group into a new
    // Properties object.
    HashSet<String> synonymSet = new HashSet<>();
    for (String[] synonymGroup: synonymGroups) {
      for (String synonym : synonymGroup) {
        if (!synonymSet.add(synonym)) {
          // TODO: we should consider doing more formal validations around this
          log.warn("reused synonym={}", synonym);
        }
      }
    }
    // Properties is a very old Java class which uses Enumeration to iterate through keys rather
    // than an iterator.  The keys are always String, but we have to deal with an
    // Enumeration<Object> anyway, probably for historical purposes.
    Properties newProps = new Properties();
    for (Enumeration<String> keyEnumerator = (Enumeration<String>) props.propertyNames();
         keyEnumerator.hasMoreElements(); ) {
      String key = keyEnumerator.nextElement();
      if (!synonymSet.contains(key)) {
        newProps.put(key, props.get(key));
      }
    }
    // Process each synonym group.
    for (String[] synonymGroup: synonymGroups) {
      String target = synonymGroup[0];
      List<String> deprecated = new ArrayList<>();
      for (int i = 1; i < synonymGroup.length; i++) {
        if (props.containsKey(synonymGroup[i])) {
          deprecated.add(synonymGroup[i]);
        }
      }
      if (deprecated.isEmpty()) {
        // No deprecated key(s) found.
        if (props.containsKey(target)) {
          newProps.put(target, props.get(target));
        }
        continue;
      }
      StringBuilder synonymString = new StringBuilder(deprecated.get(0));
      for (int i = 1; i < deprecated.size(); i++) {
        synonymString.append(", ");
        synonymString.append(deprecated.get(i));
      }
      if (props.containsKey(target)) {
        // Ignore the deprecated key(s) because the actual key was set.
        log.error(target + " was configured, as well as the deprecated synonym(s) " +
          synonymString + ".  Using the value of " + target);
        newProps.put(target, props.get(target));
      } else if (deprecated.size() > 1) {
        log.error("The configuration keys " + synonymString + " are deprecated and may be " +
          "removed in the future.  Additionally, this configuration is ambigous because " +
          "these configuration keys are all synonyms for " + target + ".  Please update " +
          "your configuration to have only " + target + " set.");
        newProps.put(target, props.get(deprecated.get(0)));
      } else {
        log.warn("Configuration key " + deprecated.get(0) + " is deprecated and may be removed " +
          "in the future.  Please update your configuration to use " + target + " instead.");
        newProps.put(target, props.get(deprecated.get(0)));
      }
    }
    return newProps;
  }
}
