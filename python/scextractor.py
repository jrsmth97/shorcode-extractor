import re

class ShortcodeExtractor:    
    def __init__(self):
        self.text = ""
        self.bracket_type = "curly"
        self.bracket_open = "{"
        self.bracket_close = "}"
        self.variables = {}
        self.regex_pattern = ""

    def set_text(self, text):
        self.text = text
        return self

    def set_bracket_type(self, type):
        self.bracket_type = type
        return self

    def set_variables(self, variables):
        self.variables = variables
        return self

    def extract(self):
        self.regex_build()
        regex_variables = re.findall(self.regex_pattern, self.text)
        new_text = self.text
        for va in regex_variables:
            val = self.variables.get(va)
            if val:
                new_text = new_text.replace(str(self.bracket_open) + str(va) + str(self.bracket_close), val)

        return new_text
    
    def regex_build(self):
         match self.bracket_type:
            case 'curly':
                self.regex_pattern = r'{({*[^{}]*}*)}'
                self.bracket_open = "{"
                self.bracket_close = "}"
            case 'square':
                self.regex_pattern = r'\[([^\]\[\r\n]*)\]'
                self.bracket_open = "["
                self.bracket_close = "]"
            case 'parentheses':
                self.regex_pattern = r'\(([^\)]+)\)'
                self.bracket_open = "("
                self.bracket_close = ")"
            case _:
                print("Invalid bracket type")
                exit()
