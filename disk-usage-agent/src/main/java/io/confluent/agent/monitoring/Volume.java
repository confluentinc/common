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

import java.io.IOException;
import java.nio.file.FileStore;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Volume implements VolumeMBean {
  private static final Logger log = LoggerFactory.getLogger(Volume.class);

  private final String dir;
  private final FileStore store;

  public Volume(String dir) throws IOException {
    this.dir = dir;
    store = Files.getFileStore(Paths.get(dir));
  }

  @Override
  public long getTotal() {
    long total = -1;
    try {
      total = this.store.getTotalSpace();
    } catch (IOException e) {
      log.error("Error while getting total disk space", e);
    }
    return total;
  }

  @Override
  public long getUsed() {
    long used = -1;
    try {
      used = (this.store.getTotalSpace() - this.store.getUnallocatedSpace());
    } catch (IOException e) {
      log.error("Error while getting total disk space", e);
    }
    return used;
  }

  @Override
  public long getAvailable() {
    long avail = -1;
    try {
      avail = this.store.getUsableSpace();
    } catch (IOException e) {
      log.error("Error while getting total disk space", e);
    }
    return avail;
  }

  @Override
  public double getPercentUsed() {
    return ((double)this.getUsed() / (double)this.getTotal()) * 100;
  }

  @Override
  public double getPercentAvailable() {
    return ((double)this.getAvailable() / (double)this.getTotal()) * 100;
  }

  @Override
  public String getMountpoint() {
    return dir;
  }

  @Override
  public String getDeviceName() {
    return this.store.name();
  }

  @Override
  public String toString() {
    return "Volume{" +
        "store=" + store.toString() +
        ", total=" + getTotal() +
        ", used=" + getUsed() +
        ", available=" + getAvailable() +
        ", percentUsed=" + getPercentUsed() +
        ", percentAvailable=" + getPercentAvailable() +
        ", mountpoint='" + getMountpoint() + '\'' +
        ", deviceName='" + getDeviceName() + '\'' +
        '}';
  }

}
