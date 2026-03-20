package ui

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/harshalve/ctx-tool/internal/storage"
)

func HandleList(store *storage.Storage) {
	projects, err := store.ListAll()
	if err != nil {
		fmt.Printf("❌ Error retrieving projects: %v\n", err)
		return
	}

	if len(projects) == 0 {
		fmt.Println("\n📭 No contexts saved yet. Try saving one with:")
		fmt.Println("   ctx save <name>")
		return
	}

	fmt.Println("\n📂 Saved Contexts:")

	// Create a tabwriter to handle column alignment
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintln(w, "NAME\tBRANCH\tLINKS\tTASKS")
	fmt.Fprintln(w, "----\t------\t-----\t-----")

	for _, name := range projects {
		p, err := store.Load(name)
		if err != nil {
			// If one file is corrupted, we skip it and continue
			continue
		}

		fmt.Fprintf(w, "%s\t%s\t%d\t%d\n",
			p.Name,
			p.Branch,
			len(p.Links),
			len(p.Terminals))
	}

	w.Flush() // Write the buffered data to the terminal
	fmt.Println("")
}
