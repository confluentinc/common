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

import java.util.Arrays;
import java.util.stream.Collectors;
import org.apache.kafka.connect.json.JsonConverter;
import org.apache.logging.log4j.core.Core;
import org.apache.logging.log4j.core.Layout;
import org.apache.logging.log4j.core.config.Property;
import org.apache.logging.log4j.core.config.plugins.Plugin;
import org.apache.logging.log4j.core.config.plugins.PluginElement;
import org.apache.logging.log4j.core.config.plugins.PluginFactory;

@Plugin(
    name = "StructuredJsonLayout",
    category = Core.CATEGORY_NAME,
    elementType = Layout.ELEMENT_TYPE,
    printObject = true
)
public final class StructuredJsonLayoutPlugin {
  @PluginFactory
  public static StructuredLayout createLayout(
      @PluginElement("Properties") final Property[] properties) {
    final JsonConverter converter = new JsonConverter();
    converter.configure(
        Arrays.stream(properties).collect(
            Collectors.toMap(Property::getName, Property::getValue)
        ),
        false
    );
    return new StructuredLayout(struct -> converter.fromConnectData("", struct.schema(), struct));
  }
}
