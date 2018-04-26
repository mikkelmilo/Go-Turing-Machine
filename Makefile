dir=Compiler/antlr-parser/

build:
	java -jar ~bin/antlr-4.7-complete.jar -Dlanguage=Go $(dir)/TML.g4
clean:
	rm -f $(dir)/tml_lexer.go
	rm -f $(dir)/tml_base_listener.go
	rm -f $(dir)/tml_parser.go
	rm -f $(dir)/tml_listener.go
	rm -f $(dir)/*.tokens