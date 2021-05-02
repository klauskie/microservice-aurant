package util

import (
	"github.com/robfig/cron/v3"
	"klauskie.com/microservice-aurant/party-service/repository"
	"log"
)

func ClearCacheCron() {
	log.Println("Creating ClearCacheCron...")
	c := cron.New()
	_, err := c.AddFunc("@midnight", func() {
		log.Println("Running cache clear on party map...")
		repository.GetPartyRepository().ClearAll()
	} )
	if err != nil {
		log.Println("clearCacheCron: " + err.Error())
	}
	c.Start()
}
