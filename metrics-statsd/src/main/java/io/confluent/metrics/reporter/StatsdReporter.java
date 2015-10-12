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
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.confluent.metrics.reporter;

import io.confluent.common.metrics.KafkaMetric;
import io.confluent.common.metrics.MetricsReporter;

import java.util.List;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.timgroup.statsd.NonBlockingStatsDClient;
import com.timgroup.statsd.StatsDClient;

public class StatsdReporter implements MetricsReporter {

	public static final Logger LOG = LoggerFactory
			.getLogger(StatsdReporter.class);

	public static final String STATSD_HOST_NAME = "metrics.statsd.host";
	public static final String STATSD_PORT_NAME = "metrics.statsd.port";
	public static final String STATSD_PREFIX_NAME = "metrics.statsd.prefix";
	
	String statsdHost;
	int statsdPort = 8125;
	String statsdPrefix = "schema.metrics";
	transient StatsDClient statsd;
	
	public StatsdReporter(){
		
	}
	
	public StatsdReporter(StatsDClient clientimpl){
		statsd = clientimpl;
	}
	
	@Override
	public void configure(Map<String, ?> configs) {

		if (configs.containsKey(STATSD_HOST_NAME)) {
			statsdHost = (String) configs.get(STATSD_HOST_NAME);
		}

		if (configs.containsKey(STATSD_PORT_NAME)) {
			statsdPort = Integer.parseInt((String) configs.get(STATSD_PORT_NAME));
		}

		if (configs.containsKey(STATSD_PREFIX_NAME)) {
			statsdPrefix = (String) configs.get(STATSD_PREFIX_NAME);
		}

		System.out.println("Configure statsd host properties :  " + statsdHost + " " + statsdPrefix );

	}

	@Override
	public void init(List<KafkaMetric> metrics) {
		if(statsd == null){
			statsd = new NonBlockingStatsDClient(statsdPrefix, statsdHost,
				statsdPort);
		}	
	}
	
	@Override
	public void metricChange(KafkaMetric metric) {
		report(metric.metricName().name(), (int) metric.value());
	}

	private void report(String s, int number) {
		LOG.info("reporting: {}={}", s, number);
		statsd.gauge(s, number);
	}

	@Override
	public void close() {
		LOG.info("Stopping Statsd Client");
		statsd.stop();

	}
	

}
