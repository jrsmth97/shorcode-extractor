class ShortcodeExtractor {
    #text = "";
    #bracketType = "curly";
    #regexPattern;
    #variables = {};

    constructor() {}

    setText(text) {
        this.#text = text;
        return this;
    }
    
    setVariables(variables) {
        this.#variables = variables;
        return this;
    }
    
    setBracketType(type) {
        this.#bracketType = type;
        return this;
    }

    extract() {
        this.#regexBuild();
        const variables = [...this.#text.matchAll(this.#regexPattern)];
        return variables.reduce((txt, va) => {
            const val = this.#variables[va[1]] || null
            if (val) txt = txt.replace(va[0], val)
            return txt
        }, this.#text)
    }

    #regexBuild() {
        switch (this.#bracketType) {
            case 'curly':
                this.#regexPattern = new RegExp('{({*[^{}]*}*)}', 'g');
            break;
            case 'square':
                this.#regexPattern = new RegExp('\\[([^\\]\\[\\r\\n]*)\\]', 'g');
            break;
            case 'parentheses':
                this.#regexPattern = new RegExp('\\(([^\\)]+)\\)', 'g');
            break;
            default:
                throw new Error('invalid bracket type');

        }
    }

}
