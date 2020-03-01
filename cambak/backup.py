import copy as cp

from os import path, mkdir
from glob import glob
from shutil import copy

from .cameras import Camera
from .scanner import scan_folder


def backup_camera(
    src: str,
    dest: str,
    name: str,
    type_: Camera,
    force: bool = False,
    **kwargs
):
    def get_fullpaths(paths):
        existing_files = []

        for p in paths:
            fullpaths = glob(path.join(src, p))

            for fullpath in fullpaths:
                if path.isdir(fullpath):
                    existing_files.append(fullpath)
                else:
                    print("%s not exists. Skipping it.")

        return existing_files

    files = []
    dates = set()
    src_path = {}

    # Generate folder pictures fullpaths
    src_path["Pictures"] = get_fullpaths(type_.img_folders)
    src_path["RAW"] = get_fullpaths(type_.raw_folders)
    src_path["Videos"] = get_fullpaths(type_.vid_folders)

    # Check if the dest folder exists
    if path.isdir(dest) is False:
        raise IOError("Destination folder doesn't exists.")

    # Get all files categorized
    for folder_type, folders in src_path.items():
        if folder_type == "Pictures":
            extensions = type_.img_extensions
        elif folder_type == "RAW":
            extensions = type_.raw_extensions
        elif folder_type == "Videos":
            extensions = type_.vid_extensions

        for p in folders:
            files.extend(scan_folder(p, extensions, folder_type))

    # Get all dates and create all folders
    for _, date, _ in files:
        if date not in dates:
            dates.add(date)
            date_path = path.join(dest, date)

            if not path.exists(date_path):
                mkdir(date_path)
                mkdir(path.join(date_path, name))
                mkdir(path.join(date_path, name, "Pictures"))
                mkdir(path.join(date_path, name, "RAW"))
                mkdir(path.join(date_path, name, "Videos"))

    total_files = len(files)
    remaining_files = cp.copy(total_files)
    print("Total of %s files to copy." % total_files)

    # Copy all files in their correct folders
    for f, date, f_type in files:
        remaining_files -= 1
        dest_path = path.join(dest, date, name, f_type)
        basename = path.basename(path.join(dest_path, f))

        if (force is False and not path.exists(basename)) or (force is True):
            try:
                print(
                    "Copying %s - Remaining: %s/%s"
                    % (f, remaining_files, total_files)
                )
                copy(f, dest_path)
            except Exception as err:
                print(
                    "!!! Unable to copy %r file. Cause: \n\t%s" % (f, str(err))
                )
