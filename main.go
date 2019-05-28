package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chmllr/imgtb/health"
	"github.com/chmllr/imgtb/imp"
	"github.com/chmllr/imgtb/index"

	"github.com/fatih/color"
)

func main() {
	lib := flag.String("lib", "", "path to the photo library")
	src := flag.String("src", "", "source directory")
	deep := flag.Bool("deep", false, "deep check (includes md5 comparison)")
	flag.Parse()

	if len(flag.Args()) != 1 {
		printHelp()
		os.Exit(1)
	}

	if *lib == "" {
		var err error
		*lib, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := flag.Args()[0]

	switch cmd {
	case "import":
		if *src == "" {
			log.Fatal("no source folder specified")
		}
		fmt.Printf("importing to %q from %q...\n", *lib, *src)
		refs, err := imp.Import(*lib, *src)
		if err != nil {
			log.Fatalf("couldn't import: %v", err)
		}
		sealed, err := index.Index(*lib)
		if err != nil {
			log.Fatalf("couldn't get index: %v", err)
		}
		for _, ref := range sealed {
			refs = append(refs, ref)
		}
		index.Save(*lib, refs)
	case "repair":
		fmt.Printf("repairing %q...\n", *lib)
		files, err := index.Report(*lib, true)
		if err != nil {
			log.Fatalf("couldn't get report: %v", err)
		}
		index.Save(*lib, files)
	case "health":
		fmt.Printf("checking (deep: %t) health of %q...\n", *deep, *lib)
		libRefs, err := index.Report(*lib, *deep)
		if err != nil {
			log.Fatalf("couldn't get report: %v", err)
		}
		corrupted, found, sealed, duplicates, err := health.Verify(*lib, *deep, libRefs)
		if err != nil {
			log.Fatalf("couldn't verify: %v", err)
		}
		for _, path := range corrupted {
			color.Red("File %s is corrupted!", path)
		}
		for path := range sealed {
			color.Red("File %s is missing!", path)
		}
		for _, paths := range duplicates {
			color.Yellow("These files are duplicates:")
			for _, path := range paths {
				color.Yellow(" - %s", path)
			}
		}
		for path := range found {
			color.Cyan("File %s is new!", path)
		}
		if len(corrupted) == 0 && len(sealed) == 0 && len(duplicates) == 0 && len(found) == 0 {
			if *deep {
				fmt.Printf("%q is in perfect health! ✅\n", *lib)
			} else {
				fmt.Printf("%q is in a good health (use --deep for a complete check)! ✅\n", *lib)
			}
		}
	default:
		fmt.Printf("Error: unknown command %q\n\n", cmd)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`Usage: imgtb --lib <PATH> [OPTIONS] <COMMAND>

Avaliable commands:

import (accepts option --src <PATH>):
	Imports all media files from the specified source path into the lib folder.
	It creates the corresponding folder structure (<lib>/YYYY/MM/DD) if necessary.
	If no source folder was specified, the current directory is used.

repair:
	Seals all existing files with their md5 hashes into the lib index. It does not
	make any mutating operations on the library!

health (accepts option --deep):
	Checks existing file structure against the index. This command can detect 
	missing, modified, duplicated and new files. If option 'deep' is proveded, 
	checks the file hash as well.`)

}
