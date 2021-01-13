package cmds

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/golevi/cfallow/cfa"
)

var sleepDuration = time.Millisecond * 300

// AddFile a file of IPs.
// There are some sleeps in this function to help prevent hammering the CF API.
//
// https://api.cloudflare.com/
// The Cloudflare API sets a maximum of 1,200 requests in a five minute period.
// Or 240/minute or 4 calls every second.
// Allowing 250ms for each call.
func AddFile(filename string) {
	fmt.Println("Starting...")

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	ips := strings.Split(string(data), "\n")

	for _, ac := range cfa.ListAccounts().Result {
		fmt.Println(ac.Name)
		line("=", 80)

		for _, ip := range ips {
			hash := md5.New()
			hash.Write([]byte(ip))
			h := hash.Sum(nil)

			name := fmt.Sprintf("%x", h)

			fmt.Printf("Rule hash %s\n", name)
			fmt.Printf("IP %s\n", ip)

			rules := cfa.GetRulesByNote(ac.ID, name)
			for _, r := range rules.Result {
				del := cfa.DeleteAccessRule(ac.ID, r.ID)
				if del {
					fmt.Printf("Deleted rule %v\n", r.ID)
				}
				time.Sleep(sleepDuration)
			}

			car := cfa.CreateAccessRule(ac.ID, ip, name)
			if car {
				fmt.Println("IP Added...")
			}
			line("-", 80)
			time.Sleep(sleepDuration)
		}
		time.Sleep(sleepDuration)
	}

	fmt.Println("Done!")
}

func line(char string, len int) {
	for i := 0; i < len; i++ {
		fmt.Print(char)
	}
	fmt.Print("\n")
}
