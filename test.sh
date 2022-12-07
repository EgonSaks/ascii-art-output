go run . --output=banner.txt "hello" standard
# cat -e banner.txt

go run . --output=banner.txt "Hello There\!" shadow
# cat -e banner.txt

go run . --output test.txt banana standard 
# cat test.txt

go run . --output=test00.txt "First\nTest" standard
# cat test00.txt

go run . --output=test01.txt "hello" standard
# cat test01.txt

go run . --output=test02.txt "123 -> #$%" standard
# cat test02.txt

go run . --output=test03.txt "432 -> #$%&@" shadow
# cat test03.txt

go run .  --output=test04.txt "There" shadow
# cat test04.txt

go run . --output=test05.txt "123 -> \"#$%@" tinkertoy
# cat test05.txt

go run . --output=test06.txt "2 you" tinkertoy
# cat test06.txt

go run . --output=test07.txt "Testing long output\!" standard
# cat test07.txt

go run . --output=test08.txt "abcDEF" standard
# cat test08.txt

go run . --output=test010.txt "#%$" standard
# cat test10.txt

go run . --output=test011.txt "abcDEF 123" standard
# cat test11.txt

#go run . --output=test08.txt "abcDEF" randomFont
# could not read the content in the file: open : no such file or directory

# go run . --output=test010.txt "#%$" randomFont
# could not read the content in the file: open : no such file or directory

# go run . --output=test011.txt "abcDEF 123" randomFont
# could not read the content in the file: open : no such file or directory



