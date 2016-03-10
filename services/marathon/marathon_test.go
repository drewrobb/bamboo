package marathon

import (
	"testing"

	. "github.com/QubitProducts/bamboo/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	conf "github.com/QubitProducts/bamboo/configuration"
	"github.com/QubitProducts/bamboo/services/service"
)

func TestAcl(t *testing.T) {
	Convey("#acl", t, func() {

		s := service.Service{Acl: "foo.bar.com"}

		Convey("should return internal correctly", func() {

			label := "internal"
			h := conf.HAProxy{HostnameLabel: &label}

			Convey("should return internal if internal bamboo and internal hostname label", func() {
				a := App{Labels: map[string]string{InternalHostnameLabel: "foo.internal.bar.com"}}
				So(Acl(h, a, s), ShouldEqual, "foo.internal.bar.com")

			})
			Convey("should return default if internal bamboo and no internal hostname label", func() {
				a := App{}
				So(Acl(h, a, s), ShouldEqual, "foo.bar.com")
			})
		})

		Convey("should return external correctly", func() {
			label := "external"
			ha := conf.HAProxy{HostnameLabel: &label}

			Convey("should return external if external bamboo and external hostname label", func() {
				a := App{Labels: map[string]string{ExternalHostnameLabel: "foo.external.bar.com"}}
				So(Acl(ha, a, s), ShouldEqual, "foo.external.bar.com")
			})

			Convey("should return default if external bamboo and no internal hostname label", func() {
				a := App{}
				So(Acl(ha, a, s), ShouldEqual, "foo.bar.com")
			})
		})

		Convey("should return service acl correctly", func() {
			label := ""
			ha := conf.HAProxy{HostnameLabel: &label}
			a := App{}
			Convey("should return service acl if no balancer type", func() {
				So(Acl(ha, a, s), ShouldEqual, "foo.bar.com")
			})
			Convey("should return service acl if internal and no internal acl", func() {
				label := "internal"
				ha = conf.HAProxy{HostnameLabel: &label}
				So(Acl(ha, a, s), ShouldEqual, "foo.bar.com")
			})
			Convey("should return service acl if external and no external acl", func() {
				label := "external"
				ha = conf.HAProxy{HostnameLabel: &label}
				So(Acl(ha, a, s), ShouldEqual, "foo.bar.com")
			})
		})
	})
}

func TestParseHealthCheckPathTCP(t *testing.T) {
	Convey("#parseHealthCheckPath", t, func() {
		checks := []marathonHealthCheck{
			marathonHealthCheck{"/", "TCP", 0},
			marathonHealthCheck{"/foobar", "TCP", 0},
			marathonHealthCheck{"", "TCP", 0},
		}
		Convey("should return no path if all checks are TCP", func() {
			So(parseHealthCheckPath(checks), ShouldEqual, "")
		})
	})
}

func TestParseHealthCheckPathHTTP(t *testing.T) {
	Convey("#parseHealthCheckPath", t, func() {
		checks := []marathonHealthCheck{
			marathonHealthCheck{"/first", "HTTP", 0},
			marathonHealthCheck{"/", "HTTP", 0},
			marathonHealthCheck{"", "HTTP", 0},
		}
		Convey("should return the first path if all checks are HTTP", func() {
			So(parseHealthCheckPath(checks), ShouldEqual, "/first")
		})
	})
}

func TestParseHealthCheckPathMixed(t *testing.T) {
	Convey("#parseHealthCheckPath", t, func() {
		checks := []marathonHealthCheck{
			marathonHealthCheck{"", "TCP", 0},
			marathonHealthCheck{"/path", "HTTP", 0},
			marathonHealthCheck{"/", "HTTP", 0},
		}
		Convey("should return the first path if some checks are HTTP", func() {
			So(parseHealthCheckPath(checks), ShouldEqual, "/path")
		})
	})
}
