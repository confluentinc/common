/*
 * Copyright 2017 Confluent Inc.
 */

package io.confluent.kafkaensure.cli;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;

import net.sourceforge.argparse4j.ArgumentParsers;
import net.sourceforge.argparse4j.inf.ArgumentParser;
import net.sourceforge.argparse4j.inf.ArgumentParserException;
import net.sourceforge.argparse4j.inf.Namespace;

import org.apache.kafka.common.utils.Utils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;

import io.confluent.kafkaensure.TopicEnsure;
import io.confluent.kafkaensure.TopicSpec;

import static net.sourceforge.argparse4j.impl.Arguments.store;

/**
 * This command ensures that a topic exists and has valid config.
 * where:
 * config                 : path to properties with client config.
 * file                   : file with topic spec.
 * timeout                : timeout in ms for all operations.
 */
public class TopicEnsureCommand {

  private static final Logger log = LoggerFactory.getLogger(TopicEnsureCommand.class);
  public static final String TOPIC_ENSURE = "topic-ensure";

  private static ArgumentParser createArgsParser() {
    ArgumentParser topicEnsure = ArgumentParsers
        .newArgumentParser(TOPIC_ENSURE)
        .defaultHelp(true)
        .description("Check if topic exists and is valid.");

    topicEnsure.addArgument("--timeout", "-t")
        .action(store())
        .required(true)
        .type(Integer.class)
        .metavar("TIMEOUT_IN_MS")
        .help("Time (in ms) to wait for service to be ready.");

    topicEnsure.addArgument("--config", "-c")
        .action(store())
        .type(String.class)
        .metavar("CONFIG")
        .required(true)
        .help("Client config.");

    topicEnsure.addArgument("--file", "-f")
        .action(store())
        .type(String.class)
        .metavar("FILE_CONFIG")
        .required(true)
        .help("Topic config file.");

    topicEnsure.addArgument("--create-if-not-exists")
        .action(store())
        .type(Boolean.class)
        .setDefault(false)
        .help("Create topic if it does not exist.");

    return topicEnsure;
  }

  public static void main(String[] args) {

    ArgumentParser parser = createArgsParser();
    boolean success = false;
    try {
      Namespace res = parser.parseArgs(args);
      log.debug("Arguments {}. ", res);

      TopicEnsure topicEnsure = new TopicEnsure(Utils.loadProps(res.getString("config")));
      ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
      TopicSpec.Topics topics = mapper.readValue(
          new File(res.getString("file")), TopicSpec.Topics.class
      );

      for (TopicSpec spec : topics.topics()) {
        success = topicEnsure.topicExists(spec, res.getInt("timeout"));
        System.err.printf("Topic [ %s ] exists ? %s\n", spec.name(), success);
        if (success) {
          success = topicEnsure.validateTopic(spec, res.getInt("timeout"));
          System.err.printf("Topic spec [ %s ] valid ? %s\n", spec, success);
        } else if (res.getBoolean("create_if_not_exists")) {
          success = topicEnsure.createTopic(spec, res.getInt("timeout"));
          System.err.printf("Topic [ %s ] created with spec: [ %s ] \n", spec.name(), spec);
        }
      }
    } catch (ArgumentParserException e) {
      if (args.length == 0) {
        parser.printHelp();
        success = true;
      } else {
        parser.handleError(e);
        success = false;
      }
    } catch (Exception e) {
      log.error("Error while running topic-ensure {}.", e);
      success = false;
    }

    if (success) {
      System.exit(0);
    } else {
      System.exit(1);
    }
  }
}
