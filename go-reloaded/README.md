# Text Completion & Editing Tool

## 📖 Introduction
This project is a simple text completion, editing, and auto-correction tool written in **Go**.  
It processes a text file, applies a series of transformations, and outputs the corrected text into another file.  

The tool is designed to:
- Practice Go file system (`fs`) APIs
- Improve skills in string and number manipulation
- Encourage good coding practices and unit testing
- Provide peer-audited code review experience

---
## 🚀 Usage
```bash
# Compile and run
go run . sample.txt result.txt

# Example
cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom...

go run . sample.txt result.txt

cat result.txt

It was the best of times, it was the worst of TIMES, it was the age of wisdom...
