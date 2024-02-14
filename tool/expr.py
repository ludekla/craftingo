import os
from typing import List

structs = [
    'Binary[R ReturnType]: left Expr[R], operator tk.Token, right Expr[R]',
    'Grouping[R ReturnType]: expression Expr[R]',
    'Literal[R ReturnType]: value Expr[R]',
    'Unary[R ReturnType]: operator tk.Token, expression Expr[R]'
]

top_section = '''package grammar

import tk "glox/pkg/tokens"

type ReturnType interface { string | float64 | int }'''

def define_expr(exprs: List[str]):
    parts = [
        'type Expr[R ReturnType] interface {',
        f'    Accept(v Visitor[R]) R',
        '}'
    ]
    return '\n'.join(parts)

def setup_struct(expr: str):
    clsname, signature = expr.split(': ')
    pieces = signature.split(', ')
    top = f'type {clsname} ' + 'struct {'
    body = '    ' + '\n    '.join(piece for piece in pieces)
    bottom = '}'
    return '\n'.join([top, body, bottom])

def setup_visitor(exprs: List[str]):
    parts = ['type Visitor[R ReturnType] interface {']
    for expr in exprs:
        clsname, *_ = expr.split(':')
        clsname = clsname.replace('ReturnType', '').strip()
        n = clsname.find('[')
        pure_clsname = clsname[:n]
        parts.append(f'   Visit{pure_clsname}({clsname[0].lower()} {clsname}) R')
    parts.append('}')
    return '\n'.join(parts)

def implement_visitor(exprs: List[str]):
    parts = []
    for expr in exprs:
        clsname, *_ = expr.split(':')
        obj = clsname[0].lower()
        clsname = clsname.replace('ReturnType', '').strip()
        n = clsname.find('[')
        pure_clsname = clsname[:n]
        code_lines = [
            f'func ({obj} {clsname}) ' + 'Accept(v Visitor[R]) R {',
            f'    return v.Visit{pure_clsname}({obj})',
            '}'
        ]
        parts.append('\n'.join(code_lines))
    return '\n\n'.join(parts)

def write():
    with open('pkg/grammar/expr.go', mode='w') as fout:
        fout.write(top_section + '\n\n')
        fout.write(define_expr(structs) + '\n\n')
        for struct in structs:
            struct_str = setup_struct(struct)
            fout.write(struct_str + '\n\n')
        visitor = setup_visitor(structs)
        fout.write(visitor + '\n\n')
        impls = implement_visitor(structs)
        fout.write(impls)


if __name__ == '__main__':

    write()

    os.system('go fmt pkg/grammar/*')





