# loaded config


class Config:
    project_name = ""
    version = ""
    username = ""
    password = ""
    download_path = {   ".zip":"",
                        ".tar.gz":""
                    }

class TestConfig(Config):
    DEBUG = True
