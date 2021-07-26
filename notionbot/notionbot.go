package notionbot

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"notion-integration-bot/config"
	"notion-integration-bot/model"
	"time"

	"github.com/tidwall/gjson"
)


func CheckDepartment(url string) (isChange bool, detailSet []model.DepartmentDetail){
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request.Header.Set("Notion-Version", "2021-05-13")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer " + config.NotionSecretKey)

	client := &http.Client{}
    response, err := client.Do(request)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
        log.Fatal(err)
    }
    jsonFormBody := string(bodyBytes)

	isChange = false

	for _, page := range gjson.Get(jsonFormBody, `results`).Array() {
		rawCreatedTime := gjson.Get(page.String(), `created_time`)
		createdTime , err := time.Parse(time.RFC3339, rawCreatedTime.Str)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		rawEditTime := gjson.Get(page.String(), `last_edited_time`)
		editTime, err := time.Parse(time.RFC3339, rawEditTime.Str)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		loc, _ := time.LoadLocation("UTC")
		current := time.Now().In(loc)
		currentBefore := current.Add(-10 * time.Second)

		//justify time problem
		createdTime = createdTime.Add(40 * time.Second)
		editTime = editTime.Add(40 * time.Second)
		
		fmt.Println("before")

		if(createdTime.Before(current) && createdTime.After(currentBefore)){
			isChange = true

			assigneeSet := gjson.Get(page.String(), `properties.Assignee.people`).Array()
			title := gjson.Get(page.String(), `properties.Projects.title.0.text.content`).Str
			taskType := gjson.Get(page.String(), `properties.Type.select.name`).Str
			status := gjson.Get(page.String(), `properties.Type.select.name`).Str
			priority := gjson.Get(page.String(), `properties.Priority.select.name`).Str

			for _, assignee := range assigneeSet{

				detail := model.DepartmentDetail{
					Action: "Create",
					Title: title,
					AssigneeEmail: gjson.Get(assignee.String(), `person.email`).Str,
					FieldSet: make([]model.Field, 0),
				}

				detail.FieldSet = append(detail.FieldSet, model.Field{"TaskType", taskType})
				detail.FieldSet = append(detail.FieldSet, model.Field{"Status", status})
				detail.FieldSet = append(detail.FieldSet, model.Field{"Priority", priority})
				
				detailSet = append(detailSet, detail)
			}

		} else if (editTime.Before((current)) && editTime.After(currentBefore)){
			isChange = true

			assigneeSet := gjson.Get(page.String(), `properties.Assignee.people`).Array()
			title := gjson.Get(page.String(), `properties.Projects.title.0.text.content`).Str
			taskType := gjson.Get(page.String(), `properties.Type.select.name`).Str
			status := gjson.Get(page.String(), `properties.Status.select.name`).Str
			priority := gjson.Get(page.String(), `properties.Priority.select.name`).Str
			pageLink := gjson.Get(page.String(), `url`).Str

			for _, assignee := range assigneeSet{

				detail := model.DepartmentDetail{
					Action: "Update",
					Title: title,
					AssigneeEmail: gjson.Get(assignee.String(), `person.email`).Str,
					FieldSet: make([]model.Field, 0),
				}

				detail.FieldSet = append(detail.FieldSet, model.Field{"TaskType", taskType})
				detail.FieldSet = append(detail.FieldSet, model.Field{"Status", status})
				detail.FieldSet = append(detail.FieldSet, model.Field{"Priority", priority})
				detail.FieldSet = append(detail.FieldSet, model.Field{"PageLink", pageLink})
				
				detailSet = append(detailSet, detail)
			}

		}
	}
	return
}