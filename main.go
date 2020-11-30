package main

import (
	"fmt"
	"os"

	"github.com/golevi/cfallow/cfa"
)

func main() {
	ip := cfa.GetIP()
	name, _ := os.Hostname()

	fmt.Println("Starting...")

	fmt.Printf("Your IP %s\n", ip)
	fmt.Println("=========================================================")
	for _, ac := range cfa.ListAccounts().Result {
		fmt.Println(ac.Name)
		rules := cfa.GetRulesByNote(ac.ID, name)

		for _, r := range rules.Result {
			del := cfa.DeleteAccessRule(ac.ID, r.ID)
			if del {
				fmt.Printf("Deleted rule %v\n", r.ID)
			}
		}

		// Create rule
		car := cfa.CreateAccessRule(ac.ID, ip, name)
		if car {
			fmt.Println("IP Added...")
		}

		fmt.Println("=========================================================")
	}

	fmt.Println("Done!")
}
