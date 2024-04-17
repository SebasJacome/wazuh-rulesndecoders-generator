.DEFAULT_GOAL := open

clean:
	rm -f ./api/bundledResources.go
	rm -f ./bundled.go
	rm -rf *.app
	rm -f *.exe
.PHONY:clean

bundle: clean
	fyne bundle -o ./bundledResources.go ./api/conf.toml
	fyne bundle -o ./bundled.go ./assets/Logos_UP.png
.PHONY:bundle

package: bundle
	fyne package
.PHONY:package

open: package
	open Wazuh\ RD.app
.PHONY:open
