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

import org.apache.kafka.clients.admin.AdminClient;
import org.apache.kafka.clients.admin.Config;
import org.apache.kafka.clients.admin.CreateTopicsOptions;
import org.apache.kafka.clients.admin.CreateTopicsResult;
import org.apache.kafka.clients.admin.DescribeConfigsResult;
import org.apache.kafka.clients.admin.DescribeTopicsOptions;
import org.apache.kafka.clients.admin.DescribeTopicsResult;
import org.apache.kafka.clients.admin.NewTopic;
import org.apache.kafka.clients.admin.TopicDescription;
import org.apache.kafka.common.config.ConfigResource;
import org.apache.kafka.common.errors.UnknownTopicOrPartitionException;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;
import java.util.concurrent.ExecutionException;

public class TopicEnsure {

  private final AdminClient adminClient;

  public TopicEnsure(Properties props) {
    this.adminClient = AdminClient.create(props);
  }

  public boolean createTopic(TopicSpec spec, int timeOut) throws Exception {
    NewTopic newTopic = new NewTopic(
        spec.name(), spec.partitions(), (short) spec.replicationFactor()
    );
    newTopic.configs(spec.config());
    CreateTopicsResult result = adminClient.createTopics(
        Collections.singletonList(newTopic), new CreateTopicsOptions().timeoutMs(timeOut)
    );
    result.all().get();
    return true;
  }

  public boolean validateTopic(TopicSpec spec, int timeOut) throws Exception {
    // Describe topic.
    DescribeTopicsResult topicDescribeResult = adminClient.describeTopics(
        Collections.singletonList(spec.name()), new DescribeTopicsOptions().timeoutMs(timeOut)
    );
    TopicDescription topic = topicDescribeResult.all().get().get(spec.name());

    // Get topic config.
    ConfigResource configResource = new ConfigResource(ConfigResource.Type.TOPIC, spec.name());
    DescribeConfigsResult configResult = adminClient.describeConfigs(
        Collections.singletonList(configResource)
    );
    Map<ConfigResource, Config> resultMap = configResult.all().get();
    Config config = resultMap.get(configResource);

    // Create actual TopicSpec.
    Map<String, String> actualConfig = new HashMap<>();
    for (Map.Entry<String, String> entry : spec.config().entrySet()) {
      actualConfig.put(entry.getKey(), config.get(entry.getKey()).value());
    }

    TopicSpec actualSpec = new TopicSpec(
        topic.name(), topic.partitions().size(),
        topic.partitions().get(0).replicas().size(), actualConfig
    );

    boolean isTopicValid = actualSpec.equals(spec);
    if (!isTopicValid) {
      System.err.printf(
          "Invalid topic [ %s ] ! Expected %s but got %s", spec.name(), spec, actualSpec
      );
    }

    return isTopicValid;
  }

  public boolean topicExists(TopicSpec spec, Integer timeOut) throws Exception {
    try {
      DescribeTopicsResult topicDescribeResult = adminClient.describeTopics(
          Collections.singletonList(spec.name()), new DescribeTopicsOptions().timeoutMs(timeOut)
      );
      topicDescribeResult.all().get().get(spec.name());
    } catch (ExecutionException e) {
      if (e.getCause() instanceof UnknownTopicOrPartitionException) {
        return false;
      } else {
        throw e;
      }
    }
    return true;
  }
}
