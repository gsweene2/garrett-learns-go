import os

omit_list = [ "convert-to-rst.py"]

def _get_files():
    files = [f for f in os.listdir('.') if os.path.isfile(f)]
    for file in omit_list:
        files.remove(file)
    return sorted(files)

def _write_title(writer, title):
    writer.write(f"# {title}" + "\n\n")

def _write_disclaimer(writer):
    writer.write(
        "This rst document is generated programatically from https://github.com/gsweene2/garrett-learns-go/blob/master/convert-to-md.py\n\n"
    )

def _write_header(writer, title):
    writer.write(f"## {title}\n\n")

def _write_subheader(writer, subheader):
    writer.write(f"### {subheader}\n\n")

def _write_code_block(writer, language):
    writer.write(f"```{language}" + "\n")

def _write_end_code_block(writer):
     writer.write("```" + "\n")

def _write_code_lines(writer, lines):
    for line in lines:
        writer.write(line)

with open(
    "/Users/garrettsweeney/git/garrett-mkdocs/docs/go.md", "w", newline=""
) as writer:
    _write_title(writer, "Go")
    _write_disclaimer(writer)
    
    # Write description and prepare code block for import lines to follow
    files = _get_files()
    # Remove README.md
    files.remove('README.md')
    files.remove('go.sum')
    files.remove('go.mod')
    files.remove('convert-to-md.py')

    for file in files:
        _write_header(writer, file)
        _write_subheader(writer, "Example")
        code = "md" if '.md' in file else "go"
        _write_code_block(writer, "go")
        with open(file) as f:
            _write_code_lines(writer, f.readlines())
        _write_end_code_block(writer)
        writer.write("\n\n")
