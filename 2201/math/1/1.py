
def show_print(p,a,b):
    return [(x,y) for x in range(p) for y in range(p) if (y*y - (x*x*x + a*x+b))%p == 0]

a = 2
b = 3
p = 97
print(show_print(p,a,b))