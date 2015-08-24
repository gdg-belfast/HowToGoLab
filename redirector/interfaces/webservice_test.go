package interfaces

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"errors"
	"net/http"
	"net/http/httptest"
)

func TestWebService(t *testing.T) {

	Convey(`Given I have a WebService`, t, func() {

		stubMapper := &stubHostMapper{}

		ws := &WebService{
			MappingInteractor: stubMapper,
		}

		Convey(`Given I have a Valid URL`, func() {

			receivedHostname := ""

			stubMapper.GetHostMappingFunc = func(host string) (string, error) {
				receivedHostname = host
				return "http://google.com", nil
			}

			Convey(`Given the hostname does not have a port`, func() {

				req, err := http.NewRequest("GET", "http://example.com/foo", nil)
				if err != nil {
					panic(err)
				}

				Convey(`When I go to the URL`, func() {

					recorder := httptest.NewRecorder()
					ws.Redirect(recorder, req)

					Convey(`Then I will be redirected`, func() {
						So(recorder.Code, ShouldEqual, http.StatusMovedPermanently)
					})

					Convey(`Then the original hostname was "example.com"`, func() {
						So(receivedHostname, ShouldEqual, "example.com")
					})
				})
			})

			Convey(`Given the hostname does have a port`, func() {

				req, err := http.NewRequest("GET", "http://example.com:9000/foo", nil)
				if err != nil {
					panic(err)
				}

				Convey(`When I go to the URL`, func() {

					recorder := httptest.NewRecorder()
					ws.Redirect(recorder, req)

					Convey(`Then the original hostname was "example.com"`, func() {
						So(receivedHostname, ShouldEqual, "example.com")
					})
				})
			})
		})
		Convey(`Given I have an Invalid URL`, func() {

			stubMapper.GetHostMappingFunc = func(host string) (string, error) {
				return "", errors.New("Invalid URL")
			}

			Convey(`When I go to the URL`, func() {

				req, err := http.NewRequest("GET", "http://unknownaddress.com", nil)
				if err != nil {
					panic(err)
				}

				recorder := httptest.NewRecorder()
				ws.Redirect(recorder, req)

				Convey(`Then I will get an Internal Server Error`, func() {
					So(recorder.Code, ShouldEqual, http.StatusBadRequest)
				})
			})

		})

	})

}

type stubHostMapper struct {
	GetHostMappingFunc func(hostname string) (string, error)
}

func (stub *stubHostMapper) GetHostMapping(hostname string) (string, error) {
	return stub.GetHostMappingFunc(hostname)
}
