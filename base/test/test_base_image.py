import os
import unittest

import confluent.docker_utils as utils


class BaseImageTest(unittest.TestCase):

    def setUp(self):
        self.image = "{0}confluentinc/cp-base:{1}".format(os.environ["DOCKER_REGISTRY"], os.environ["DOCKER_TAG"])

    def test_image_build(self):
        self.assertTrue(utils.image_exists(self.image))

    def test_java_install(self):
        cmd = "java -version"
        expected = 'OpenJDK Runtime Environment (Zulu 8.'
        output = utils.run_docker_command(image=self.image, command=cmd)
        self.assertTrue(expected in output)

    def test_dub_exists(self):
        self.assertTrue(utils.path_exists_in_image(self.image, "/usr/local/bin/dub"))
        self.assertTrue(utils.path_exists_in_image(self.image, "/usr/local/bin/cub"))


if __name__ == '__main__':
    unittest.main()
