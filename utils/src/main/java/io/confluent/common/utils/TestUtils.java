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

package io.confluent.common.utils;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;

/**
 * Utility functions to be used in unit/integration tests
 */
public class TestUtils {

  /**
   * Create a temporary directory. The directory and any contents will be deleted when the test
   * process terminates.
   */
  public static File tempDirectory() {
    final File file;
    try {
      file = Files.createTempDirectory("confluent").toFile();
    } catch (final IOException ex) {
      throw new RuntimeException("Failed to create a temp dir", ex);
    }
    file.deleteOnExit();

    Runtime.getRuntime().addShutdownHook(new Thread() {
      @Override
      public void run() {
        try {
          Utils.delete(file);
        } catch (IOException e) {
          System.out.println("Error deleting " + file.getAbsolutePath());
        }
      }
    });

    return file;
  }
}
