package horseracing

type Race struct {
	Id                 int
	Name               string
	Date               string // date of race
	Region             string // region of race
	Course             Course
	Offtime            string // race off time
	Racetype           string // type of racing (flat/hurdle/chase etc)
	Raceclass          string // race class
	Pattern            string // race pattern
	Ratingband         string // rating restrictions
	Agebandrestriction string // age restrictions
	Sexrestriction     string // sex restrictions
	Distance           string // in metres
	Going              string // going description
	Surface            string // surface turf/dirt/aw
	Ran                int    // number of runners in race
}

type Course struct {
	Id     int
	Name   string
	Region int
}

type Region struct {
	Id   int
	Name string
}

type Runner struct {
	Horse          int
	Race           int
	Racecardnumber int
	Position       int    // finished position
	Draw           int    // stall number
	Overbeaten     int    // total number of lengths beaten
	Beaten         int    // lengths behind nearest horse in front
	Age            string // age of horse at race time
	Weight         string // weight in pounds
	Headgear       string
	Time           string // time in minutes/seconds
	Odds           string // decimal odds
	Jockeyid       int
	Trainerid      int
	Prizemoney     int
	Officialrating int
	RPRating       int
	TSrating       int
	Ownerid        int
	Comment        string
}

type Horse struct {
	ID        int
	Name      string
	SireID    int // Father
	DamID     int // Mother
	DamsireID int // Father of the Dam
	Sex       string
}

type Jockey struct {
	Id   int
	Name string
}

type Trainer struct {
	Id   int
	Name string
}

type Owner struct {
	Id      int
	Name    string
	Silkurl string
}

type RacingPostRecord struct {
	// Race information
	RaceDate               string `csv:"date"`
	RaceRegion             string `csv:"region"`
	CourseID               string `csv:"course_id"`
	Course                 string `csv:"course"`
	RaceID                 string `csv:"race_id"`
	RaceOfftime            string `csv:"off"`
	RaceName               string `csv:"race_name"`
	RaceType               string `csv:"type"`
	RaceClass              string `csv:"class"`
	RacePattern            string `csv:"pattern"` // (Group 1 etc)
	RatingbandRestrictions string `csv:"rating_band"`
	AgebandRestriction     string `csv:"age_band"`
	SexRestriction         string `csv:"sex_rest"`
	Distance               string `csv:"dist_m"`  // in metres
	Going                  string `csv:"going"`   // going description
	Surface                string `csv:"surface"` // surface Turf/Dirt/AW
	Ran                    string `csv:"ran"`     // number of runners in race

	// Specific Runner information
	RacecardNumber   string `csv:"num"`
	FinishedPosition string `csv:"pos"`     // runner finished position
	Draw             string `csv:"draw"`    // stall number
	Overbeaten       string `csv:"ovr_btn"` // total number of lengths beaten
	Beaten           string `csv:"btn"`     // lengths behind nearest horse in front
	HorseID          string `csv:"horse_id"`
	HorseName        string `csv:"horse"`
	HorseAge         string `csv:"age"`
	HorseSex         string `csv:"sex"`
	HorseWeight      string `csv:"lbs"` // weight in pounds
	Headgear         string `csv:"hg"`
	FinishTime       string `csv:"time"` // time taken in minutes/seconds
	DecimalOdds      string `csv:"dec"`
	JockeyID         string `csv:"jockey_id"`
	JockeyName       string `csv:"jockey"`
	TrainerID        string `csv:"trainer_id"`
	TrainerName      string `csv:"trainer"`
	PrizeMoney       string `csv:"prize"`
	OfficialRating   string `csv:"or"`
	RPRRating        string `csv:"rpr"`
	TSRating         string `csv:"ts"`
	SireID           string `csv:"sire_id"`
	SireName         string `csv:"sire"`
	DamID            string `csv:"dam_id"`
	DamName          string `csb:"dam"`
	DamsireID        string `csv:"damsire_id"`
	DamsireName      string `csv:"damsire"`
	OwnerID          string `csv:"owner_id"`
	OwnerName        string `csv:"owner"`
	SilkURL          string `csv:"silk_url"` // URL of silk colours
	Comment          string `csv:"comment"`  // Form in running comments
}
