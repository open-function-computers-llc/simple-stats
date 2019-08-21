package main

import "testing"

func TestCanParseDFCommand(t *testing.T) {
	raw := `
Filesystem      Size  Used Avail Use% Mounted on
dev             7.8G     0  7.8G   0% /dev
run             7.8G  1.2M  7.8G   1% /run
/dev/sdb1       229G  116G  102G  54% /
tmpfs           7.8G  287M  7.6G   4% /dev/shm
tmpfs           7.8G     0  7.8G   0% /sys/fs/cgroup
tmpfs           7.8G   16M  7.8G   1% /tmp
/dev/sda1       916G  500G  370G  58% /home
tmpfs           1.6G   52K  1.6G   1% /run/user/1000
`
	discs := processDF(raw)
	if discs[0].MountPoint != "/" && discs[0].Size != "229G" { // could check other params, but whatever
		t.Error("df output not parsed correctly")
	}
}
