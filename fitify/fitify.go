package fitify

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/renstrom/fuzzysearch/fuzzy"
    "strings"
    "math/rand"
)

type aiJSON struct {
  Cnt string `json: cnt`
}
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
   "how can i train",
   "what can i train using",
 }
 var case_3  = []string{"tell me exercises that train using",
   "which exercises train using",
   "how can i train using",
   "how can i workout my using",
   "workouts using",
   "work outs using"}

 var case_5 = []string{"image",
    "pictures",
    "photo",
    "pic",
    "img"}
 var case_6 = []string{"Does use ", "Is needed for ", "Are needed for "}
 var case_7 = []string{"Does train", "Does workout ", "Does work out "}
 var case_8 = []string{"hi" , "hello", "greetings", "hey", "sup", "howdy", "salam alaikom", "salam 3alaikom", "hallo", "bonjour"}
 var case_9 = []string{"I'm not really sure how to answer that", "I'm not sure", "I don't understand what you mean", "Try asking that another way"}

func Case1(message string) (string){
   input := []string{message}
   f := false
   for _,q := range case_1 {
     fuzzyMatches := fuzzy.Find(strings.ToLower(q), input)
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
      if containsIgnoreCase(message, eq.Name){
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
        result += ex.Name + ", \n" // + ex.Description
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
     fuzzyMatches := fuzzy.Find(strings.ToLower(q), input)
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
     if containsIgnoreCase(message, m.Name){
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
        result += ex.Name + ", \n" // + ex.Description
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
     fuzzyMatches := fuzzy.Find(strings.ToLower(q), input)
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
     if containsIgnoreCase(message, m.Name){
       muscleID = m.Id
     }
   }
   if muscleID == -1{
     fmt.Println("fail muscle")
     return "err"
   }
      equipmentID := -1
    for _,eq := range equipments {
      if containsIgnoreCase(message, eq.Name){
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
        result += ex.Name + ", \n" // + ex.Description
        for _,img := range exercises_img {
          if img.Exercise == ex.Id {
            result += "<img src=\""+img.Image + "\" style=\"width:350px;height:200px;\"/>" +"\n"
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
     fuzzyMatches := fuzzy.Find(strings.ToLower(q), input)
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
     if containsIgnoreCase(message, ex.Name){
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
       result +=  "<img src=\""+img.Image + "\" style=\"width:350px;height:200px;\"/>" +"\n"
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
     fuzzyMatches := fuzzy.Find(strings.ToLower(q), input)
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
     if containsIgnoreCase(message, ex.Name){
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
     if containsIgnoreCase(message, eq.Name){
       fmt.Println("equipment found")
       eqID = eq.Id
       break
     }

   }

   for _,eq := range exer.Equipment{
     fmt.Printf("is %d == %d ?\n" , eq , eqID)
     if eqID == eq {
       return "Yes, in some variations of" + exer.Name
     }
   }
   if(eqID != -1){
     return "Nope."
   }
   return "I'm not really sure"
 }

func Case7(message string) (string){
   input := []string{message}
   f := false

   for _, q := range case_7 {
     fuzzyMatches := fuzzy.Find(strings.ToLower(q), input)
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
     if containsIgnoreCase(message, ex.Name){
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
     if containsIgnoreCase(message, m.Name){
       fmt.Println("muscle found")
       mID = m.Id
       mus = m
       break
     }

   }

  if mID != -1 && exer.Category == mus.Id{
    return "Yes."
  }
  if(mID != -1){
    return "No."
  }
  return "I'm not really sure"
 }

func Case8(message string)(string){
  input := []string{message}
  f := false

  for _, q := range case_8 {
    fuzzyMatches := fuzzy.Find(strings.ToLower(q), input)
    if len(fuzzyMatches) != 0{
      f = true
    }
  }

  if !f {
    fmt.Println("flag is false")
    return "err"
  }
  return case_8[rand.Intn(len(case_8))]
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

 func GetScriptResponse(message string)(string){
   message = strings.Replace(message, " ", "+", -1)
   var obj aiJSON
   url := "http://api.acobot.net/get?bid=413&key=3LQGr50kQPVmmR9x&uid=1234567890&msg=" + message;

   resp := getJSON(url)
   defer resp.Body.Close()
   if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {

   log.Println(err)
   }
   return obj.Cnt
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

func CaseMatch(message string) (string){
  message = strings.ToLower(message)

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
  // result = Case8(message)
  // if result != "err"{
  //   return  result + "\n"
  // }

  // return case_9[rand.Intn(len(case_9))]
  return GetScriptResponse(message)
  }

func containsElement(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
  }

func containsIgnoreCase(str1 string, str2 string) (bool){
  return strings.Contains(strings.ToLower(str1), strings.ToLower(str2))
  }
