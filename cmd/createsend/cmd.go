package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sourcegraph/createsend-go/createsend"
)

var apiclient *createsend.APIClient

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: createsend command [OPTS] ARGS...\n")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "The commands are:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tlist-clients")
		fmt.Fprintln(os.Stderr, "\tlist-lists       CLIENT")
		fmt.Fprintln(os.Stderr, "\tlists-for-email CLIENT EMAIL")
		fmt.Fprintln(os.Stderr, "\tlist-subscribers LIST (active|unconfirmed|unsubscribed|bounced|deleted)")
		fmt.Fprintln(os.Stderr, "\tget-subscriber   LIST EMAIL")
		fmt.Fprintln(os.Stderr, "\tadd-subscriber   LIST EMAIL")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Common arguments:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tCLIENT:\ta client ID")
		fmt.Fprintln(os.Stderr, "\tLIST:\ta list ID")
		fmt.Fprintln(os.Stderr, "\tEMAIL:\temail address")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Run `createsend command -h` for more information.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}
	log.SetFlags(0)

	apiKey := os.Getenv("CREATESEND_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: you must set your Campaign Monitor API key in the CREATESEND_API_KEY environment variable.")
	}
	authClient := &http.Client{
		Transport: &createsend.APIKeyAuthTransport{APIKey: apiKey},
	}
	apiclient = createsend.NewAPIClient(authClient)

	subcmd := flag.Arg(0)
	remaining := flag.Args()[1:]
	switch subcmd {
	case "list-clients":
		listClients(remaining)
	case "list-lists":
		listLists(remaining)
	case "lists-for-email":
		listsForEmail(remaining)
	case "list-subscribers":
		listSubscribers(remaining)
	case "get-subscriber":
		getSubscriber(remaining)
	case "add-subscriber":
		addSubscriber(remaining)
	}
}

func listClients(args []string) {
	clients, err := apiclient.ListClients()
	if err != nil {
		log.Fatalf("Error listing clients: %s\n", err)
	}
	if len(clients) == 0 {
		fmt.Println("No clients found.")
		return
	}
	for _, c := range clients {
		fmt.Printf("%-24s %s\n", c.Name, c.ClientID)
	}
}

func listLists(args []string) {
	if len(args) != 1 {
		log.Println("list-subscribers takes 1 argument.")
		flag.Usage()
	}

	clientID := args[0]
	lists, err := apiclient.ListLists(clientID)
	if err != nil {
		log.Fatalf("Error listing lists: %s\n", err)
	}
	if len(lists) == 0 {
		fmt.Println("No lists found.")
		return
	}
	for _, c := range lists {
		fmt.Printf("%-24s %s\n", c.Name, c.ListID)
	}
}

func listsForEmail(args []string) {
	if len(args) != 2 {
		log.Println("lists-for-email takes 2 arguments.")
		flag.Usage()
	}

	clientID, email := args[0], args[1]
	lists, err := apiclient.ListsForEmail(clientID, email)
	if err != nil {
		log.Fatalf("Error listing lists for email address %q: %s\n", email, err)
	}
	if len(lists) == 0 {
		fmt.Printf("No lists found for email address %q.\n", email)
		return
	}
	for _, c := range lists {
		fmt.Printf("%-44s %s  %s\n", c.ListName, c.ListID, c.DateSubscriberAddedStr)
	}
}

func listSubscribers(args []string) {
	if len(args) != 2 {
		log.Println("list-subscribers takes 2 arguments.")
		flag.Usage()
	}

	listID, group := args[0], createsend.SubscriberGroup(args[1])
	subs, err := apiclient.ListSubscribers(listID, group, nil)
	if err != nil {
		log.Fatalf("Error listing subcribers for list %q: %s\n", listID, err)
	}
	if len(subs) == 0 {
		fmt.Println("No subscribers found.")
		return
	}
	for _, c := range subs {
		fmt.Printf("%-24s %s\n", c.EmailAddress, c.Name)
	}
}

func getSubscriber(args []string) {
	if len(args) != 2 {
		log.Println("get-subscriber takes 2 arguments.")
		flag.Usage()
	}

	listID, email := args[0], args[1]
	sub, err := apiclient.GetSubscriber(listID, email)
	if err != nil {
		log.Fatalf("Error getting subcriber %q for list %q: %s\n", email, listID, err)
	}
	fmt.Printf("%+v\n", sub)
}

func addSubscriber(args []string) {
	if len(args) != 2 {
		log.Println("get-subscriber takes 2 arguments.")
		flag.Usage()
	}

	listID, email := args[0], args[1]
	err := apiclient.AddSubscriber(listID, createsend.NewSubscriber{EmailAddress: email})
	if err != nil {
		log.Fatalf("Error adding subcriber %q to list %q: %s\n", email, listID, err)
	}
	fmt.Printf("Added subscriber %q to list %q.\n", email, listID)
}
