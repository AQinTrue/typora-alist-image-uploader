import uuid
import hashlib
import base64
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad


class DataCrypto:
    def __init__(self):
        self.key = str(uuid.getnode()).encode('utf-8')

    def create_aes_iv(self):
        hash_object = hashlib.sha256(self.key)
        return hash_object.digest()[-16:]

    def create_aes_key(self):
        hash_object = hashlib.sha256(self.key)
        return hash_object.digest()[:32]

    def encrypt(self, message):
        cipher = AES.new(self.create_aes_key(), AES.MODE_CBC, self.create_aes_iv())
        message = pad(message.encode('utf-8'), AES.block_size)
        ct_bytes = cipher.encrypt(message)
        ct_bytes = base64.b64encode(ct_bytes)
        return ct_bytes

    def decrypt(self, ct):
        ct_bytes = base64.b64decode(ct)
        cipher = AES.new(self.create_aes_key(), AES.MODE_CBC, self.create_aes_iv())
        pt = unpad(cipher.decrypt(ct_bytes), AES.block_size)
        return pt.decode('utf-8')


if __name__ == '__main__':
    data = DataCrypto()
    print(data.create_aes_key())
    en_data = data.encrypt("Hello World")
    print(en_data)
    de_data = data.decrypt(en_data)
    print(de_data)
