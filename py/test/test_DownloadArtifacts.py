import os
import sys
import json
path = os.path.dirname(os.path.abspath(__file__))
sys.path.append(path+'/../confluent/build/download_artifacts')

import unittest.mock as mock
from unittest.mock import patch, mock_open
from download_artifacts import DownloadArtifacts

test_pom = []
with open(path+'/data/pom.xml') as pom:
    for line in pom:
        test_pom.append(line)

test_maven_settings = []
with open(path+'/data/settings.xml') as settings:
    for line in settings:
        test_maven_settings.append(line)

nexus_result = json.loads(open(path+'/data/nexus.json').read())

@mock.patch('builtins.open', new_callable=mock_open, create=True)
def test_get_project_name_version(mock_open):
    mock_open.return_value.__enter__ = mock_open
    mock_open.return_value.__iter__ = mock.Mock(
        return_value = iter(test_pom))
    da = DownloadArtifacts()
    da.get_project_name_version()
    assert da.cg.project_name == "common"
    assert da.cg.version == "5.3.0"

@mock.patch('builtins.open', new_callable=mock_open, create=True)
def test_get_maven_settings(mock_open):
    mock_open.return_value.__enter__ = mock_open
    mock_open.return_value.__iter__ = mock.Mock(
        return_value = iter(test_maven_settings))
    da = DownloadArtifacts()
    da.get_maven_settings()
    assert da.cg.username == "xiang username"
    assert da.cg.password == "xiang password"

def mocked_requests_get(*args, **kwargs):
    class MockResponse:
        def __init__(self, json_data, status_code):
            self.json_data = json_data
            self.status_code = status_code

        def json(self):
            return self.json_data

    return MockResponse(nexus_result, 200)

@mock.patch('download_artifacts.requests.get', side_effect = mocked_requests_get)
def test_get_artifact_uri(mock_open):
    da = DownloadArtifacts()
    da.get_artifact_uri()
    assert da.cg.artifacts == {
                                'zip':{
                                    'file_ext':'.zip',
                                    'file_url':'https://nexus.confluent.io/repository/maven-snapshots/io/confluent/common-package/5.3.0-SNAPSHOT/common-package-5.3.0-20190413.162837-58-package.zip'
                                },
                                'tar.gz':{
                                    'file_ext':'.tar.gz',
                                    'file_url':'https://nexus.confluent.io/repository/maven-snapshots/io/confluent/common-package/5.3.0-SNAPSHOT/common-package-5.3.0-20190413.162837-58-package.tar.gz'
                                }
                            }