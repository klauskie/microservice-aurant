package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"klauskie.com/microservice-aurant/party-service/models"
	"net/http"
)

func GetItemDefinitionMap(itemIdList []string) map[string]models.CatalogItem {
	url := URL_CATALOG_API + "/item-definition"
	jsonValue, _ := json.Marshal(itemIdList)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	resultMap := make(map[string]models.CatalogItem)
	if err := json.Unmarshal(body, &resultMap); err != nil {
		fmt.Println(err.Error())
	}

	return resultMap
}

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
