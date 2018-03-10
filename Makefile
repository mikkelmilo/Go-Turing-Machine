build:
	java -jar ~bin/antlr-4.7-complete.jar -Dlanguage=Go -o parser TM-Language/TML.g4
clean:
	rm -f parser/TM-Language/*_lexer.go
	rm -f parser/TM-Language/*_base_listener.go
	rm -f parser/TM-Language/*_parser.go
	rm -f parser/TM-Language/*_listener.go
	rm -f parser/TM-Language/*.tokens