Bowser.app: clean Makefile src/*.go src/*.h src/*.m src/*.plist assets/*.icns assets/*.png
	mkdir -p dist/Bowser.app/Contents/MacOS dist/Bowser.app/Contents/Resources
	cd src && go build -i -o ../dist/Bowser.app/Contents/MacOS/bowser
	cp assets/bowser.png dist/Bowser.app/Contents/MacOS
	cp src/Info.plist dist/Bowser.app/Contents
	cp assets/icon.icns dist/Bowser.app/Contents/Resources

.PHONY: install
install: Bowser.app
	cp -Rf dist/Bowser.app /Applications

.PHONY: clean
clean:
	-rm -Rf dist/Bowser.app
