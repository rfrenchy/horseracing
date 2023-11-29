package domain

type Race struct {
        Id int
        Name string
        Date string // date of race
        Region string // region of race
        Course Course
        Offtime string // race off time
        Racetype string // type of racing (flat/hurdle/chase etc)
        Raceclass string // race class
        Pattern string // race pattern
        Ratingband string // rating restrictions
        Agebandrestriction string // age restrictions
        Sexrestriction string // sex restrictions
        Distance string // in metres
        Going string // going description
        Surface string // surface turf/dirt/aw
        Ran int // number of runners in race
}

type Course struct {
        Id int
        Name string
        Region int
}

type Region struct {
        Id int
        Name string
}

type Runner struct {
        Horse int
        Race int
        Racecardnumber int
        Position int // finished position
        Draw int // stall number
        Overbeaten int // total number of lengths beaten
        Beaten int // lengths behind nearest horse in front 
        Age string // age of horse at race time
        Weight string // weight in pounds
        Headgear string 
        Time string // time in minutes/seconds
        Odds string // decimal odds 
        Jockeyid int
        Trainerid int
        Prizemoney int 
        Officialrating int
        RPRating int
        TSrating int
        Ownerid int
        Comment string
}

type Horse struct {
        Id int
        Name string
        Sireid int
        Damsireid int
        Sex string
}

type Jockey struct {
        Id int
        Name string
}

type Trainer struct {
        Id int
        Name string
}

type Owner struct {
        Id int
        Name string
        Silkurl string
}
