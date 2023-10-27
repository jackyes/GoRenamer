# GoRenamer

This program provides a convenient way to rename files in a directory using the provided flags. It supports adding prefixes, appending strings, renaming with a current date, and even using regex patterns for more advanced renaming capabilities.

## Features

- **Prefix & Append**: Easily add a prefix or append a string to file names.
- **Date Formatting**: Add current date (YYYY-MM-DD) as a prefix or append it to file names.
- **String Replacement**: Replace a specific string in the filename.
- **RegEx Support**: Use regular expressions for advanced renaming.
- **File Type Filter**: Rename only files with a specific extension.
- **Remove by RegEx**: Delete files whose names match a specific regex pattern.
- **Directory Rename**: Apply the renaming also to directories.

## Usage:  
  
Usage of ./GoRename:  
  -append string  
    Add <string> at the end of filename  
  -dr  
    Add -dr to modify also Directory name  
  -FileType string  
    Rename only file with extension (Example: .dwg)  
  -path string  
    Select path.  
  -prefix string  
    Add <string> before filename  
  -RegExPattern string  
    RegEx Pattern. (Use with -RegExReplace <string>.  
  -RegExReplace string  
    String to replace the matched regex. (Use with -RegExPattern <string>.)  
  -RegExRemove string  
    Remove file if match regex.  
  -StrReplacer string  
    String to be replaced on renamed filename. (Use with -StrToReplace <string>.)  
  -StrToReplace string  
    String to be replaced on original filename. (Use with -StrReplacer <string>.)  
  -ta  
    Append YYYY-MM-DD to filename.  
  -tp  
    Add YYYY-MM-DD as prefix.  
  -h  
    help for ./GoRename  
    
## Examples:  
  
Prefix all files in the current directory with "new_"  
./GoRename -path . -prefix new_  
  
Append ".bak" to the end of all files in the current directory  
./GoRename -path . -append .bak  
  
Replace all instances of "old" with "new" in the filenames of all files in the current directory  
./GoRename -path . -StrToReplace old -StrReplacer new  
  
Use regular expression to remove all numbers from the filenames of all files in the current directory  
./GoRename -path . -RegExRemove '[0-9]+'  
  
Remove all files in the current directory that have the extension ".txt"  
./GoRename -path . -RegExRemove '.txt$'  
  
Add the current date and time to the beginning of the filenames of all files in the current directory  
./GoRename -path . -tp  
  
Add the current date and time to the end of the filenames of all files in the current directory  
./GoRename -path . -ta  
  


## Notes

- Always back up your data before making bulk renaming operations.
- Ensure you have the necessary permissions for renaming and deleting in the specified directory.
- The -dr flag must be used to modify the names of directories.
- The -FileType flag can be used to filter the files that are renamed.
- The -h flag can be used to display this help message.
