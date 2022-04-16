/**
Contains a list of files formats used for the collection of all files.

The RAW list came from an output of ImageMagick with the command: `convert -list Format | grep DNG`
*/

package files

var picturesExtensions = []string{
	"heif",
	"jpeg",
	"jpg",
	"tiff",
}

var rawExtensions = []string{
	// Generic RAW formats
	"dcraw", // Raw Photo Decoder (dcraw)
	"dng",   // Digital Negative
	"iiq",   // Phase One Raw Image Format
	"raw",   // Raw
	"rmf",   // Raw Media Format

	// Canon RAW formats
	"cr2", // Canon Digital Camera Raw Image Format
	"cr3", // Canon Digital Camera Raw Image Format
	"crw", // Canon Digital Camera Raw Image Format

	// Epson RAW formats
	"erf", // Epson RAW Format

	// Fujifilm RAW formats
	"raf", // Fuji CCD-RAW Graphic File

	// Hasselblad RAW formats
	"3fr", // Hasselblad CFV/H3D39II

	// Kodak RAW formats
	"dcr", // Kodak Digital Camera Raw Image File
	"k25", // Kodak Digital Camera Raw Image Format
	"kdc", // Kodak Digital Camera Raw Image Format

	// Miyama RAW formats
	"mef", // Mamiya Raw Image File

	// Nikon RAW formats
	"nef", // Nikon Digital SLR Camera Raw Image File
	"nrw", // Nikon Digital SLR Camera Raw Image File

	// Olympus RAW formats
	"orf", // Olympus Digital Camera Raw Image File

	// Panasonic RAW formats
	"rw2", // Panasonic Lumix Raw Image

	// Pentax RAW formats
	"pef", // Pentax Electronic File

	// Sigma RAW formats
	"x3f", // Sigma Camera RAW Picture File

	// Sony RAW formats
	"arw", // Sony Alpha Raw Image Format
	"mrw", // Sony (Minolta) Raw Image File
	"sr2", // Sony Raw Format 2
	"srf", // Sony Raw Format
}

var moviesExtensions = []string{
	"mp4",
	"mov",
}
