import os

from os import path
from datetime import datetime


def scan_folder(src, extensions, type_):
    files = []
    exts = []

    # Ensure all extensions are in lower case
    for ext in extensions:
        exts.append(ext.lower())

    for f in os.listdir(src):
        fullpath = path.join(src, f)
        extension = path.splitext(f)[-1].lower()

        if path.isdir(fullpath):
            files.extend(scan_folder())

        if extension in exts:
            created_at = datetime.fromtimestamp(
                os.stat(fullpath).st_ctime
            ).strftime("%Y-%m-%d")

            files.append((fullpath, created_at, type_))

    return files
