package usecases

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/gdg-belfast/HowToGoLab/redirector/domain"

	"errors"
)

func TestHostMapInteractor(t *testing.T) {

	Convey(`Given a HostMapInteractor`, t, func() {

		stubRepo := &stubHostMapRepository{}

		mapper := &HostMapper{
			Repo: stubRepo,
		}

		Convey(`Given I have a URL that exists`, func() {

			mapping := &domain.HostMap{
				From: "example.org",
				To:   "example.com",
			}
			readFuncCalled := false

			stubRepo.ReadFunc = func(host string) (*domain.HostMap, error) {
				readFuncCalled = true
				return mapping, nil
			}

			Convey(`When I call the GetHostMapping method`, func() {

				url, err := mapper.GetHostMapping(mapping.From)

				Convey(`Then I will call the Read method of the HostMapRepository`, func() {
					So(readFuncCalled, ShouldBeTrue)
				})

				Convey(`Then I will get the URL to redirect to`, func() {
					So(url, ShouldEqual, mapping.To)
				})

				Convey(`Then I will get a nil error`, func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey(`Given accessing the Repo raises an error`, func() {
			stubRepo.ReadFunc = func(host string) (*domain.HostMap, error) {
				return nil, errors.New("Bad Repo")
			}

			Convey(`When I call the GetHostMapping method`, func() {

				url, err := mapper.GetHostMapping("insignificant.org")

				Convey(`Then the URL will be blank`, func() {
					So(url, ShouldEqual, "")
				})

				Convey(`Then I will get an error`, func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})

}

// stubHostMapRepository is a stub used for testing which can be used
// for mocking in tests
type stubHostMapRepository struct {
	CreateFunc func(*domain.HostMap) error
	ReadFunc   func(host string) (*domain.HostMap, error)
	UpdateFunc func(*domain.HostMap) error
	DeleteFunc func(*domain.HostMap) error
}

func (stub *stubHostMapRepository) Create(hm *domain.HostMap) error {
	return stub.CreateFunc(hm)
}

func (stub *stubHostMapRepository) Read(host string) (*domain.HostMap, error) {
	return stub.ReadFunc(host)
}

func (stub *stubHostMapRepository) Update(hm *domain.HostMap) error {
	return stub.UpdateFunc(hm)
}

func (stub *stubHostMapRepository) Delete(hm *domain.HostMap) error {
	return stub.DeleteFunc(hm)
}
