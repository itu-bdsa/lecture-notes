.PHONY: all watch preview

all:
	../../remark_presi.py Slides.md ./Slides.html

preview:
	../../remark_presi.py --preview Slides.md ./Slides.html

Slides.pdf: Slides.md
	pandoc --pdf-engine=xelatex Slides.md -o Slides.pdf

watch:
	fswatch -o Slides.md | xargs -n1 -I{} make preview
