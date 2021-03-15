from itertools import product

remap = {(0,0): '1',(0,1): '2',(0,2): '3',
         (1,0): '4',(1,1): '5',(1,2): '6',
         (2,0): '7',(2,1): '8',(2,2): '9',
         (3,0):'-1',(3,1): '0',(3,2):'-1' }

inputs = {val : key
          for key, val
          in remap.items()}

def get_pins(observed):
    return map(''.join,product(*[
        [
            remap[(y1+y2,x1+x2)]
            for y2,x2 in [(1,0),(0,-1),( 0,0),(0,1),(-1,0)]
            if 0<=y1+y2<=3
            and 0<=x1+x2<=2
            and remap[(y1+y2,x1+x2)] != '-1'
        ] for y1,x1 in [ inputs[o] for o in observed ]
    ]))