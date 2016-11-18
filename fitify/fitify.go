package fitify

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "github.com/renstrom/fuzzysearch/fuzzy"
    "strings"
    //"net/url"
    //"google.golang.org/appengine"
    //"google.golang.org/appengine/urlfetch"
)
type exerciseJSON struct {
  Next string `json: next`
  Results []exercise `json: results`
}
type muscleJSON struct {
  Next string `json: next`
  Results []muscle `json: results`
}
type equipmentJSON struct {
  Next string `json: next`
  Results []equipment `json: results`
}
type imagesJSON struct {
  Next string `json: next`
  Results []exercise_img `json: results`
}
type exercise struct{
  Id int `json: id`
  Description string `json: description`
  Name string `json: name`
  Category int `json: category`
  Equipment []int `json: equipment`
}
 type muscle struct {
   Id int `json: id`
   Name string `json: name`
 }
 type equipment struct {
   Id int `json: id`
   Name string `json: name`
 }
 type exercise_img struct {
   Id int `json: id`
   Image string `json: image`
   Exercise int `json: exercise`
 }

 var exercises []exercise = make([]exercise,500,500)
 var muscles []muscle = make([]muscle,500,500)
 var equipments []equipment = make([]equipment,500,500)
 var exercises_img []exercise_img = make([]exercise_img,500,500)

 // case_1 := []string{"tell me exercises using"}
 // case_2 := []string{"tell me exercises"}
 // case_3 := []string{"what the hell?"}
 //

 var case_1  = []string{"tell me exercises that use",
   "tell me exercises using",
   "what are exercises that use",
   "how can i work out using",
   "how can i workout using",
   "how can i train using",
 }
 var case_2  = []string{"tell me exercises that train",
   "tell me exercises that workout",
   "tell me exercises that work out",
   "what are exercises that train",
   "what are exercises that workout",
   "what are exercises that work out",
   "how can i workout",
   "how can i work out",
 }
 var case_3  = []string{"tell me exercises that train using",
   "which exercises train using",
   "how can i train using",
   "how can i workout my using",
   "workouts using"}
 var case_5 = []string{"Show me images for"}
 var case_6 = []string{"Does use ", "Is needed for "}
 var case_7 = []string{"Does train"}

func Case1(message string) (string){
   input := []string{message}
   f := false
   for _,q := range case_1 {
     fuzzyMatches := fuzzy.Find(q, input)
     if len(fuzzyMatches) !=0 {
       f = true
     }
   }
   if !f {
     fmt.Println("fail no matches")
     return "err"
   }

      equipmentID := -1
    for _,eq := range equipments {
      if strings.Contains(message, eq.Name){
          equipmentID = eq.Id
        }
      }
    if equipmentID == -1{
        fmt.Println("fail equipment")
        return "err"
      }

      result := ""
    for _,ex := range exercises{
      if containsElement(ex.Equipment, equipmentID) {
        result += ex.Name + "\n" // + ex.Description
      }
    }
    if result == ""{
      return "err"
    }
    return result
 }

func Case2(message string) (string){
   input := []string{message}
   f := false
   for _,q := range case_2 {
     fuzzyMatches := fuzzy.Find(q, input)
     if len(fuzzyMatches) !=0 {
       f = true
     }
   }
   if !f {
     fmt.Println("fail no matches")
     return "err"
   }

       muscleID := -1
   for _,m := range muscles {
     if strings.Contains(message, m.Name){
       muscleID = m.Id
     }
   }
   if muscleID == -1{
     fmt.Println("fail muscle")
     return "err"
   }

      result := ""
    for _,ex := range exercises{
      if ex.Category == muscleID {
        result += ex.Name + "\n" // + ex.Description
      }
    }
    if result == ""{
      return "err"
    }
    return result
 }

