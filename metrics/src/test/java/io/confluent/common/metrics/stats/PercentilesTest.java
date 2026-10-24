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
package io.confluent.common.metrics.stats;

import io.confluent.common.metrics.MetricConfig;
import io.confluent.common.metrics.MetricName;
import org.junit.Test;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.fail;

/**
 * Unit tests for class {@link Percentiles}.
 *
 * @see Percentiles
 */
public class PercentilesTest {

  @Test
  public void testValue() {
    Percentiles.BucketSizing percentiles_BucketSizing = Percentiles.BucketSizing.LINEAR;
    Percentiles percentiles = new Percentiles(1846, 1846, percentiles_BucketSizing, new Percentile[0]);
    MetricConfig metricConfig = new MetricConfig();
    percentiles.record(metricConfig, 28L, 1792L);

    assertEquals(Double.POSITIVE_INFINITY, percentiles.value(metricConfig, 1846, 27.78662642648307), 0.01);
  }

  @Test
  public void testFailsToCreatePercentilesTaking5ArgumentsThrowsIllegalArgumentException() {
    try {
      new Percentiles(44, (-121.169), 44, Percentiles.BucketSizing.LINEAR, null);
      fail("Expecting exception: IllegalArgumentException");
    } catch (IllegalArgumentException e) {
      assertEquals(Percentiles.class.getName(), e.getStackTrace()[0].getClassName());
    }
  }

  @Test
  public void testFailsToCreatePercentilesTaking4ArgumentsThrowsIllegalArgumentException() {
    try {
      new Percentiles(58, 58, null, new Percentile[7]);
      fail("Expecting exception: IllegalArgumentException");
    } catch (IllegalArgumentException e) {
      assertEquals(Percentiles.class.getName(), e.getStackTrace()[0].getClassName());
    }
  }

  @Test
  public void testCombineWithEmptyList() {
    Percentiles.BucketSizing percentiles_BucketSizing = Percentiles.BucketSizing.LINEAR;
    Percentile[] percentileArray = new Percentile[0];
    Percentiles percentiles = new Percentiles(146, 1846, percentiles_BucketSizing, percentileArray);
    MetricConfig metricConfig = new MetricConfig();
    List<SampledStat.Sample> list = percentiles.samples;

    assertEquals(Double.NaN, percentiles.combine(list, metricConfig, 1846), 0.01);
  }

  @Test
  public void testStatsThrowsNullPointerException() {
    Percentiles.BucketSizing percentiles_BucketSizing = Percentiles.BucketSizing.LINEAR;
    Percentile[] percentileArray = new Percentile[2];
    Map<String, String> hashMap = new HashMap<>();
    MetricName metricName = new MetricName("|~a`", "|~a`", hashMap);
    Percentile percentile = new Percentile(metricName, (-961));
    percentileArray[0] = percentile;
    Percentiles percentiles = new Percentiles((-961), 0.0, percentiles_BucketSizing, percentileArray);

    try {
      percentiles.stats();
      fail("Expecting exception: NullPointerException");
    } catch (NullPointerException e) {
      assertEquals(Percentiles.class.getName(), e.getStackTrace()[0].getClassName());
    }
  }
}