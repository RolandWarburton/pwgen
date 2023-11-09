# Passphrase Generator

**WIP**

I really like passwords in the format of `^(\S+-){2,}\d+$`.

There are some tools that exist that meet this requirement,
in particular [pgen](https://github.com/ctsrc/Pgen).

I wrote this one mainly as a programming exercise.

## Goal

* Generate passphrases in a *word-word-number* format.
* Variable number of words using CLI flags
* Embed a word list inside the binary.
* Decently fast and lightweight.