func Case3(message string) (string){
   input := []string{message}
   f := false
   for _,q := range case_3 {
     fuzzyMatches := fuzzy.Find(q, input)
     if len(fuzzyMatches) !=0 {
       f = true
     }
   }
   if !f {
     fmt.Println("fail no matches")
     return "err"
   }
       muscleID := -1
   for _,m := range muscles {
     if strings.Contains(message, m.Name){
       muscleID = m.Id
     }
   }
   if muscleID == -1{
     fmt.Println("fail muscle")
     return "err"
   }
      equipmentID := -1
    for _,eq := range equipments {
      if strings.Contains(message, eq.Name){
          equipmentID = eq.Id
        }
      }
    if equipmentID == -1 {
        fmt.Println("fail equipment")
        return "err"
      }
      result := ""

    for _,ex := range exercises {
      if ex.Category == muscleID && containsElement(ex.Equipment, equipmentID) {
        result += ex.Name + "\n" // + ex.Description
        for _,img := range exercises_img {
          if img.Exercise == ex.Id {
            result += "Image: " + img.Image + "\n"
          }
        }
      }
    }


    if result == "" {
      return "err"
    }
    return result
 }


func Case5(message string) (string){
   input := []string{message}
   var result string
   f := false
   for _, q := range case_5 {
     fuzzyMatches := fuzzy.Find(q, input)
     if len(fuzzyMatches) != 0{
       f = true
     }
   }

   if !f {
     fmt.Println("flag is false")
     return "err"
   }
   exID := -1

   for _, ex := range exercises{
     if ex.Name == "" {
       continue
     }
     if strings.Contains(message, ex.Name){
       exID = ex.Id
       fmt.Println("found exercise" + ex.Name)
       break;
     }
   }

   if exID == -1 {
     fmt.Println("exID -1")
     return "err"
   }

   for _, img := range exercises_img{
     if img.Image == "" {
       continue
     }
     if img.Exercise == exID {
       fmt.Printf("%s image found %d" , img.Image, exID)
       result +=  img.Image + "\n"
     }
   }
   fmt.Println("didn't find image")
   if result == "" {
     return "err"
   }
   return result
  }

func Case6(message string) (string){
   input := []string{message}
   f := false

   for _, q := range case_6 {
     fuzzyMatches := fuzzy.Find(q, input)
     if len(fuzzyMatches) != 0{
       f = true
     }
   }

   if !f {
     fmt.Println("flag is false")
     return "err"
   }

   exID := -1
   var exer exercise
   for _, ex := range exercises{
     if ex.Name == "" {
       continue
     }
     if strings.Contains(message, ex.Name){
       exID = ex.Id
       exer = ex
       fmt.Println("found exercise" + ex.Name)
       break;
     }
   }

   if exID == -1 {
     fmt.Println("exID -1")
     return "err"
   }

   eqID := -1
   for _, eq := range equipments{
     if eq.Name == ""{
       continue
     }
     if strings.Contains(message, eq.Name){
       fmt.Println("equipment found")
       eqID = eq.Id
       break
     }

   }

   for _,eq := range exer.Equipment{
     fmt.Printf("is %d == %d ?\n" , eq , eqID)
     if eqID == eq {
       return "YES!"
     }
   }

   return "NO"
 }

func Case7(message string) (string){
   input := []string{message}
   f := false

   for _, q := range case_7 {
     fuzzyMatches := fuzzy.Find(q, input)
     if len(fuzzyMatches) != 0{
       f = true
     }
   }

   if !f {
     fmt.Println("flag is false")
     return "err"
   }

   exID := -1
   var exer exercise
   for _, ex := range exercises{
     if ex.Name == "" {
       continue
     }
     if strings.Contains(message, ex.Name){
       exID = ex.Id
       exer = ex
       fmt.Println("found exercise" + ex.Name)
       break;
     }
   }

   if exID == -1 {
     fmt.Println("exID -1")
     return "err"
   }

   mID := -1
   var mus muscle
   for _, m := range muscles{
     if m.Name == ""{
       continue
     }
     if strings.Contains(message, m.Name){
       fmt.Println("muscle found")
       mID = m.Id
       mus = m
       break
     }

   }

  if mID != -1 && exer.Category == mus.Id{
    return "YES"
  }

   return "NO"
 }

