package dao

import (
	"testing"
)

func init() {

}

func TestDevice(t *testing.T) {
	e := new(DeviceEntity)
	e.Sys = "a"
	e.SysVersion = "a"
	e.PhoneModel = "a"
	e.MacAddr = "a"
	e.PackageVersion = "a"
	e.SimStatus = "a"
	e.Insert()
}
