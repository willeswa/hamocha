# Hamocha CLI Tool

A CLI tool in Go that organizes files within a directory into folders based on their file type. This tool was developed as a learning project for handling file operations in Go.

## Features
- Categorizes files by type (e.g., Images, Documents, Videos).
- Creates destination folders if they don’t already exist.
- Moves files into their respective folders by extension.
- Easily accessible from anywhere on your system once installed.

---

## Installation Steps

### 1. Clone the Repository


```
git clone https://github.com/yourusername/hamocha.git
cd hamocha
```

### Build the Executable
```
go build -o hamocha
```
### Move to PATH
Move the executable to a directory in your PATH to make it accessible globally:

```
sudo mv hamocha /usr/local/bin/
```

### Usage
Navigate to the directory you want to organize and run:

```
hamocha <directory-path>
````

For example:

```
hamocha /Users/username/Downloads
````

This command organizes all files in the Downloads folder into subfolders by type (e.g., Images, Documents, Videos).

Supported File Types
Currently, Hamocha supports these categories and extensions:

```
Images: .jpg, .jpeg, .png, .gif, .bmp
Documents: .pdf, .doc, .docx, .txt, .xls, .xlsx, .ppt, .pptx
Videos: .mp4, .avi, .mov, .mkv
Music: .mp3, .wav, .aac
Archives: .zip, .rar, .tar, .gz
Additional types and extensions can be easily added to the fileTypes map in the code.
```

Example Output
After running the tool, it will display messages indicating the folders created and files moved, such as:

Created folder: Images
Moved file sample.jpg to Images/sample.jpg
Created folder: Documents
Moved file resume.pdf to Documents/resume.pdf
Future Improvements
Add more file types.
Handle duplicate filenames in the destination folder.
Enable organization by other criteria (e.g., date modified).
