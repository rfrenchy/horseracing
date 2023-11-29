package models

import "github.com/gocarina/gocsv"

type ScrapeRecord struct {
        /** Race Info */
        RaceDate string  `csv:"date"`
        RaceRegion string  `csv:"region"`
        CourseId string `csv:"course_id"`
        Course string `csv:"course"`

        RaceId string `csv:"race_id"`
        RaceOfftime string `csv:"off"` 
        RaceName string `csv:"race_name"`

        Racetype string `csv:"type"` 
        Raceclass string  `csv:"class"`
        RacePattern string `csv:"pattern"` // (Group 1 etc)

        RatingbandRestrictions string `csv:"rating_band"`
        AgebandRestriction string `csv:"age_band"`
        SexRestriction string `csv:"sex_rest"`

        Distance string `csv:"dist_m"` // in metres

        Going string `csv:"going"` // going description
        Surface string `csv:"surface"` // surface Turf/Dirt/AW

        /** Runner Info */
        RacecardNumber string `csv:"num"`
        FinishedPosition string `csv:"position"` // runner finished position
        Draw string `csv:"draw"` // stall number

        Overbeaten string `csv:"ovr_btn"` // total number of lengths beaten
        Beaten string `csv:"btn"` // lengths behind nearest horse in front 

        HorseId string `csv:"horse_id"`
        HorseName string `csv:"horse"` 
        HorseAge string `csv:"age"` 
        HorseSex string `csv:"sex"` 

        HorseWeight string `csv:"lbs"` // weight in pounds

        Headgear string `csv:"hg"`

        FinishTime string `csv:"time"` // time taken in minutes/seconds

        DecimalOdds string `csv:"dec"`

        JockeyId string `csv:"jockey_id"`
        Jockey string `csv:"jockey"`
        TrainerId string `csv:"trainer_id"`
        Trainer string `csv:"trainer"`

        PrizeMoney string `csv:"prize"`

        OfficialRating string `csv:"or"`
        RPRRating string `csv:"rpr"`
        TSRating string `csv:"ts"`

        SireId string `csv:"sire_id"`
        SireName string `csv:"sire"`
        DamId string `csv:"dam_id"`
        DamName string `csb:"dam"`
        DamsireId string `csv:"damsire_id"`
        DamsireName string `csv:"damsire"`

        OwnerId string `csv:"owner_id"`
        OwnerName string `csv:"owner"`
        SilkURL string `csv:"silk_url"` // URL of silk colours
        
        Comment string `csv:"comment"` // Form in running comments
}

