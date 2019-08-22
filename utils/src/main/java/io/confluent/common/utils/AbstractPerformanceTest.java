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

package io.confluent.common.utils;

public abstract class AbstractPerformanceTest {

  private static final long NS_PER_MS = 1000000L;
  private static final long NS_PER_SEC = 1000 * NS_PER_MS;
  private static final long MIN_SLEEP_NS = 2 * NS_PER_MS;
  private static final long TARGET_RATE_UPDATE_PERIOD_MS = 100;

  protected PerformanceStats stats;

  protected long targetIterationRate = 0;

  public AbstractPerformanceTest(long numEvents, int reportingInterval) {
    stats = new PerformanceStats(numEvents, reportingInterval);
  }

  /** Constructor which provides a 'reasonable default' reportingInterval. */
  public AbstractPerformanceTest(long numEvents) {
    this(numEvents, 5000);
  }

  /**
   * The code that is executed once per iteration of the performance test.
   */
  protected abstract void doIteration(PerformanceStats.Callback cb);

  /**
   * Returns true if the test has reached its completion criteria.
   */
  protected abstract boolean finished(int iteration);

  /**
   * Returns true if the test is running faster than its target pace.
   */
  protected abstract boolean runningFast(int iteration, float elapsed);

  /**
   * Returns the target number of iterations that should be executed per second. A value of 0
   * indicates no rate limiting should be applied. This method will be executed periodically to
   * allow the rate to adjust when tests want a value other than iteration rate to remain
   * constant, e.g. throughput.
   * @return the target number of iterations, or 0 if there should be no limit
   */
  protected float getTargetIterationRate(int iteration, float elapsed) {
    return targetIterationRate;
  }

  /**
   * Run the performance test, trying to achieve the target rate returned by
   * {@link #getTargetIterationRate(int, float)}.
   * @throws InterruptedException if interrupted
   */
  protected void run() throws InterruptedException {
    long sleepDeficitNs = 0;
    long start = System.currentTimeMillis();
    long targetRateUpdated = start;
    float currentTargetRate = getTargetIterationRate(0, 0);
    long sleepTime = (long)(NS_PER_SEC / currentTargetRate);
    for (int i = 0; !finished(i); i++) {
      long sendStart = System.currentTimeMillis();

      /*
       * Maybe sleep a little to control throughput. Sleep time can be a bit inaccurate for
       * times < 1 ms so instead of sleeping each time instead wait until a minimum sleep time
       * accumulates (the "sleep deficit") and then make up the whole deficit in one longer sleep.
       *
       * The ordering of this code and the call to doIteration is the opposite of what you
       * might expect because we need to be able to respond immediately to quick changes in
       * target rate at the very start of the test when the user overrides #targetRate(), and
       * be careful about the iteration # and elapsed time parameters passed to #targetRate()
       * and #runningFast(). The code is simpler/requires fewer calls to System#currentTimeMillis
       * in this order.
       */
      if (currentTargetRate > 0) {
        float elapsed = (sendStart - start) / 1000.f;
        if (i == 1 || (i > 1 && sendStart - targetRateUpdated > TARGET_RATE_UPDATE_PERIOD_MS)) {
          currentTargetRate = getTargetIterationRate(i, elapsed);
          sleepTime = (long)(NS_PER_SEC / currentTargetRate);
          targetRateUpdated = sendStart;
        }
        if (i > 0 && elapsed > 0 && runningFast(i, elapsed)) {
          sleepDeficitNs += sleepTime;
          if (sleepDeficitNs >= MIN_SLEEP_NS) {
            long sleepMs = sleepDeficitNs / 1000000;
            long sleepNs = sleepDeficitNs - sleepMs * 1000000;
            Thread.sleep(sleepMs, (int) sleepNs);
            sleepDeficitNs = 0;
          }
        }
      }

      PerformanceStats.Callback cb = stats.nextCompletion(sendStart);
      doIteration(cb);
    }

    /* print final results */
    stats.printTotal();
  }

  /**
   * Run the performance test and try to hit a fixed target rate. If you use this method, you
   * should not override the {@link #getTargetIterationRate(int, float)} method.
   * @param iterationsPerSec target number of iterations per second
   * @throws InterruptedException if interrupted
   */
  protected void run(long iterationsPerSec) throws InterruptedException {
    targetIterationRate = iterationsPerSec;
    run();
  }
}
