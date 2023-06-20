SOURCE	= ../README.md
TARGET_DIR	= ./content
TARGET_FILE	= $(TARGET_DIR)/_index.md

build_with_hugo: clean
	@mkdir -p $(TARGET_DIR)
	@cp $(SOURCE) $(TARGET_FILE)
	@sed -i -e 's/.\/docs\/static\/images/.\/images/g' $(TARGET_FILE)
	hugo --ignoreCache

clean:
	rm -rf public/*
	rm -rf $(TARGET_DIR)
