# GOOEY

`gooey` is a no-frills package I wrote that wraps end-user-friendly interactive dropdowns, tables, and input prompts for `go` command line tools I use that assume a set of standard defaults involving the imported packages.

### EXAMPLE

```
package main

const defaultExitHelpText = ":q or :quit will quit the program"

func main() {
  confirmPrompt, err := gooey.NewConfirmPrompt("launch chrome?", true, true, defaultExitHelpText)
  if err != nil {
    log.Fatal(err)
  }

  if confirmPrompt.LaunchWithResponse() {
    launchChrome()
  }
}

func launchChrome() {
  // launch chrome browser
}

```