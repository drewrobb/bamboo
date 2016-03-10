package configuration

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBalancerType(t *testing.T) {

	Convey("should find balancer types correctly", t, func() {

		Convey("When the label is internal", func() {
			label := "internal"
			h := HAProxy{HostnameLabel: &label}

			So(*h.HostnameLabel, ShouldEqual, "internal")
			So(h.BalancerType(), ShouldEqual, InternalBalancerType)
		})

		Convey("When the label is external", func() {
			label := "external"
			h := HAProxy{HostnameLabel: &label}

			So(*h.HostnameLabel, ShouldEqual, "external")
			So(h.BalancerType(), ShouldEqual, ExternalBalancerType)
		})

		Convey("When the label is nil", func() {
			var label *string
			label = nil
			h := HAProxy{HostnameLabel: label}

			So(h.HostnameLabel, ShouldEqual, nil)
			So(h.BalancerType(), ShouldEqual, EmptyBalancerType)
		})
	})
}
