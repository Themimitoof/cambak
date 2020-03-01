class Camera:
    """Abstract class of Camera."""

    img_folders = None
    raw_folders = None
    vid_folders = None

    img_extensions = None
    raw_extensions = None
    vid_extensions = None


def get_all_cameras():
    import pkgutil, importlib, inspect

    cameras = {}

    for _, modname, _ in pkgutil.walk_packages(
        path=__path__, prefix=__name__ + ".", onerror=lambda x: None
    ):
        modulesource = importlib.import_module(modname)
        for module in inspect.getmembers(modulesource, inspect.isclass):
            class_name = module[0]

            if not cameras.get(class_name):
                cameras[class_name] = module[1]

    # Pop the 'Camera' abstract class
    cameras.pop("Camera", None)

    return cameras
