#SHORTCODE EXTRACTOR<br/>

#Python Edition<br/>

```
replace inside bracket shortcode with dynamic value
```
<br/>

#EXAMPLE USAGE
```
shortcode = ShortcodeExtractor()
shortcode.set_text("test string with {variable} inserted and (variable2) also [notexistvariable]")
shortcode.set_bracket_type('parentheses')
shortcode.set_variables({
     'variable': 'hello',
     'variable2': 'world',
     'variable3': 'say hi',
})
result = shortcode.extract()
print(result)

```