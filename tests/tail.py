import os
from datetime import datetime
from uuid import uuid4
import time
import json


def write_data_tofile():
    with open('D:\\home\\audit_modsec.log', 'a') as f:
        for i in range(10):
            f.write(json.dumps({"datetime": str(datetime.now()), "uid": str(uuid4()), "index": i}))
            f.write('\n')
    f.close()


def test():
    while True:
        print('---------write--------')
        write_data_tofile()
        time.sleep(1)


if __name__ == '__main__':
    test()

