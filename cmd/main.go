package main

import "github.com/tiohlognm/pgback/internal"

func main() {
	internal.BackupData("/dev/null")
}
