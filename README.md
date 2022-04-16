# Cambak

![GitHub](https://img.shields.io/github/license/themimitoof/cambak) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/themimitoof/cambak) [![Maintainability](https://api.codeclimate.com/v1/badges/7ab0233767b6c405fdf2/maintainability)](https://codeclimate.com/github/Themimitoof/cambak/maintainability)

Cambak is a small but powerful tool that simplify and automate the derushing process of your cameras.

Cambak is flexible and customizable. You can choose what kind of medias (pictures, RAW files, movies) you want to copy, the structure of the destination folder, the name of the camera you used, and how you want to handle copy conflicts (skip or replace/merge). If you always use the same settings, you can put them into a configuration and call cambak in an instant!


## Note for users using the Python version

Great news folks! After two years, Cambak have been fully rewritten in Go. Welcome to a single binary file, increased performances and bye-bye cameras architectures and annoying bugs! Give place to this new version, easier to use and more flexible!

We hope you will love this new version. Please be aware that the default structure of the destination folder have changed. Instead of `year-month-day/camera-name/{Pictures,RAW,Videos}`, this new version uses the structure `year/month-day/camera-name/{Pictures,RAW,Videos}`. If you want to keep the old structure, use the format `%y/%m-%d/%n/%t` in the configuration file or with the flag `--format`.

If you search the source code of the Python version, it is still available via the [old branch](https://github.com/Themimitoof/cambak/tree/old) or via the [v0.1.x](https://github.com/Themimitoof/cambak/tags) tags.

## Features

 - Choose to extract Pictures, RAW files, movies or all three type of files
 - Extract to a configurable output folder
 - Skip or merge files that already exists in the destination folder
 - Delete source files after copy
 - (Run in Dry-run mode to see before executing the real command.)


## Planned for the future

 - Read EXIF information for:
   - implementing a date range filter
   - implementing a rating range filter (for example to only copy the pictures you rated 4 or 5 stars)
 - Improve the way to read the creation date of each file
 - Create a daemon for automating the derush process when a SD Card is mounted on the system
 - (Probably one day, probably never, for now is a simple idea, nothing more 👀) Add a GUI to visualize, sort and prepare an extraction process. (this idea is more for fun of learning something new, there are better softwares that does that task very well.)

## Installation

### GitHub release binaries

All releases are available [here](https://github.com/Themimitoof/cambak/releases) and contains compiled versions for:

 - Linux, MacOS, FreeBSD, OpenBSD running on the `amd64` architecture
 - Linux, MacOS, FreeBSD, OpenBSD running on the `arm64` architecture
 - Linux, FreeBSD, OpenBSD running on the `armhf` architecture
 - Debian, RPM and APK (Alpine) packages are available for `amd64`, `arm64` and `armhf`.

**Note:** Until Cambak have not been tested on Windows, no binaries will be provided.

### Manual install

You can install Cambak by using the command `go install` if you have Go installed on your machine:

```bash
go install github.com/themimitoof/cambak@latest
```

You can also clone the repository and build the project if you want the latest version in development:

```bash
git clone https://github.com/themimitoof/cambak.git
cd cambak
go build
./cambak --help
```


## Usage

Cambak is very easy to use. To get the list of commands available, type `cambak help`:

```bash
 → cambak help
Cambak is a simple but powerful too for derushing cameras.

The program use a configuration file located in '$HOME/.config/cambak.yaml'.
During the first execution, a default configuration file will be created. You
can override it by an another configuration file by using the --config flag.

For more information, please consult: https://github.com/themimitoof/cambak.

Usage:
  cambak [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  extract     Copy files from a source media to a local/remote destination
  help        Help about any command
  version     Return the version of Cambak

Flags:
      --config string   Path of the configuration file (default "/home/themimitoof/.config/cambak.yaml")
  -h, --help            help for cambak

Use "cambak [command] --help" for more information about a command.
```

To see how to use the _extraction_ method, type `cambak extract --help`:

```bash
 → cambak extract --help
The cambak extrator will copy/extract files from a source media (eg:
SD card, MTP drive, local/remote folder) to a local or remote destination folder.

By default, the folder destination structure is the following:

<destination folder>
└── <YEAR>
    └── <MONTH>-<DAY>
        └── <CAMERA_NAME>
            ├── Pictures
            ├── RAW
            └── Movies

You can change the destination format by using the '--format' flag or change the value
in the configuration file.

For more information, please consult: https://github.com/themimitoof/cambak.

Usage:
  cambak extract [flags]

Aliases:
  extract, copy, cp

Flags:
  -A, --all             Import all medias files type
  -c, --clean           Delete source file after been copied
      --dry-run         Only log what the extractor will do if this flag was not set
  -f, --format string   Structure format in the destination folder.
  -h, --help            help for extract
  -m, --merge           Merge the source file if it already exists in the destination folder
  -M, --movies          Import movies files
  -n, --name string     Name of the camera
  -P, --pictures        Import pictures files
  -R, --raws            Import RAWs files
  -s, --skip            Skip the source file if it already exists in the destination folder

Global Flags:
      --config string   Path of the configuration file (default "/home/themimitoof/.config/cambak.yaml")
```

**Note:** the command extract can also be invoked with the alias `copy` or `cp`.

To start your first job, plug a SD Card on your computer (or use a local folder that contains images, RAW files and/or movies from a Digital Camera) and run the below command:

```bash
→ cambak extract -n MyCameraName /mnt/my-sdcard ~/backup-camera
Generating the default configuration file...
✅ 30 files collected (20 pictures, 8 RAWs, 2 movies). 0 files skipped.
Copying files...
 100% |██████████████████████████████████████████████| (30/30, 71 it/s)
✨ All files have been copied!
```

Now, if you take a look on the destination folder, here's your files sorted by year, date, camera name and by type of files:

```bash
→ tree ~/backup-camera
/home/themimitoof/backup-camera
└── 2022
    ├── 03-15
    │   └── MyCameraName
    │       ├── Movies
    │       │   └── 20191231_203102.mp4
    │       ├── Pictures
    │       │   ├── DSC00150.JPG
    │       │   ├── DSC00150.smsmsm.JPG
    │       │   ├── DSC00151.jpg
    │       │   └── DSC00152.JpG
    │       └── RAW
    │           ├── DSC00150.ARW
    │           └── DSC00151.arw
    └── 03-16
        └── MyCameraName
            ├── Movies
            │   └── 20191231_203102.mp4
            ├── Pictures
            │   ├── DSC00150.JPG
            │   ├── DSC00150.smsmsm.JPG
            │   ├── DSC00151.jpg
            │   └── DSC00152.JpG
            └── RAW
                ├── DSC00150.ARW
                └── DSC00151.arw
```

Houray! You can now try to update the [configuration file](#configuration) (generated during the first execution) for having the same behavior each time you use Cambak. You can always override the settings on the fly by using the different options available (check `cambak extract --help` for more information).

## Configuration

By default, the configuration file is located into the folder `$HOME/.config/cambak.yaml` and use the markup language `YAML`.

The default configuration look like this:

```yaml
extract:
    pictures: true
    raws: true
    movies: true
    destination: ""
    format: '%y/%m-%d/%n/%t'
    conflict: skip
    camera_name: Camera
    clean_after_copy: false
```

| Configuration key | Type | Description |
|-------------------|------|-------------|
| `extract` |
| `pictures` | `boolean` | Specify if we want by default to extract pictures |
| `raws` | `boolean` | Specify if we want by default to extract RAW pictures files |
| `movies` | `boolean` | Specify if we want by default to extract movies |
| `destination` | `string (path)` | Specify the default destination path |
| `format` | `string` | Specify the folder structure in the destination folder. See _[Customize the destination structure](#customize-the-destination-structure)_ for more information |
| `conflict` | `string` | Specify how to handle conflict files if it already exist in the destination folder. Two values are available : `skip` and `merge` (default: `skip`). |
| `camera_name` | `string` | Specify the default name of the Camera |
| `clean_after_copy` | `boolean` | Specify if Cambak need to delete the file from the source media/folder after being copied to the destination folder |

Here's my personal configuration as an example:

```yaml
extract:
    pictures: true
    raws: true
    movies: true
    destination: "/mnt/okinawa-smb/Medias/backup-brut"
    format: '%y/%m-%d/%n/%t'
    conflict: skip
    camera_name: A6400
    clean_after_copy: false
```

It will :
 - copy all pictures, RAW files and movies
 - copy all files to my home server via a Samba share folder
 - skip all files if they already exist on my home server
 - the name of my camera is (Sony) `A6400`
 - do not clean the SD Card after copying the files to my home server (I format my SD Card via my camera few weeks later just for safety)

With that configuration, I just need to execute this command each time I need to extract all my pictures after a session:

```bash
→ cambak cp /mnt/my-sdcard
✅ 30 files collected (20 pictures, 8 RAWs, 2 movies). 0 files skipped.
Copying files...
 100% |██████████████████████████████████████████████| (30/30, 71 it/s)
✨ All files have been copied!
```

That's it! I just need to specify the source folder and press enter, [_simple, basique_](https://www.youtube.com/watch?v=2bjk26RwjyU).

### Customize the destination structure

The format of the destination structure is very flexible and gives you the possibility to structure your media files like you want.

Here's the list of available verbs:

| Verb | Description |
|------|-------------|
| `%y` | Year (e.g.: 2022) |
| `%m` | Month (e.g.: `04`) |
| `%d` | Day (e.g.: `01`) |
| `%n` | Camera name |
| `%t` | Media type (Pictures, RAW, Movies) |

For example, if you want to have one folder with the complete date, you don't want to have a folder with the name of the camera name but a folder for each type of medias, the format will be: `%y-%m-%d/%t`.

If you want something even more customized (for example, to give a name to a shot session/trip/etc.), you can use something like this: `Stockholm Trip - %y-%m/%t`.

In case you want more flexibility and a verb is missing, feel free to [open an issue](https://github.com/Themimitoof/cambak/issues) (if it not already exists) or edit the _[files/format.go](files/format.go)_ file and send a [Pull request](https://github.com/Themimitoof/cambak/pulls).

## Supported files

Cambak support a large set of file extensions (the base of this list is based from what ImageMagick supports):

| Extension | Type |
|-----------|------|
| `HEIF` | Pictures |
| `JPEG` | Pictures |
| `JPG` | Pictures |
| `TIFF` | Pictures |
| `DCRAW` | RAW files |
| `DNG` | RAW files |
| `IIQ` | RAW files |
| `RAW` | RAW files |
| `RMF` | RAW files |
| `CR2` | RAW files |
| `CR3` | RAW files |
| `CRW` | RAW files |
| `ERF` | RAW files |
| `RAF` | RAW files |
| `3FR` | RAW files |
| `DCR` | RAW files |
| `K25` | RAW files |
| `KDC` | RAW files |
| `MEF` | RAW files |
| `NEF` | RAW files |
| `NRW` | RAW files |
| `ORF` | RAW files |
| `RW2` | RAW files |
| `PEF` | RAW files |
| `X3F` | RAW files |
| `ARW` | RAW files |
| `MRW` | RAW files |
| `SR2` | RAW files |
| `SRF` | RAW files |
| `MP4` | Movies |
| `MOV` | Movies |

In case an extension file is missing, feel free to [open an issue](https://github.com/Themimitoof/cambak/issues) (if it not already exists) or edit the _[files/file_extensions.go](files/file_extensions.go)_ file and send a [Pull request](https://github.com/Themimitoof/cambak/pulls).


## Contributions

In case you have a suggestion or want a new feature, feel free to open a [discussion](https://github.com/Themimitoof/cambak/discussions).

If you found a bug, you can [open an issue](https://github.com/Themimitoof/cambak/issues).

In case you are able to fix by yourself a bug, enhance the code or implement a new feature, feel free to send a [Pull request](https://github.com/Themimitoof/cambak/pulls).


## License

This project is under the [MIT license](LICENSE). Feel free to use, contribute, fork and do what you want with it. Please keep all licenses, copyright notices and mentions in case you use, re-use, steal, fork code from this repository.
