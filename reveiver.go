package main

import (
	"encoding/json"
	"fmt"
	"github.com/gen2brain/beeep"
	"log"
	"net/http"
	"os"
)

type MsgStruct struct {
	Phone           string
	Slot, Msg, Date string
}

func msgHandler(w http.ResponseWriter, r *http.Request) {
	var msgT MsgStruct
	d := json.NewDecoder(r.Body)
	fmt.Println(r.RemoteAddr)
	er := d.Decode(&msgT)
	if er != nil {
		panic(er)
	}
	str := fmt.Sprintf("%s,FROM %s, %s", msgT.Date, msgT.Phone, msgT.Slot)
	err := beeep.Notify(str, msgT.Msg, "assets/information.png")
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", msgHandler)
	port := "8081"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Printf("Listening: http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
