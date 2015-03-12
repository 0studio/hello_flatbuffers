# -*- coding:utf-8 -*-
.PHONY: build
build:
	flatc -g demo.idl
	cd demo && go install
	flatc -g demo2.idl
	cd demo2 && go install

# 	g++ -o nul  -Wall -Wextra  -pedantic -fsyntax-only -Wno-variadic-macros -std=c99 $(INCLUDES) -S ${CHK_SOURCES}
