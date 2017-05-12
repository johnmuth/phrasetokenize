# phrasetokenize
you know, for phrases

Splits an input string into a slice of strings, where each string is either a word or a phrase - a phrase if wrapped in quotes.

E.g., 

```
Phrasetokenize(`"Hello world," she said`)
\\\returns []string{`"Hello world,"`, "she", "said"}
```
