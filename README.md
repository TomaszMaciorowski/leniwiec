# leniwiec
Program functionality:
Scanning the source directory: The program traverses the source directory, searching for files with specified extensions (e.g., .jpg, .png).
Reading the date from EXIF metadata: If an image file contains EXIF metadata (e.g., the photo's creation date), the program uses that date to generate the file name.
File naming: Each file is given a name in the format YYYYMMDDHHMMSS_INDEX.jpg, where INDEX is the file's sequential number.
Moving files: Files are moved from the source directory to the destination directory with their new names.

Requirements:
The program requires three arguments:
Path to the source directory.
Path to the destination directory.
List of file extensions separated by commas (e.g., jpg,png).
Example usage:j
leniwiec.exe /path/to/source /path/to/destination jpg,png
