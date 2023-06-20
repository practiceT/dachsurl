SOURCE	= ../README.md
TARGET	= ./content/_index.md

build_with_hugo: clean
	@cp $(SOURCE) $(TARGET)
	@sed -i -e 's/.\/docs\/static\/images/.\/images/g' $(TARGET)
	hugo --ignoreCache

clean:
	rm -rf public/*
	rm -f $(TARGET)
