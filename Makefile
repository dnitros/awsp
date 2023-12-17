BIN_DIR = /usr/local/bin

help:          ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

install:       ## Install Target
	GOOS= GOARCH= GOARM= GOFLAGS= go build -o ${BIN_DIR}/_awsp_prompt
	cp scripts/_awsp ${BIN_DIR}/_awsp
	@echo "Add <alias awsp=\"source _awsp\"> to your .bash_profile/.bashrc/zshrc then open new terminal or source that file"

uninstall:     ## Uninstall Target
	rm -f ${BIN_DIR}/_awsp
	rm -f ${BIN_DIR}/_awsp_prompt