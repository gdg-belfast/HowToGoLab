package domain

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"encoding/json"
)

func TestHostMapping(t *testing.T) {

	Convey(`Given a HostMapping`, t, func() {

		mapping := &HostMap{
			From: "apple.com",
			To:   "google.com",
		}

		Convey(`When It is encoded to JSON`, func() {
			result, err := json.Marshal(mapping)

			Convey(`Then there will be no error in conversion`, func() {
				So(err, ShouldBeNil)
			})

			Convey(`Then it will produce valid JSON`, func() {
				So(string(result), ShouldEqual, `{"From":"apple.com","To":"google.com"}`)
			})
		})
	})

	Convey(`Given a JSON Encoded HostMapping`, t, func() {

		mapping := []byte(`{"From":"apple.com","To":"google.com"}`)

		Convey(`When It is unmarshalled from JSON`, func() {

			var result *HostMap

			err := json.Unmarshal(mapping, &result)

			Convey(`Then there will be no error in conversion`, func() {
				So(err, ShouldBeNil)
			})

			Convey(`Then the From Field will be set correctly`, func() {
				So(result.From, ShouldEqual, "apple.com")
			})

			Convey(`Then the To Field will be set correctly`, func() {
				So(result.To, ShouldEqual, "google.com")
			})
		})
	})
}
