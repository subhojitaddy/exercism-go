package partyrobot

import(
    f "fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return f.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return f.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	returnString := Welcome(name) + "\n" + "You have been assigned to table "
    if table < 100 {
        if table < 10 {
            returnString = returnString + f.Sprintf("00%d", table)
        } else {
        	returnString = returnString + f.Sprintf("0%d", table)
        }
    } else {
        returnString = returnString + f.Sprintf("%d", table)
    }

    returnString = returnString + f.Sprintf(". Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", direction, distance, neighbor)
    
    return returnString
}
