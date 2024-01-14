package urlPgxRepository

const (
	SaveShortAndFullQuery  = `insert into url (original_url, short_url) values ($1, $2) on conflict do nothing`
	GetFullUrlByShortQuery = `select original_url from url where short_url = $1`
	CleanUpOldRecords      = `delete from url where creation_time < now() - interval '10 seconds'`
)
