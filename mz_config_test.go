package util

import "testing"


func Test_MzConfig(t *testing.T) {
    config, err := ReadMzConfig("./config.test.ini")
    if err != nil {
        t.Error(err)
    }

    val := config.Get("main.foo", "invalid")
    if val != "bar" {
        t.Error("Get failed to find correct value")
    }

    val = config.Get("part1.foo", "invalid")
    if val != "blah" {
        t.Error("Get failed to find correct value")
    }

    val = config.Get("main.unknown", "known")
    if val != "known" {
        t.Error("Get failed to return default")
    }

    bval := config.GetFlag("main.truth")
    if ! bval {
        t.Error("Flag incorrectly detected")
    }

    bval = config.GetFlag("main.gorp")
    if ! bval {
        t.Error("Flag incorrectly detected")
    }
}

