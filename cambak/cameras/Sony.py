from . import Camera


class SonyNex(Camera):
    """General support for Sony NEX cameras (Alpha 5, 6, 7 and 9)"""

    img_folders = ["DCIM/*MSDCF"]
    raw_folders = img_folders
    vid_folders = ["PRIVATE/M4ROOT/CLIP"]

    img_extensions = [".JPG"]
    raw_extensions = [".ARW"]
    vid_extensions = [".MP4"]


class SonyAlpha(SonyNex):
    """Shortcut for Sony Alpha cameras"""


class SonyHDR(SonyNex):
    """
    Support for Sony "HDR" Actions cams. Similar to the Sony NEX architecture.
    """
