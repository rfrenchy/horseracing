package main

import (
	"os"
	"testing"
	// "github.com/gocarina/gocsv"
)

var csv = `date,region,course,off,race_name,type,class,pattern,rating_band,age_band,sex_rest,dist,dist_f,dist_m,going,ran,num,pos,draw,ovr_btn,btn,horse,age,sex,lbs,hg,time,secs,dec,jockey,trainer,prize,or,rpr,sire,dam,damsire,owner,comment
2008-04-30,GB,Ascot,2:10,Ascot Annual Badgeholders Conditions Stakes,Flat,Class 3,,,2yo,,5f,5f,1006,Good To Soft,7,1,1,1,0,0,Baycat (IRE),2,G,127,,1:4.85,64.85,3.75,James Doyle,Jonathan Portman,6231,–,92,One Cool Cat (USA),Greta DArgent (IRE),Great Commotion,A S B Portman,Held up behind leaders - hung left and led over 1f out - kept on well - ridden out(op 9-2)`

func TestRacingPostRecord(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Error("unable to create file for testing")
	}
	defer os.Remove(f.Name())

	if _, err := f.Write([]byte(csv)); err != nil {
		t.Error("unable to write to file for testing")
	}

	x := &Course{}
	x.Name = "yes"

	// records := []*RacingPostRecord{}
	// if err := gocsv.UnmarshalFile(f, &records); err != nil {
	// 	t.Error("unable to unmarshall csv to racingpostrecord")
	// }

	// if len(records) < 1 {
	//         t.Error("incorrect total of records unmarshalled")
	// }
}

func TestOwner(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Error("unable to create file for testing")
	}
	defer os.Remove(f.Name())

	if _, err := f.Write([]byte(csv)); err != nil {
		t.Error("unable to write to file for testing")
	}

}
