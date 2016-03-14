package configuration

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBalancerType(t *testing.T) {

	Convey("should find balancer types correctly", t, func() {

		Convey("When the label exists", func() {
			label := "internal"
			h := HAProxy{HostnameLabel: &label}

			So(*h.HostnameLabel, ShouldEqual, "internal")
			So(string(h.BalancerType()), ShouldEqual, "internal")
		})

		Convey("When the label is nil", func() {
			var label *string
			label = nil
			h := HAProxy{HostnameLabel: label}

			So(h.HostnameLabel, ShouldEqual, nil)
			So(string(h.BalancerType()), ShouldEqual, EmptyBalancerType)
		})
	})
}
