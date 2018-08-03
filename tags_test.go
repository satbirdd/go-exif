package exif

import (
	"testing"

	"github.com/dsoprea/go-logging"
)

func TestGet(t *testing.T) {
	im := NewIfdMappingWithStandard()
	ti := NewTagIndex(im)

	it, err := ti.Get(TiffIfdPathStandard, 0x10f)
	log.PanicIf(err)

	if it.Is(TiffIfdPathStandard, 0x10f) == false || it.IsName(TiffIfdPathStandard, "Make") == false {
		t.Fatalf("tag info not correct")
	}
}

func TestGetWithName(t *testing.T) {
	im := NewIfdMappingWithStandard()
	ti := NewTagIndex(im)

	it, err := ti.GetWithName(TiffIfdPathStandard, "Make")
	log.PanicIf(err)

	if it.Is(TiffIfdPathStandard, 0x10f) == false || it.Is(TiffIfdPathStandard, 0x10f) == false {
		t.Fatalf("tag info not correct")
	}
}
