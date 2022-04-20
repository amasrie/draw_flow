# Vue Application

Application based on [this example](https://github.com/jerosoler/drawflow-vue3-example)

If you want to create an isolated Vue container, follow these instructions:

```sh
docker run -dit --name vue_app -p 9501:9501 -p 9601:9601 -v "$PWD:/home/vue" vue
docker start vue_app
docker exec -it vue_app bash
```

Make sure that the `dev` script inside the `package.json` file is using port `9501` while the file `vite.config.js` is using port `9601`.

## Vue 3 + Vite

This template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

## Recommended IDE Setup

- [VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=johnsoncodehk.volar)

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run dev
```

After that, the application could be accessed from `localhost:9501`

### Dependencies

To properly use this application, make sure the `HOST` variable in the `drawflow.vue` file is pointing to  your service.


## Programming Graph - Node Editor

This dashboard contains a list of nodes that allows the clients to create basic programs that will be translated into Python code.

### General graph rules

* Cycles are not allowed (ie: Node1 -> Node2, Node2 -> Node1).
* Each node may have a different amount of inputs and outputs depending of their type.
* Every input and output may have up to one connection.

### Available Nodes

#### Boolean Nodes

This node represent a boolean constant.

* Attributes: 
	- A switch that represent the value of the boolean (`True` or `False`)
* Inputs: None
* Outputs: 
	- `output_1`: Any boolean expression

#### Number Nodes

This node represent a numeric constant with up to 5 decimals.

* Attributes: 
	- A numeric input that represent the value of the number
* Inputs: None
* Outputs: 
	- `output_1`: Any numeric expression

#### Variable Nodes

This node represent a variable and as such it may represent any kind of expression.

* Attributes: 
	- A text input that indicates the name of the variable. Note that the input must have a valid python variable name, otherwise the color will turn red.
* Inputs: None
* Outputs: 
	- `output_1`: Connects with almost any kind of node

#### Arithmetic Operator Nodes

This node allows to create a basic operation that takes numeric expressions and returns a number.

* Attributes: 
	- A select input with the available unary and binary numeric operations.
* Inputs: 
	- `input_1`: A numeric expression
	- `input_2`: A numeric expression (hidden when a unary operator is selected)
* Outputs: 
	- `output_1`: Any numeric expression

#### Comparison Operator Nodes

This node allows to create a basic operation that takes numeric expressions and returns a boolean.

* Attributes: 
	- A select input with the available binary comparison operations.
* Inputs: 
	- `input_1`: A numeric expression
	- `input_2`: A numeric expression
* Outputs: 
	- `output_1`: Any boolean expression

#### Logical Operator Nodes

This node allows to create a basic operation that takes boolean expressions and returns a boolean.

* Attributes: 
	- A select input with the available unary and binary logical operations.
* Inputs: 
	- `input_1`: A boolean expression
	- `input_2`: A boolean expression (hidden when a unary operator is selected)
* Outputs: 
	- `output_1`: Any boolean expression

#### Assign Nodes

This instruction node associates a variable with an expression.

* Attributes: None
* Inputs: 
	- `input_1`: The previous instruction
	- `input_2`: The variable that will store the expression
	- `input_3`: A constant or expression (numerical, boolean or a variable) that will be assigned to the defined variable in `input_2`
* Outputs: 
	- `output_1`: The next instruction (optional)

#### Print Nodes

This instruction node represent a print instruction.

* Attributes: None
* Inputs: 
	- `input_1`: The previous instruction
	- `input_2`: A constant or expression (numerical, boolean or a variable) that will be printed
* Outputs: 
	- `output_1`: The next instruction (optional)

#### Conditional Nodes

This instruction node verifies if a condition is satisfied in order to execute another block of instructions.

* Attributes: 
	- A select input with the types of conditions available (first condition, next condition or otherwise)
* Inputs: 
	- `input_1`: The previous instruction
	- `input_2`: A boolean constant or expression that will be verified (hidden when else is selected)
* Outputs: 
	- `output_1`: The next external instruction (optional)
	- `output_2`: The first instruction inside the condition block

#### Loop Nodes

This instruction node loops over a range of numbers and executes a block of instructions for each of them.

* Attributes: None
* Inputs: 
	- `input_1`: The previous instruction
	- `input_2`: A variable that will loop over the defined range
	- `input_3`: A numeric constant or expression that represent the first element in the range
	- `input_4`: A numeric constant or expression that represent the last element in the range
* Outputs: 
	- `output_1`: The next external instruction (optional)
	- `output_2`: The first instruction inside the loop block

### AST Representation

Consider the following rules as a basic representation of this language:

```
_start: _seq_operations
_seq_operations: _operations _seq_operations | _operations
_operations: _assign | _cond | _loop | _print
_print: print(_expr)
_assign: VAR = _expr
_expr: _numexpr | _bolexpr
_numexpr: NUMBER | VAR | _numexpr + _numexpr | _numexpr - _numexpr | - _numexpr | _numexpr * _numexpr | _numexpr / _numexpr | _numexpr // _numexpr | _numexpr mod _numexpr | _numexpr ** _numexpr
_bolexpr: True | False | VAR | _bolexpr and _bolexpr | _bolexpr or _bolexpr | not _bolexpr | _numexpr < _numexpr | _numexpr <= _numexpr | _numexpr > _numexpr | _numexpr >= _numexpr | _numexpr == _numexpr | _numexpr != _numexpr 
_cond: if _bolexpr: _seq_operations _next_cond
_next_cond: elif _bolexpr: _seq_operations _next_cond | else: _seq_operations
_loop: for VAR in range(_numexpr, _numexpr): _seq_operations
```

Notice that:
* Everything with suffix `_` is considered a rule.
* `_start` is the starting rule.
* A rule may execute different options, each of them separated by a pipe (`|`).
* `VAR` represents a variable name.
* `NUMBER` represents a number.
* Any other element is a reserved word of the language.
* Indentation is not considered for these rules.
