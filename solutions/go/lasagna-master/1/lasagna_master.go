package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int)(int) {
    if time == 0{
        time = 2
    }
    return len(layers)*time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string)(noodleNeeded int, sauceNeeded float64) {
    l := len(layers)
    for i := 0; i<l; i++ {
        if layers[i] == "noodles" {
            noodleNeeded += 50
        }
        if layers[i] == "sauce" {
            sauceNeeded += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList = append(myList[:len(myList)-1], friendList[len(friendList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
    var it float64 = float64(portions)/2
    for _, val := range quantities {
        scaledQuantities = append(scaledQuantities, val * it)
    }
    return
    
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
