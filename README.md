# phrasetokenize
you know, for phrases

Splits an input string into a slice of strings, where each string is either a word or a phrase.

E.g., 

```
Phrasetokenize(`"Hello world," she said`)
\\\returns []string{`"Hello world,"`, "she", "said"}
```
