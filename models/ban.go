package models

import (
	"strconv"
	"time"

	"github.com/pasarsakomandiri/shared/database"
)

type Ban struct {
	Id           int64
	Ip_address   string
	Attempt      int
	Ban_time     string
	Created_date string
	Created_by   int
}

func BanFlushAttempt(ip string) error {
	_, err := database.Db.Exec("UPDATE bans SET attempt=?, last_update_date=? WHERE ip_address=?", 0, time.Now().String(), ip)
	return err
}

func BanGetInfoByHost(ip string) (Ban, error) {
	ban := Ban{}
	err := database.Db.Get(&ban, "SELECT id, ip_address, attempt, ban_time, created_date, created_by FROM bans WHERE ip_address=?", ip)
	return ban, err
}

func BanCreateNewAddress(ip string) error {
	_, err := database.Db.Exec("INSERT INTO bans(ip_address, attempt, ban_time, created_date, created_by) VALUES(?, ?, ?, ?, ?)", ip, 1, time.Now().String(), time.Now().String(), -1)
	return err
}

func BanUpdateAddress(ban Ban) error {
	_, err := database.Db.Exec("UPDATE bans SET attempt=?, ban_time=?, last_update_date=?, updated_by=? WHERE id="+strconv.FormatInt(ban.Id, 16), ban.Attempt, ban.Ban_time, time.Now().String(), -1)
	return err
}
