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
package io.confluent.common.metrics;

import io.confluent.common.metrics.stats.Avg;
import io.confluent.common.metrics.stats.Total;
import io.confluent.metrics.reporter.StatsdReporter;

import org.junit.Test;

import com.timgroup.statsd.NoOpStatsDClient;

public class StatsdReporterTest {
	
	
	@Test
	  public void testStatsdReporter() throws Exception {
	    Metrics metrics = new Metrics();
	    metrics.addReporter(new StatsdReporter(new NoOpStatsDClient()));
	    Sensor sensor = metrics.sensor("schema.requests");
	    sensor.add(new MetricName("get.id.avg", "grp1"), new Avg());
	    sensor.add(new MetricName("get.id.total", "grp2"), new Total());
	    Sensor sensor2 = metrics.sensor("schema.error");
	    sensor2.add(new MetricName("get.id.error", "grp1"), new Total());
	    sensor2.add(new MetricName("get.id.test", "grp1"), new Total());
	  }
	
}
