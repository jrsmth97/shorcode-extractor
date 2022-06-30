#SHORTCODE EXTRACTOR<br/>

#Javascript Edition<br/>

```
replace inside bracket shortcode with dynamic value
```
<br/>

#EXAMPLE USAGE
```
const shortcode = new ShortcodeExtractor()
const result = shortcode.setText("test string with {variable} inserted and {variable2} also {notexistvariable}")
                        .setVariables({
                            variable: 'hello',
                            variable2: 'world',
                            variable3: 'say hi',
                        })
                        .extract()

console.log(result)

```