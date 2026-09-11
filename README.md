# ts

Tools for the **Top Secret: New World Order** espionage RPG.

## `nwo/gen` — attribute generator

Rolls a random set of the five TS:NWO attributes — Nerve, Suave, Pulse,
Intellect, Reflex — as die-type ratings (`d4` … `d12+d4`) and prints them as a
Markdown table:

```
| Attribute | Level |
|:----|----:|
| Nerve | d8 |
| Suave | d6 |
| Pulse | d10 |
| Intellect | d6 |
| Reflex | d8 |
```

`attributes/interface.go` defines the `Level` scale and the
`ToLevelRoundUp` / `HalfRoundedUp` helpers used to convert calculated values
back to legal die types, plus the two standard starting arrays
(`ArrayOne`, `ArrayTwo`).

## Run

The code imports the module path `ts`, so run from the repo root (init a module
first if needed):

```bash
go mod init ts   # only if go.mod is missing
go run ./nwo/gen
```
