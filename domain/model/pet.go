package model

type TinyUrl struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}

func (TinyUrl) TableName() string { return "url" }