func GetImages(){

   var obj imagesJSON
   url := "https://wger.de/api/v2/exerciseimage/"
   f := true
   for f && url != ""{
     resp := getJSON(url)
     defer resp.Body.Close()
     if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
     log.Println(err)
     }

       exercises_img = append(exercises_img,obj.Results...)

       if url == obj.Next {
       f = false
     }
       url = obj.Next
   }
   fmt.Println("Images Done!")
 }

func GetEquipments(){

   var obj equipmentJSON
   url := "https://wger.de/api/v2/equipment/?language=2"
   f := true
   for f && url != ""{
     resp := getJSON(url)
     defer resp.Body.Close()
     if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
     log.Println(err)
     }

       equipments = append(equipments,obj.Results...)
       if url == obj.Next {
       f = false
     }
       url = obj.Next
   }
   fmt.Println("Equipments Done!")

 }

func GetMuscles(){

   var obj muscleJSON
   url := "https://wger.de/api/v2/exercisecategory"
   f := true
   for f && url != ""{
     resp := getJSON(url)
     defer resp.Body.Close()
     if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
     log.Println(err)
     }
       muscles = append(muscles,obj.Results...)
      if url == obj.Next {
       f = false
     }
       url = obj.Next

   }
   fmt.Println("Muscles Done!")
 }

func GetExercises(){

  var obj exerciseJSON
  url := "https://wger.de/api/v2/exercise/?language=2"
  f := true
  fmt.Printf("Loading ")
  for f && url != "" {
    resp := getJSON(url)
    defer resp.Body.Close()
    if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
    log.Println(err)
    }
      fmt.Printf("[]")
      exercises = append(exercises,obj.Results...)

      if url == obj.Next {
      f = false
    }
      url = obj.Next
  }
  fmt.Println("Exercises Done!")
}

func getJSON(inputUrl string) (*http.Response) {
  url := fmt.Sprintf(inputUrl)

  req, err := http.NewRequest("GET", url, nil)
   if err != nil {
  log.Fatal("NewRequest: ", err)
  }
  client := &http.Client{}

  resp, err := client.Do(req)
   if err != nil {
  log.Fatal("Do: ", err)
  }

  return resp

}

func home(w http.ResponseWriter, r *http.Request) {
    // var resultObj obj
    // url := "https://wger.de/api/v2/exercise/?language=2"
    // resultObj = getJSON(url);
    // for _, x := range resultObj.Results {
    //   fmt.Fprintf(w, "%v \n", x.Name)
    // }

    //fmt.Fprintf(w, "<h1>Welcome to the servers bitches</h1>") // write data to response
}

func api(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("welcome.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Fprintf(w, r.Form["username"][0])

    }
}

func CaseMatch(message string) (string){
  var result string
  result = Case3(message)
  if result != "err"{
    return  result +"\n"
  }
  result = Case6(message)
  if result != "err"{
    return  result + "\n"
  }
  result = Case7(message)
  if result != "err"{
    return  result + "\n"
  }
  result = Case5(message)
  if result != "err"{
    return  result + "\n"
  }
  result = Case1(message)
  if result != "err"{
    return  result + "\n"
  }
  result = Case2(message)
  if result != "err"{
    return  result + "\n"
  }
  return "err"

}

func containsElement(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


// func main() {
//     http.HandleFunc("/", home) // setting router rule
//     http.HandleFunc("/api", api)
//     	fmt.Printf("Listening on port %d...\n", 3000)
//       getMuscles()
//       getEquipments()
//       getImages()
//       getExercises()
//
//       for true {
//         fmt.Println("Hello, how can I help you?")
//         reader := bufio.NewReader(os.Stdin)
//         message, _ := reader.ReadString('\n')
//         fmt.Println("Scanned message = " + message)
//         fmt.Println(caseMatch(message))
//       }
//
//     err := http.ListenAndServe(":3001", nil) // setting listening port
//     if err != nil {
//         log.Fatal("ListenAndServe: ", err)
//     }
// }
