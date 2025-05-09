#!/usr/bin/env python3
"""
convert_bw.py – resize images to 28×28, increase contrast, make B/W
Usage:  python convert_bw.py <input_dir> <output_dir>
Requires: Pillow  (pip install pillow)
"""

from pathlib import Path
from PIL import Image, ImageOps, ImageEnhance
import sys

def process_image(src: Path, dst_dir: Path,
                  contrast=2.0, thresh=90):
    with Image.open(src) as im:
        im = ImageOps.fit(im.convert("L"), (28, 28),
                          Image.Resampling.LANCZOS)
        im = ImageEnhance.Contrast(im).enhance(contrast)
        bw = im.point(lambda p: 255 if p > thresh else 0, mode="1")

        out = dst_dir / (src.stem + ".png")   # force lossless PNG
        bw.save(out, format="PNG", dither=Image.NONE)

def main():
    if len(sys.argv) != 3:
        print("Usage: python convert_bw.py <input_dir> <output_dir>")
        sys.exit(1)

    in_dir, out_dir = Path(sys.argv[1]), Path(sys.argv[2])
    out_dir.mkdir(parents=True, exist_ok=True)

    for f in in_dir.iterdir():
        if f.is_file():
            try:
                process_image(f, out_dir)
                print(f"✔ {f.name}")
            except Exception as e:
                print(f"✖ {f.name}: {e}")

if __name__ == "__main__":
    main()