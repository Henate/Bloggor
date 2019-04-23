package main

import (
	"time"
	"log"

	"github.com/robfig/cron"

	"github.com/Henate/Bloggor/models"
)

func main() {
	log.Println("Starting...")

	c := cron.New()	//根据本地时间创建一个新（空白）的 Cron job runner
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}