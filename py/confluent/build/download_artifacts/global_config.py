class DownloadArtifactsConfig:
    project_name = ""
    version = ""
    username = ""
    password = ""
    artifacts = {
                    "zip":{
                        "file_ext":".zip",
                        "file_url":""
                    },
                     "tar.gz":{
                         "file_ext":".tar.gz",
                         "file_url":""
                     }
                    }
    DEBUG = False