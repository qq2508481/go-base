package gcacheclient

import (
	"go-base/lib/gcache"
	"go-base/lib/util"
	"time"
)

func TestDriver() {
	// get by driver
	cache := gcache.D("default")
	util.Dump(cache.Set("test1", "test1"))
	util.Dump(cache.SetTTL("test2", "test2", 5*time.Second))
	util.Dump(cache.Get("test1"))
	util.Dump(cache.Del("test1"))
}

func TestRegion() {
	// get user region
	cache := gcache.R("user")
	util.Dump(cache.Set("user1", "user1"))
	util.Dump(cache.SetTTL("user2", "user2", 5*time.Second))
	util.Dump(cache.Get("user1"))
	util.Dump(cache.Del("user1"))
}
