package main

import (
	"os/exec"
	"log"
	"fmt"
	"encoding/json"
)

type Result struct {
	Code int `json:"errcode"`
	Msg string `json:"errmsg"`
}

func main() {
	//send dingding robot message
	r, err := exec.Command("bash", "-c", `curl https://oapi.dingtalk.com/robot/send?access_token=e78992df5c1aa619bcf6b83044dc341d122fc \
        -H 'Content-Type: application/json' \
        -d '
          {"msgtype": "text",
            "text": {
                "content": "test message, pleat ignore these"
             }
          }'`).Output()

    if err != nil {
    	fmt.Println(err)
    	return
	}

	// convert json result to struct
    result := Result{}
    err = json.Unmarshal(r, &result)
    if err != nil {
    	fmt.Println("json decode failed: ", err)
	}

	fmt.Printf("%+v\n", result)
	if result.Code != 0 {
		fmt.Println(result.Msg)
	}

	//find the command path
	path, err := exec.LookPath("curl")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

	// call date print current date
	out, err := exec.Command("date", `+"%Y-%m-%d"`).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", string(out))

}
