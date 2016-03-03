package models
import (
	"github.com/jmoiron/sqlx"
	"time"
	"strconv"
	"github.com/pasarsakomandiri/shared/database"
)

type Ban struct {
	Id int64
	Ip_address string
	Attempt int
	Ban_time string
	Created_date string
	Created_by int
}

func BanGetInfoByHost(db *sqlx.DB, ip string) (Ban, error){
	ban:= Ban{}
	err := database.Db.Get(&ban, "SELECT id, ip_address, attempt, ban_time, created_date, created_by FROM bans WHERE ip_address=?", ip)
	return ban, err
}

func BanCreateNewAddress(db *sqlx.DB, ip string) error {
	_, err := database.Db.Exec("INSERT INTO bans(ip_address, attempt, ban_time, created_date, created_by) VALUES(?, ?, ?, ?, ?)", ip, 1, time.Now().String(), time.Now().String(), -1)
	return err
}

func BanUpdateAddress(db *sqlx.DB, ban Ban) error {
	_, err := database.Db.Exec("UPDATE bans SET attempt=?, ban_time=?, last_update_date=?, updated_by=? WHERE id="+strconv.FormatInt(ban.Id, 16), ban.Attempt, ban.Ban_time, time.Now().String(), -1)
	return err
}