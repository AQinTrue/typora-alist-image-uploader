import os
import json
import requests
import time
from datetime import datetime
from DataCrypto import DataCrypto
from ComParse import ComParse
from FileHandle import FileHandle


class QinULoad:
    def __init__(self):
        self.cryptographer = DataCrypto()
        self.command_parser = ComParse()
        self.Config = "ULConfig.json"
        self.time_format = "%Y_%m_%d_%H_%M_%S"
        self.upload_message = "Upload Success:"
        self.file_handle = FileHandle()

    def initialize(self):
        try:
            with open(f"{self.file_handle.base_path}/{self.Config}", 'r') as config_file:
                encrypted_content = json.load(config_file)
                return (self.cryptographer.decrypt(encrypted_content['url']),
                        self.cryptographer.decrypt(encrypted_content['path']),
                        self.cryptographer.decrypt(encrypted_content['token']))
        except Exception as e:
            print(f"Initialization error: {e}")
            return

    def rename_file(self, file_path):
        try:
            file_extension = file_path.split('.')[-1]
            timestamp_millis = int(round(time.time() * 1000))
            dt_object = datetime.fromtimestamp(timestamp_millis / 1000.0)
            microseconds = int(dt_object.microsecond / 1000)
            formatted_time = dt_object.strftime(self.time_format) + '_' + f'{microseconds:03d}'
            return f"{formatted_time}.{file_extension}"
        except Exception as e:
            print(f"Renaming error: {e}")
            return

    def debug(self, file_path):
        file_name = self.rename_file(file_path)
        self.upload_message += f"\nhttps://test.org/{file_name}"

    def write_configuration(self):
        url = self.cryptographer.encrypt(self.command_parser.url)
        path = self.cryptographer.encrypt(self.command_parser.path)
        token = self.cryptographer.encrypt(self.command_parser.token)
        with open(f"{self.file_handle.base_path}/{self.Config}", 'w+') as f:
            content = {
                "url": url.decode(),
                "path": path.decode(),
                "token": token.decode(),
            }
            f.write(json.dumps(content))

    def request_file_upload(self, file_path, server_url, storage_path, token):
        try:
            file_name = self.rename_file(file_path)
            headers = {
                'Authorization': token,
                'Content-Type': 'application/octet-stream',
                'Content-Length': str(os.path.getsize(file_path)),
                'file-path': f'{storage_path}/{file_name}'
            }
            with open(file_path, 'rb') as file:
                response = requests.put(server_url + '/api/fs/put', headers=headers, data=file)
                if response.status_code == 200:
                    self.upload_message += f"\n{server_url}/d/{storage_path}/{file_name}"
                else:
                    self.upload_message += "\nError"
        except Exception as e:
            print(f"File upload error: {e}")

    def process(self):
        try:
            if self.command_parser.type == self.command_parser.Types.UpLoad:
                server_url, storage_path, token = self.initialize()
                if server_url and storage_path and token:
                    for file_path in self.command_parser.upload_args:
                        self.request_file_upload(file_path, server_url, storage_path, token)
            elif self.command_parser.type == self.command_parser.Types.Debug:
                for file_path in self.command_parser.upload_args:
                    self.debug(file_path)
            elif self.command_parser.type == self.command_parser.Types.UpConfig:
                self.write_configuration()
                print("Success!")
                return
            print(self.upload_message)
        except Exception as e:
            print(f"Processing error: {e}")


if __name__ == '__main__':
    QinULoad().process()
