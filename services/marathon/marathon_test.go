package marathon

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAcl(t *testing.T) {
	Convey("#acl", t, func() {

		s := service.Service{Acl: "hdr(host) -i foo.bar.com"}

		Convey("should return internal correctly", func() {

			label := "internal"
			h := conf.HAProxy{HostnameLabel: &label}
			labelJson := `{
				"internal": "foo.internal.bar.com"
			}`

			Convey("should return balancer type acl if hostname label for balancer type", func() {
				a := App{Labels: map[string]string{hostnameConfiguration: labelJson}}
				So(Acl(h, a, s), ShouldEqual, conf.AclFormat("foo.internal.bar.com"))

			})
		})

		Convey("should return service acl correctly", func() {
			label := ""
			ha := conf.HAProxy{HostnameLabel: &label}
			a := App{}
			Convey("should return service acl if no balancer type", func() {
				So(Acl(ha, a, s), ShouldEqual, s.Acl)
			})

			Convey("should return service acl if balancer type and no balancer rtype rule ", func() {
				label := "internal"
				ha = conf.HAProxy{HostnameLabel: &label}
				So(Acl(ha, a, s), ShouldEqual, s.Acl)
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
