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

func TestCanSkipAnySnapMountPoints(t *testing.T) {
	raw := `
/dev/vda1                                       54081164  15998308  35182952  32% /
/dev/loop0                                         63360     63360         0 100% /var/lib/snapd/snap/core20/1081
/dev/loop1                                           384       384         0 100% /var/lib/snapd/snap/bpytop/249
/dev/loop2                                         43392     43392         0 100% /var/lib/snapd/snap/certbot/1280
/dev/loop5                                        101760    101760         0 100% /var/lib/snapd/snap/core/11316
/dev/loop3                                         63232     63232         0 100% /var/lib/snapd/snap/core20/1026
/dev/loop4                                         33152     33152         0 100% /var/lib/snapd/snap/snapd/12398
/backups                                       114039976 101336060   6867992  94% /backups
/dev/loop6                                        101888    101888         0 100% /var/lib/snapd/snap/core/11420
/dev/loop7                                         33152     33152         0 100% /var/lib/snapd/snap/snapd/12704
/dev/loop8                                         43520     43520         0 100% /var/lib/snapd/snap/certbot/1343
`
	discs := processDF(raw)
	for _, disc := range discs {
		if len(disc.MountPoint) < 14 {
			continue
		}
		if disc.MountPoint[0:14] == "/var/lib/snapd" {
			t.Errorf("Disc should have been skipped. Mountpoint is: " + disc.MountPoint + " " + disc.MountPoint[0:14])
		}
	}
}
