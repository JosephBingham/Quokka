# This document goes though the syntax and design principles of the Quokka programming language

## Comments:
## Comments take the form:
    ~comment~
    ~
    comment 2
    ~


## Types:
### The native types in the Quokka progamming language are 
#### INT :: integer
#### B :: bit or bool
#### FIX :: fixed point value
#### CHAR :: character
#### QINT :: quantum INT
#### QB :: quantum B
#### QFIX :: quantum FIX
#### There is no quantum CHAR
##### What would that even mean?

### Declarations of these types look like
    QINT a = (1|.25>, -1|.25>, 0|.5>); ~declartation of a quantum integer variable :: (value1|prob1>, ...)~

    QINT aa = (1|.25>, -1|.25>, 0|>); ~can leave out last prob since it is implied~

    QINT aaa = 1|.25> -1|.25> 0|>; ~can leave out parens and commas, not pretty~

    ~representation { (prob1, value1), (prob1+prob2, value2) ... (sum(prob*), valuen) }~

    QB b = 1|.25>; ~declaration of quantum bit, only need to declare one value~


