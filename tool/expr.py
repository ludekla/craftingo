import os
from typing import List

structs = [
    'Binary[R ReturnType]: Left Expr[R], Operator tk.Token, Right Expr[R]',
    'Grouping[R ReturnType]: Expression Expr[R]',
    'Literal[R ReturnType]: Value any',
    'Unary[R ReturnType]: Operator tk.Token, Expression Expr[R]'
]

top_section = '''package grammar

import (
    tk "glox/pkg/tokens"
)

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

def setup_new(expr: str):
    clsname, struct_signature = expr.split(': ')
    pieces = []
    signature = []
    for piece in struct_signature.split(', '):
        obj, typ = piece.split()
        obj = obj.lower()
        pieces.append(obj)
        signature.append(f'{obj} {typ}')
    signature = ', '.join(signature)
    top = f'func New{clsname}({signature}) '
    clsname = clsname.replace('R ReturnType', 'R')
    top += f'{clsname}' + '{'
    body = f'    return {clsname}' + '{' + ', '.join(pieces) + '}'
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

def implement_accept(exprs: List[str]):
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
        for struct in structs:
            new_str = setup_new(struct)
            fout.write(new_str + '\n\n')
        visitor = setup_visitor(structs)
        fout.write(visitor + '\n\n')
        impls = implement_accept(structs)
        fout.write(impls)


if __name__ == '__main__':

    write()

    os.system('go fmt pkg/grammar/*')





