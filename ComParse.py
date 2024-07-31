import argparse


class ComParse:
    class Types:
        UpLoad = 'upload'
        Debug = 'debug'
        UpConfig = 'uconfig'

    def __init__(self):
        self.parser = argparse.ArgumentParser(description='Parse command line arguments')

        # 添加类型参数，使用choices限制可选值
        self.parser.add_argument('--type', type=str, choices=[self.Types.UpLoad,
                                                              self.Types.Debug, self.Types.UpConfig], help='运行状态')

        # 当type为UpLoad时，收集所有后续参数
        self.parser.add_argument('--upload_args', nargs=argparse.REMAINDER, help='参数用于上传操作')

        # 其他参数，先默认为可选
        self.parser.add_argument('--url', type=str, help='Alist网盘地址')
        self.parser.add_argument('--path', type=str, help='Alist存放路径')
        self.parser.add_argument('--token', type=str, help='管理员Token')

        # 解析命令行参数
        args = self.parser.parse_args()

        # 将参数存储为类的属性
        self.type = args.type
        self.url = args.url
        self.path = args.path
        self.token = args.token
        self.upload_args = args.upload_args


# 使用示例
if __name__ == "__main__":
    parser = ComParse()
    print("Type:", parser.type)
    if parser.type == 'upload':
        print("Upload Args:", parser.upload_args)
    elif parser.type == 'uconfig':
        print("URL:", parser.url)
        print("Path:", parser.path)
        print("Token:", parser.token)
