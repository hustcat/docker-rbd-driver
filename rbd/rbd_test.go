// +build linux

package rbd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/daemon/graphdriver"
	"os"
	"testing"
)

func init() {
	graphdriver.Register("rbd", Init)
}

func Init(home string, options []string) (graphdriver.Driver, error) {
	if err := os.MkdirAll(home, 0700); err != nil && !os.IsExist(err) {
		log.Errorf("Rbd create home dir %s failed: %v", err)
		return nil, err
	}

	rbdSet, err := NewRbdSet(home, true, options)
	if err != nil {
		return nil, err
	}

	if err := graphdriver.MakePrivate(home); err != nil {
		return nil, err
	}

	d := &RbdDriver{
		RbdSet: rbdSet,
		home:   home,
	}

	return graphdriver.NaiveDiffDriver(d), nil
}

// This avoids creating a new driver for each test if all tests are run
// Make sure to put new tests between TestRbdSetup and TestRbdTeardown
func TestRbdSetup(t *testing.T) {
	graphtest.GetDriver(t, "rbd")
}

func TestRbdCreateEmpty(t *testing.T) {
	graphtest.DriverTestCreateEmpty(t, "rbd")
}

func TestRbdCreateBase(t *testing.T) {
	graphtest.DriverTestCreateBase(t, "rbd")
}

func TestRbdCreateSnap(t *testing.T) {
	graphtest.DriverTestCreateSnap(t, "rbd")
}

func TestRbdTeardown(t *testing.T) {
	graphtest.PutDriver(t)
}
