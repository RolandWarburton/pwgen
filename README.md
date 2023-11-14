# Passphrase Generator

**WIP**

I really like passwords in the format of `^(\S+-){2,}\d+$`.

There are some nice tools that exist to meet this requirement already,
in particular [pgen](https://github.com/ctsrc/Pgen).

I wrote this one mainly as a programming exercise.

## Goal

* Generate passphrases in a *word-word-number* format.
* Variable number of words using CLI flags
* Embed a word list inside the binary.
* Decently fast and lightweight.

## Examples

```bash
# generate a password 2 words long
pwgen -w 2
# OUTPUT: coral-icon-4$

# generate a password with words >2 and <4
pwgen --gt 2 --lt 4
# OUTPUT: hurt-lash-10!

# generate a password with a custom delimiter
pwgen -d '&'
# OUTPUT: foil&ramp&6%

# generate a password with a prepended word
pwgen --prepend 'secret-'
# OUTPUT: secret-yield-ivory-4&

# generate a password with an appended word
pwgen --append '-secret'
# OUTPUT: icing-prong-secret

# generate 3 passwords
pwgen -n 3
# OUTPUT:
# avid-union-6#
# judge-boxer-10&
# wham-nacho-9$
```

## Credits

[eff.org/dice](https://www.eff.org/dice) for a well authored word list.
