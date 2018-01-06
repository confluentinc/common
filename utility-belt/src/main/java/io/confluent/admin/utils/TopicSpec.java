/*
 * Copyright 2017 Confluent Inc.
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.confluent.kafkaensure;

import java.util.List;
import java.util.Map;
import java.util.Objects;

public class TopicSpec {
  public String name;
  public int partitions;
  public int replicationFactor;
  public Map<String, String> config;

  public TopicSpec() {
  }

  public TopicSpec(String name, int partitions, int replicationFactor, Map<String, String> config) {
    this.name = name;
    this.partitions = partitions;
    this.replicationFactor = replicationFactor;
    this.config = config;
  }

  public String name() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public int partitions() {
    return partitions;
  }

  public void setPartitions(int partitions) {
    this.partitions = partitions;
  }

  public int replicationFactor() {
    return replicationFactor;
  }

  public void setReplicationFactor(int replicationFactor) {
    this.replicationFactor = replicationFactor;
  }

  public Map<String, String> config() {
    return config;
  }

  public void setConfig(Map<String, String> config) {
    this.config = config;
  }

  public void putInConfig(String key, String value) {
    this.config.put(key, value);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TopicSpec topicSpec = (TopicSpec) o;
    return partitions == topicSpec.partitions
        && replicationFactor == topicSpec.replicationFactor
        && Objects.equals(name, topicSpec.name)
        && Objects.equals(config, topicSpec.config);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, partitions, replicationFactor, config);
  }

  @Override
  public String toString() {
    return "TopicSpec={"
        + "name='" + name() + '\''
        + ", partitions='" + partitions() + '\''
        + ", replicationFactor=" + replicationFactor()
        + ", config=" + config()
        + '}';
  }

  public static class Topics {
    public List<TopicSpec> topics;

    public Topics() {
    }

    public List<TopicSpec> topics() {
      return topics;
    }

    public void setTopics(List<TopicSpec> topics) {
      this.topics = topics;
    }

    @Override
    public String toString() {
      return "Topics{"
          + "topics=" + topics
          + '}';
    }
  }
}
