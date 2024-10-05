import os


class FileHandle:
    def __init__(self):
        self.base_directory = os.getcwd()

    def create_directory(self):
        if not os.path.exists(self.base_directory):
            os.makedirs(self.base_directory)
