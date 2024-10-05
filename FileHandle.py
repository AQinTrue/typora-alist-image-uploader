import os


class FileHandle:
    def __init__(self):
        self.base_path = os.path.join(os.path.dirname(__file__), "Typora_Config")
        self.create_basedir()

    def create_basedir(self):
        if not os.path.exists(self.base_path):
            os.makedirs(self.base_path)
