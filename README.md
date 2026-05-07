# GoRenamer

Rename or delete files in bulk using prefixes, suffixes, string replacement, or regex patterns.

## Requirements

- [Go](https://go.dev/dl/) 1.18+

## Install

```
git clone https://github.com/jackyes89/GoRenamer.git
cd GoRenamer
go build -o GoRenamer.exe GoRename.go
```

## Usage

```
GoRename -path <directory> [flags]
```

The `-path` flag is always required. All other flags are optional, but only one operation runs per invocation (the first matching flag wins, checked in the order below).

## Flags

| Flag | Description |
|---|---|
| `-path` | Target directory |
| `-prefix` | Add string before each filename |
| `-append` | Insert string before the file extension |
| `-tp` | Prefix with current date (`YYYY-MM-DD`) |
| `-ta` | Append current date before the extension |
| `-StrToReplace` | Substring to find (use with `-StrReplacer`) |
| `-StrReplacer` | Replacement string (use with `-StrToReplace`) |
| `-RegExPattern` | Regex pattern to match (use with `-RegExReplace`) |
| `-RegExReplace` | Replacement for regex matches (use with `-RegExPattern`) |
| `-RegExRemove` | Regex pattern — matching files are deleted |
| `-FileType` | Filter by extension (e.g. `.txt`, `.dwg`) |
| `-dr` | Apply operation to directories as well |
| `-h` | Show help |

## Examples

Add a prefix to every file:

```
GoRename -path ./docs -prefix "draft_"
```

Append a version tag before the extension:

```
GoRename -path ./images -append "_v2"
```

Rename all `.pdf` files with today's date as prefix:

```
GoRename -path ./invoices -tp -FileType .pdf
```

Replace a substring:

```
GoRename -path ./logs -StrToReplace "error" -StrReplacer "ERROR"
```

Rename using regex capture groups:

```
GoRename -path ./photos -RegExPattern "IMG_(\d+)" -RegExReplace "Photo_$1"
```

Delete all `.tmp` files:

```
GoRename -path ./temp -RegExRemove "\.tmp$"
```

Rename directories too:

```
GoRename -path ./project -prefix "2024_" -dr
```

## Notes

- Back up your data before running bulk operations.
- You need write permission on the target directory.
- Invalid regex passed to `-RegExPattern` or `-RegExRemove` prints an error and exits instead of crashing.
