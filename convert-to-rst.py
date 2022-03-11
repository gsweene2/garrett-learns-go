import os

omit_list = [ "convert-to-rst.py"]

def _get_files():
    files = [f for f in os.listdir('.') if os.path.isfile(f)]
    for file in omit_list:
        files.remove(file)
    return sorted(files)

def _write_title(writer, title):
    writer.write(f"{title}" + "\n")
    writer.write("=" * len(title) + "\n")
    writer.write(".. meta::" + "\n")
    writer.write(f"   :description lang=en: {title} docs" + "\n\n")

def _write_disclaimer(writer):
    writer.write(".. warning::\n\n")
    writer.write(
        "    This rst document is generated programatically from https://github.com/gsweene2/garrett-learns-go/blob/master/convert-to-rst.py\n\n"
    )

def _write_header(writer, title):
    writer.write(title + "\n")
    writer.write("-" * len(title) + "\n")

def _write_subheader(writer, subheader):
    writer.write(subheader + "\n")
    writer.write("*" * len(subheader) + "\n")

def _write_code_block(writer, language):
    writer.write(f".. code-block:: {language}" + "\n")
    writer.write("  :linenos:" + "\n")
    writer.write("" + "\n")
    writer.write("" + "\n")
    writer.write("" + "\n")

def _write_code_lines(writer, lines):
    for line in lines:
        writer.write("  " + line)

with open(
    "/Users/garrettsweeney/git/garrett-docs/languages/go.rst", "w", newline=""
) as writer:
    _write_title(writer, "Go")
    _write_disclaimer(writer)
    
    # Write description and prepare code block for import lines to follow
    files = _get_files()
    for file in files:
        _write_header(writer, file)
        _write_subheader(writer, "Example")
        code = "md" if '.md' in file else "go"
        _write_code_block(writer, "go")
        with open(file) as f:
            _write_code_lines(writer, f.readlines())
        writer.write("\n\n")
