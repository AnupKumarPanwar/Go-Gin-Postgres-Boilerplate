package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Help() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 20, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "Commands")
	fmt.Fprintf(w, "\n %s\t%s\t", "serve", "# starts the server")
	fmt.Fprintf(w, "\n %s\t%s\t", "db:create", "# creates the database")
	fmt.Fprintf(w, "\n %s\t%s\t", "db:migrate", "# runs sql migrations on the database")
	fmt.Fprintf(w, "\n %s\t%s\t", "db:drop", "# drops the database")
	fmt.Fprintf(w, "\n")
}
