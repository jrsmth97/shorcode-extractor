<?php

class ShortcodeExtractor {
    private $text;
    private $bracketType = "curly";
    private $regexPattern;
    private $variables;

    public function __construct() {}

    public function setText($text) {
        $this->text = $text;
        return $this;
    }

    public function setBracketType($type) {
        $this->bracketType = $type;
        return $this;
    }

    public function setVariables($variables) {
        $this->variables = $variables;
        return $this;
    }

    public function extract() {
        $this->regexBuild();
        preg_match_all($this->regexPattern, $this->text, $regexResults);
        $newText = $this->text;
        for ($i = 0; $i < count($regexResults); $i++) {
            $val = $this->variables[$regexResults[1][$i]] ?? null;
            if ($val) {
                $newText = str_replace($regexResults[0][$i], $val, $newText);
            } 
        }

        return $newText;
    }

    private function regexBuild() {
        switch ($this->bracketType) {
            case 'curly':
                $this->regexPattern = '/{({*[^{}]*}*)}/';
            break;
            case 'square':
                $this->regexPattern = '/\[([^\]\[\r\n]*)\]/';
            break;
            case 'parentheses':
                $this->regexPattern = '/\(([^\)]+)\)/';
            break;
            default:
                throw new Error('invalid bracket type');
        }
    }
}