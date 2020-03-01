# CamBak
This small tool automate the derushing process of your SD Cards for all your
cameras.

CamBak sort all your files from your SD Card per date, camera and per type of
media and copy it to your computer or on a network volume. Here's is the
destination architecture:
```
/ *destination_folder*
├── *date_of_shots (example: 2020-02-13)*
│   └── *camera_name*
│       ├── Pictures
│       ├── RAW
│       └── Videos
```

## Supported cameras architectures
 * Sony (tested with NEX6 and HDR-AS series, a6400 and AS100V)

If a camera is not supported, you can create a new file on the `cameras`
folder with the brand name, add a class inherited of `Camera` and add each
paths and extensions for each type of medias (glog can be used for paths).

Here's is an example for Sony NEX cameras (file `Sony.py`):
```python
class SonyNex(Camera):
    """General support for Sony NEX cameras (Alpha 5, 6, 7 and 9)"""

    img_folders = ["DCIM/*MSDCF"]
    raw_folders = img_folders
    vid_folders = ["PRIVATE/M4ROOT/CLIP"]

    img_extensions = [".JPG"]
    raw_extensions = [".ARW"]
    vid_extensions = [".MP4"]
```

## Usage
```
➜ python cambak --help
usage: cambak [-h] -t TYPE -n NAME [-f] src dest

positional arguments:
  src                   Source folder (mounted card/usb camera volume)
  dest                  Destination folder (local, network volume)

optional arguments:
  -h, --help            show this help message and exit
  -t TYPE, --type TYPE  Type of camera
  -n NAME, --name NAME  Name of the camera
  -f, --force           Override if file already exists in the dest folder
```

Example:
```bash
cambak /mnt/sd-card /mnt/moon-smb/cam-backups -t SonyNex -n A6400
```

## Installation
WIP