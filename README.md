
# Table of Contents

1.  [The Problem](#org5da7aa9)
2.  [Requirements](#orgdb369ef)
3.  [Usage](#orgf6af0bf)
4.  [Demo](#org7f38546)

Small-one-file utility to convert clipboard images to text which automatically copied into
clipboard, using Tesseract, the Open Source OCR Engine


<a id="org5da7aa9"></a>

# The Problem

Taking notes from some PDFs was too painful since I had to write down the text if I wasn&rsquo;t
capable to copy it (which is the case in image-only PDFs, no text layer, which, also,
basically 99,9% of Arabic PDFs). This utility ease the process.


<a id="orgdb369ef"></a>

# Requirements

-   Tesseract
-   Unix-like OS


<a id="orgf6af0bf"></a>

# Usage

After installing (use the `go install` command), basically run:

    tesseract-to-xclip

Tesseract has script detection within &ldquo;OSD&rdquo;, but not language detection, so you cannot detect
language automatically you have to specify language. By default, it assumes that you want to
use english. To use it with your desired language, pass its ISO 639-3 identifier, for
example to run it with Arabic:

    tesseract-to-xclip ara

Check [Languages supported in different versions of Tesseract](https://tesseract-ocr.github.io/tessdoc/Data-Files-in-different-versions.html)


<a id="org7f38546"></a>

# Demo



https://user-images.githubusercontent.com/74098495/181529497-d87cc2cd-9023-489d-811d-eb4eab3de724.mp4

