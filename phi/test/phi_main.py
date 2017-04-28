import vitessedata.phi
import sys

vitessedata.phi.DeclareTypes('''
//
// BEGIN INPUT TYPES
// a int32
// b float32
// c string
// END INPUT TYPES
//
// BEGIN OUTPUT TYPES
// x int32
// y float32
// z string
// END OUTPUT TYPES
//
''')

def do_x():
    while True:
        rec = vitessedata.phi.NextInput()
        if not rec:
            sys.stderr.write("Py X: end of input\n")
            break

        sys.stderr.write("Py X: input rec " + str(rec) + "\n")
        outrec = [None, None, None]

        if rec[0] is None:
            outrec[0] = rec[0]
        else:
            outrec[0] = rec[0] * 2

        if rec[1] is None:
            outrec[1] = rec[1]
        else:
            outrec[1] = rec[1] * 2.0

        if rec[2] is None:
            outrec[2] = None 
        else:
            outrec[2] = rec[2] + rec[2] 

        vitessedata.phi.WriteOutput(outrec)

    vitessedata.phi.WriteOutput(None)

if __name__ == '__main__':
    do_x()