package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type CandidateController struct {
	beego.Controller
}

func (c *CandidateController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "candidate.tpl"
}

func setCandidateInfoFromControl(c *CandidateController) models.Candidate {
	var candidate models.Candidate

	id, err := c.GetInt64("candidateid")
	if err != nil {
		beego.Debug("failed to get candidate id : ", err.Error())
	} else {
		candidate.Id = id
	}

	candidate.Fullname = c.GetString("fullname")

	candidate.Age = c.GetString("age")

	candidate.Gender = c.GetString("gender")

	candidate.Mobile = c.GetString("mobile")

	candidate.Email = c.GetString("email")

	candidate.Workyear = c.GetString("workyear")

	candidate.Post = c.GetString("post")

	candidate.City = c.GetString("city")

	candidate.Company = c.GetString("company")

	candidate.Education = c.GetString("education")

	birth := c.GetString("birthday")
	candidate.Birthday, _ = time.Parse("2006-01-02", birth)

	return candidate

}

func (c *CandidateController) InsertOneCandidate() {
	candidate := setCandidateInfoFromControl(c)
	err := candidate.Insert()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	} else {
		c.Ctx.WriteString(strconv.FormatInt(candidate.Id, 10))
	}

}

func (c *CandidateController) UpdateOneCandidate() {
	candidate := setCandidateInfoFromControl(c)
	err := candidate.Update()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	} else {
		c.Ctx.WriteString("ok")
	}

}

func (c *CandidateController) GetCandidates() {

	var candidate models.Candidate
	ps := candidate.GetCandidates()
	// get candidate labels
	for i := range ps {
		var cand models.Candidate
		cand.Id = ps[i][0].(int64)
		ps[i] = append(ps[i], cand.GetCandidateLabels())

	}
	c.Data["json"] = &ps
	c.ServeJSON()
}

func (c *CandidateController) GetCandidateByMD5() {
	md5 := c.GetString("md5")
	var candidate models.Candidate
	candidate.Md5 = md5
	candidate.GetCandidateByMD5()
	c.Data["json"] = &candidate
	c.ServeJSON()
}

func (c *CandidateController) GetCandidateByID() {

	var candidate models.Candidate
	candidate.Id, _ = c.GetInt64("id")
	candidate.GetCandidateByID()
	c.Data["json"] = &candidate
	c.ServeJSON()
}

func (c *CandidateController) GetCandidateByLabels() {

	s := c.GetString("labelids")
	//	log.Println(s, len(s))

	labels := strings.Split(s, ",")
	//	log.Println(labels, len(labels))

	var ids []int
	for _, v := range labels {
		id, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}

	//	log.Println(ids, len(ids))

	var candidate models.Candidate
	if len(ids) > 0 {
		ps := models.GetCandidatesByLabels(ids)
		c.Data["json"] = &ps
	} else {
		ps := candidate.GetCandidates()
		c.Data["json"] = &ps
	}

	c.ServeJSON()

}
