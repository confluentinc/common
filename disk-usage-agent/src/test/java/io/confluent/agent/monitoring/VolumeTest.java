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

import org.junit.Test;
import org.junit.Ignore;
import org.junit.rules.TemporaryFolder;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;
import static org.junit.Assert.fail;


public class VolumeTest {

  @Ignore @Test
  public void testVolume() {

    try {
      Volume volume = new Volume(System.getProperty("java.io.tmpdir"));

      long total = volume.getTotal();
      long used = volume.getUsed();
      long available = volume.getAvailable();
      double percentAvailable = volume.getPercentAvailable();
      double percentUsed = volume.getPercentUsed();

      assertNotNull(volume.getDeviceName());
      assertEquals(volume.getMountpoint(), System.getProperty("java.io.tmpdir"));

      assertTrue( total > 0);
      assertTrue(used > 0);
      assertTrue(available > 0);
      assertEquals(total, used + available, total * 0.05);

      assertTrue(percentAvailable > 0);
      assertTrue(percentUsed > 0);
      assertEquals(100.0, percentUsed + percentAvailable, 0.2);
    } catch (IOException e) {
      fail("Should not fail.");
    }
  }

  @Test
  public void testDiskUsage() {

    try {
      TemporaryFolder tempFolder = new TemporaryFolder();
      String config = "service.name=kafka\ndisk.logs=/tmp";
      tempFolder.create();
      final File tempFile = tempFolder.newFile("test.properties");
      Files.write(tempFile.toPath(), config.getBytes());
      DiskUsage.premain(tempFile.getAbsolutePath(), null);
    } catch (Exception e) {
      e.printStackTrace();
      fail("Should not fail.");
    }
  }
}
