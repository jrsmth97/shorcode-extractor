#SHORTCODE EXTRACTOR<br/>

#PHP Edition<br/>

```
replace inside bracket shortcode with dynamic value
```
<br/>

#EXAMPLE USAGE
```
$shortcode = new ShortcodeExtractor();
$result = $shortcode->setText("test string with {variable} inserted and [variable2] also {notexistvariable}")
                    ->setBracketType('square')
                    ->setVariables([
                        'variable' => 'hello',
                        'variable2' => 'world',
                        'variable3' => 'say hi',
                    ])
                    ->extract();

echo $result;

```