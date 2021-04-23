/*
 * Copyright 2021 Confluent Inc.
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

package io.confluent.agent.monitoring;


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.lang.instrument.Instrumentation;
import java.lang.management.ManagementFactory;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;
import java.util.Set;
import java.util.stream.Collectors;

import javax.management.MBeanServer;
import javax.management.ObjectInstance;
import javax.management.ObjectName;


public class DiskUsage {
  private static final Logger log = LoggerFactory.getLogger(DiskUsage.class);

  public static void premain(String configfile, Instrumentation instrumentation) {

    log.info("DiskUsage Agent: config : " + configfile);
    try {
      if (configfile != null && !configfile.isEmpty()) {
        Map<String, String> config = loadConfig(configfile);
        String serviceName = config.get("service.name");
        Map<String, String> dirs = config.entrySet().stream()
            .filter(map -> map.getKey().startsWith("disk."))
            .collect(Collectors.toMap(p -> p.getKey(), p -> p.getValue()));

        MBeanServer server = ManagementFactory.getPlatformMBeanServer();

        for (String dir : dirs.keySet()) {
          String displayName = dir.replace("disk.", "");
          String volume = dirs.get(dir);

          String objectName = String.format("io.confluent.caas:type=VolumeMetrics, service=%s, dir=%s", serviceName, displayName);
          ObjectName volumeMBeanName = new ObjectName(objectName);

          Volume volumeMBean = new Volume(volume);
          server.registerMBean(volumeMBean, volumeMBeanName);

          Set<ObjectInstance> instances = server.queryMBeans(new ObjectName(objectName), null);
          ObjectInstance instance = (ObjectInstance) instances.toArray()[0];

          log.info("DiskUsage Agent: Registering object :" + instance.getObjectName() + " for class : " + instance.getClassName());

          log.info("DiskUsage Agent: Ping " + volumeMBean.toString());
        }
      } else {
        log.error("Disk Agent: No config file provided.");
      }
    } catch (Exception e) {
      log.error("Error while running Disk Agent: ", e);
    }
  }

  public static Map loadConfig(String configFile) throws IOException {
    if (!Files.exists(Paths.get(configFile))) {
      throw new RuntimeException("The config file location is invalid. " + configFile);
    }
    final Properties cfg = new Properties();
    try (InputStream inputStream = new FileInputStream(configFile)) {
      cfg.load(inputStream);
    }
    return new HashMap<>(cfg);
  }

  public static void agentmain(String agentArguments, Instrumentation instrumentation) {
    premain(agentArguments, instrumentation);
  }

  public static void main(String... args) throws InterruptedException {
    if (args.length == 1) {
      premain(args[0], null);
    } else {
      System.out.println("Need to provide path to config");
    }
  }

}
