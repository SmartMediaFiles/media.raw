package media_raw

import (
	"github.com/smartmediafiles/media/media/maps"
	"github.com/smartmediafiles/media/media/types"
)

// List of supported media.Raw file types.
const (
	ImageRaw types.FileType = "raw" // RAW Image
	ImageDNG types.FileType = "dng" // Adobe Digital Negative
)

// RawFileTypesExtensions is a map of media.Raw file types to their file extensions.
var RawFileTypesExtensions = maps.MapFileTypeExtensions{
	ImageDNG: {".dng"},
	ImageRaw: {
		".3fr",
		".ari", ".arw",
		".bay",
		".cap", ".crw", ".cr2", ".cr3",
		".data", ".dcs", ".dcr", ".drf",
		".eip", ".erf",
		".fff",
		".gpr",
		".iiq",
		".k25", ".kdc",
		".mdc", ".mef", ".mos", ".mrw",
		".nef", ".nrw",
		".obm", ".orf",
		".pef", ".ptx", ".pxn",
		".r3d", ".raf", ".raw", ".rwl", ".rwz", ".rw2",
		".srf", ".srw", ".sr2",
		".x3f",
	},
}
