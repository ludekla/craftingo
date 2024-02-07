import os

structs = [
    'Binary: left Expr, operator tk.Token, right Expr',
    'Grouping: expression Expr',
    'Literal: value Any',
    'Unary: operator tk.Token, expression Expr'
]

top_section = '''package grammar

import tk "glox/pkg/tokens"

type Expr interface {
    Any
}'''

def setup_struct(expr: str):
    clsname, signature = expr.split(': ')
    pieces = signature.split(', ')
    top = f'type {clsname} ' + 'struct {'
    body = '    ' + '\n    '.join(piece for piece in pieces)
    bottom = '}'
    return '\n'.join([top, body, bottom])

def write():
    with open('pkg/grammar/expr.go', mode='w') as fout:
        fout.write(top_section + '\n\n')
        for struct in structs:
            struct_str = setup_struct(struct)
            fout.write(struct_str + '\n\n')



if __name__ == '__main__':

    write()

    os.system('go fmt pkg/grammar/*')





