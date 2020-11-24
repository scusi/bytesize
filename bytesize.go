// Package ByteSize provides human readable byte sizes.
//
// Basically after typecast your size into int64 and ByteSize you can print
// the size as human readable string.
//
package bytesize

//
// Usage:
//
// import("ByteSize")
//
// file, err := os.Open()
// check(err)
// defer file.Close()
// fileInfo, err := file.Stat()
// check(err)
// fileSize := bytesize.ByteSize(int64(fileInfo.Size()))
// fmt.Printf("Filesize is: %s\n", fileSize)
//

import (
	"fmt"
)

// ByteSize is a datatype to have human readable byte sizes.
type ByteSize float64

const (
	_ = iota // ignore first value by assigning to blank identifier
	// KB defines a KiloByte
	KB ByteSize = 1 << (10 * iota)
	// MB defines a MegaByte
	MB
	// GB defines a GigaByte
	GB
	// TB defines a TerraByte
	TB
	// PB defines a PetaByte
	PB
	// EB defines an ExoByte
	EB
	// ZB defines a ZetaByte
	ZB
	// YB defines a YotaByte
	YB
)

// (b ByteSize) String() is the Stringer function for ByteSize.
// It returns a human readable string for the given size of bytes
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2f YB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2f ZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2f EB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2f PB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2f KB", b/KB)
	}
	return fmt.Sprintf("%.2f B", b)
}
