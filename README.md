# cheat-sheet-generator
Wanted a simplified way of managing my crib notes/reference files.

## Usage

~~Ideally, with the path to the CSG binary added to the PATH env variable, executing csg should generate markdown files _located at and based on the current working directory where the csg command is run_.~~

Let's try a usage scenario where the executable lives in the directory where the root directory for each set of reference notes also live. In this case, CSG will generate a separate file for each directory that is a sibling to the CSG executable itself.

## TODO:
- I'd like to change the order that this gets rendered in so that all leaf-level notes wind up at the end of the generated file, but that might not be trivial.
    - This will possibly require writing to the file in the middle of the file, rather than always appending to the end.
- This should also generate github style markdown in the form of README.md files.
- Once the above is implemented, the usage will have to change, since each top-level directory that CSG uses to generate files will need to be its own GH repo with a single README.md file.
    - Maybe CSG can reference a central config file that lists the path to all directories against which to run, so we can re-generate all reference files at once from any working directory easily.
    - Perhaps CSG could then output a separate .bash file or something to commit the changes for each reference file to git.
        - Ideally, only perform the commit/push if the file actually changed? This shouldn't be too difficult to do if, before re-writing all of the reference files, we gather a hash or checksum for each existing file, and then afterwords compare these with newly generated hash/checksums.
- Elegantly handle cases where multiple sections of the same heading (with different parents).
    - This could be done by specifying the link using the full path to the section, including each parent section, as in "Python > Collections > Lists | Lists" rather than just "Lists".