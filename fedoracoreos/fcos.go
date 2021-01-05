package fedoracoreos

import (
	"fmt"
	"net/url"
)

const (
	// StreamStable is the default stream
	StreamStable = "stable"
	// StreamTesting is what is intended to land in stable
	StreamTesting = "testing"
	// StreamNext usually tracks the next Fedora major version
	StreamNext = "next"
)

const (
	buildHostURL           = "https://builds.coreos.fedoraproject.org/"
	releaseIndexLocation   = "prod/streams/%s/releases.json"
	streamMetadataLocation = "streams/%s.json"
)

// mustURL parses a URL we know is valid
func mustURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return *u
}

// GetStreamURL returns the URL for the given stream
func GetStreamURL(stream string) url.URL {
	return mustURL(buildHostURL + fmt.Sprintf(streamMetadataLocation, stream))
}

// GetReleaseIndexURL returns the URL for the release index of a given stream.
// Avoid this unless you have a specific need to test a specific release.
func GetReleaseIndexURL(stream string) url.URL {
	return mustURL(buildHostURL + fmt.Sprintf(releaseIndexLocation, stream))
}

// GetCosaBuild returns the coreos-assembler build URL
func GetCosaBuild(stream, buildID, arch string) url.URL {
	return mustURL(buildHostURL + fmt.Sprintf("prod/streams/%s/builds/%s/%s/", stream, buildID, arch))
}
