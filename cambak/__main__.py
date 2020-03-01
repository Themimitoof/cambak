import argparse
import signal

from cambak.backup import backup_camera
from cambak.cameras import get_all_cameras


def keyboardInterruptHandler(signal, frame):
    print("KeyboardInterrupt has been caught. Cleaning up...")
    exit(125)


def main_cli():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "src", help="Source folder (mounted card/usb camera volume)", type=str
    )

    parser.add_argument(
        "dest", help="Destination folder (local, network volume)", type=str
    )

    parser.add_argument(
        "-t", "--type", help="Type of camera", type=str, required=True
    )
    parser.add_argument(
        "-n", "--name", help="Name of the camera", type=str, required=True
    )

    parser.add_argument(
        "-f",
        "--force",
        action="store_true",
        help="Override if file already exists in the dest folder",
    )

    args = parser.parse_args()
    cameras_types = get_all_cameras()

    signal.signal(signal.SIGINT, keyboardInterruptHandler)
    backup_camera(
        args.src,
        args.dest,
        args.name,
        cameras_types.get(args.type, None),
        args.force,
    )


if __name__ == "__main__":
    main_cli()
