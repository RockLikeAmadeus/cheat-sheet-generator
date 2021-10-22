# cheat-sheet-generator
Wanted a simplified way of managing my crib notes/reference files.

## Usage

~~Ideally, with the path to the CSG binary added to the PATH env variable, executing csg should generate markdown files _located at and based on the current working directory where the csg command is run_.~~

Let's try a usage scenario where the executable lives in the directory where the root directory for each set of reference notes also live. In this case, CSG will generate a separate file for each directory that is a sibling to the CSG executable itself.