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
package io.confluent.common.metrics.stats;

import java.util.List;

import io.confluent.common.metrics.MetricConfig;

/**
 * A {@link SampledStat} that maintains a simple average over its samples.
 */
public class Avg extends SampledStat {

  public Avg() {
    super(0.0);
  }

  @Override
  protected void update(Sample sample, MetricConfig config, double value, long now) {
    sample.value += value;
  }

  @Override
  public double combine(List<Sample> samples, MetricConfig config, long now) {
    double total = 0.0;
    long count = 0;
    for (Sample sample : samples) {
      total += sample.value;
      count += sample.eventCount;
    }
    return total / count;
  }

}
