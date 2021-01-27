package controllers

import (
	"encoding/json"
	"fmt"
	"ulivideoapi/services/es"
)

type EsDemoController struct {
	CommonController
}

// @router /es/add [*]
func (this *EsDemoController) Add() {
	body := map[string]interface{}{
		"id": 1,
		"title": "张三",
	}

	es.EsAdd("ulivideo_demo", "user_1", body)
	this.Ctx.WriteString("add")
}

// @router /es/edit [*]
func (this *EsDemoController) Edit() {
	body := map[string]interface{}{
		"id": 1,
		"title": "李四",
	}

	es.EsEdit("ulivideo_demo", "user_1", body)
	this.Ctx.WriteString("edit")
}

// @router /es/delete [*]
func (this *EsDemoController) Delete() {
	es.EsDelete("ulivideo_demo", "user_1")
	this.Ctx.WriteString("delete")
}

// @router /es/search [*]
func (this *EsDemoController) Search() {
	sort := []map[string]string{map[string]string{"id": "desc"}}

	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"id": 1,
				},
			},
		},
	}

	res := es.EsSearch("ulivideo_demo", query, 0, 10, sort)
	total := res.Total
	var resData []ResData
	for _, v := range res.Hits {
		var data ResData
		err := json.Unmarshal([]byte(v.Source), &data)
		if err != nil {
			fmt.Println(err)
		}
		resData = append(resData, data)
	}

	fmt.Println(total)
	fmt.Println(resData)
	this.Ctx.WriteString("search")
}

type ResData struct {
	Id int
	Title string
}

// @router /es/createIndex [*]
func (this *EsDemoController) CreateIndex() {
	body := map[string]interface{}{
		"settings": map[string]interface{}{
			"number_of_shards": 3,
			"number_of_replicas": 2,
		},
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type": "integer",
				},
				"title": map[string]interface{}{
					"type": "keyword",
					//"type": "text",
					//"analyzer": "ik_max_word",
					//"search_analyzer": "ik_max_word",
				},
				"sub_title": map[string]interface{}{
					"type": "keyword",
					//"type": "text",
					//"analyzer": "ik_max_word",
					//"search_analyzer": "ik_max_word",
				},
				"status": map[string]interface{}{
					"type": "integer",
				},
				"add_time": map[string]interface{}{
					"type": "integer",
				},
				"img": map[string]interface{}{
					"type": "keyword",
				},
				"img1": map[string]interface{}{
					"type": "keyword",
				},
				"channel_id": map[string]interface{}{
					"type": "integer",
				},
				"type_id": map[string]interface{}{
					"type": "integer",
				},
				"region_id": map[string]interface{}{
					"type": "integer",
				},
				"user_id": map[string]interface{}{
					"type": "integer",
				},
				"episodes_count": map[string]interface{}{
					"type": "integer",
				},
				"episodes_update_time": map[string]interface{}{
					"type": "integer",
				},
				"is_end": map[string]interface{}{
					"type": "integer",
				},
				"is_hot": map[string]interface{}{
					"type": "integer",
				},
				"is_recommend": map[string]interface{}{
					"type": "integer",
				},
				"comment": map[string]interface{}{
					"type": "integer",
				},
			},
		},
	}

	es.EsCreateIndex("ulivideo", body)
	this.Ctx.WriteString("EsCreateIndex")
}
