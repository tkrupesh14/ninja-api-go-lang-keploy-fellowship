// To run
// go run .
package main
// To import all packages at once
// go get .
import (
    "net/http"
    "github.com/gin-gonic/gin" //go get -u github.com/gin-gonic/gin
)
// superhero represents the type of data we have about a superhero.
// json:"data" is known as struct tags. Struct tags allow us to attach meta-information to corresponding struct properties. In other words, we use them to reformat the JSON response returned by the API.
type ninjahero struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Power   int64  `json:"power"`
    Special string `json:"special"`
}
// superheroes slice to initialize a few superheroes in our variable. Not using any database.
// Slice is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array.
var ninjaheros = []ninjahero{
    {ID: "1", Name: "Ninja-Man", Power: 2000, Special: "Flight, NinjaMan Strength, X-Ray Vision, Super Speed"},
    {ID: "2", Name: "Ball-Man", Power: 500, Special: "Rich"},
    {ID: "3", Name: "Fellowship-Woman", Power: 1600, Special: "NinjaMan Strength, Speed"},
    {ID: "4", Name: "Keploy-Man", Power: 1500, Special: "Flight, Weapons, Armor"},
    {ID: "5", Name: "Ninjar-Man", Power: 1100, Special: "Web, Armor, Cling to walls"},
}
func home(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{ // H is a shortcut for map[string]interface{}
        "instructions": "Add '/ninjas' to the link",
    })
}
func getNinjas(c *gin.Context) {
    // Printing all the superheroes available in the data
    c.JSON(http.StatusOK, ninjaheros)
}
func addNinjas(c *gin.Context) {
    // Creating a new object to structure superhero
    var newNinjahero ninjahero
// Call BindJSON to bind the received JSON to newSuperhero
    // BindJSON adds the data provided by user to newSuperhero
    // This is kind of "try catch" concept
    if err := c.ShouldBindJSON(&newNinjahero); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   true,
            "message": "Bad Request",
        })
        return
    }
// Add the new superhero to the slice.
    ninjaheros = append(ninjaheros, newNinjahero)
    
    // Serializing the struct as JSON and adding it to the response
    c.JSON(http.StatusCreated, newNinjahero)
}
func editNinjas(c *gin.Context) {
    id := c.Param("id")
    
    // Creating a new object to structure superhero
    var editNinjas ninjahero
// Call BindJSON to bind the received JSON to newSuperhero
    // BindJSON adds the data provided by user to newSuperhero
    // This is kind of "try catch" concept
    if err := c.ShouldBindJSON(&editNinjas); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   true,
            "message": "Bad Request",
        })
        return
    }
    
    for i, hero := range ninjaheros {
        if hero.ID == id {
            ninjaheros[i].Name = editNinjas.Name
            ninjaheros[i].Power = editNinjas.Power
            ninjaheros[i].Special = editNinjas.Special
c.JSON(http.StatusOK, editNinjas)
            return
        }
    }
// If the above statement doesn't return anything, that means the id is invalid
    c.JSON(http.StatusBadRequest, gin.H{
        "error":   true,
        "message": "Invalid",
    })
}
func removeNinjas(c *gin.Context) {
    id := c.Param("id")
    for i, hero := range ninjaheros {
        if hero.ID == id {
            // arr := [100, 200, 300, 400, 500]
            // arr[:2] = [100, 200]
            // arr[2:] = [300, 400, 500]
            // arr[2+1:] = [400. 500]
            // [100, 200][400, 500]
            ninjaheros = append(ninjaheros[:i], ninjaheros[i+1:]...) // ... is required when writing 2 slices in append function
c.JSON(http.StatusOK, gin.H{
                "message": "Item Deleted",
            })
            return
        }
    }
// If the above statement doesn't return anything, that means the id is invalid
    c.JSON(http.StatusBadRequest, gin.H{
        "error":   true,
        "message": "Invalid",
    })
}
func main() {
    // Creating a gin router with default middleware: logger and recovery (crash-free) middleware
    router := gin.Default()
    router.GET("/", home)
    router.GET("/ninjas", getNinjas)
    router.POST("/ninjas", addNinjas)
    router.PUT("/ninjas/:id", editNinjas)
    router.DELETE("/ninjas/:id", removeNinjas)
router.Run() // By default it serves on :8080 unless a PORT environment variable was defined.
    // router.Run(":3000") for customized PORT
}