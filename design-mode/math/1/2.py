# ECC在Fp域上加法、倍乘运算
# 求value在Fp域的逆——用于分数求逆
def get_inverse(value, p):
    for i in range(1, p):
        if (i * value) % p == 1:
            return i
    return -1

# 求最大公约数——用于约分化简
def get_gcd(x, y):
    if y == 0:
        return x
    else:
        return get_gcd(y, x % y)

# 计算P+Q函数
def calculate_p_q(x1,y1,x2,y2,a,b,p):
    flag = 1  # 控制符号位

    # 若P = Q，则k=[(3x1^2+a)/2y1]mod p
    if x1 == x2 and y1 == y2:
        member = 3 * (x1 ** 2) + a  # 计算分子
        denominator = 2 * y1        # 计算分母

    # 若P≠Q，则k=(y2-y1)/(x2-x1) mod p
    else:
        member = y2 - y1
        denominator = x2 - x1
        if member* denominator < 0:
            flag = 0
            member = abs(member)
            denominator = abs(denominator)

    # 将分子和分母化为最简
    gcd_value = get_gcd(member, denominator)
    member = member // gcd_value
    denominator = denominator // gcd_value

    # 求分母的逆元
    inverse_value = get_inverse(denominator, p)
    k = (member * inverse_value)

    if flag == 0:
        k = -k
    k = k % p

    # 计算x3,y3
    """
        x3≡k^2-x1-x2(mod p)
        y3≡k(x1-x3)-y1(mod p)
    """
    x3 = (k ** 2 - x1 - x2) % p
    y3 = (k * (x1 - x3) - y1) % p
    return [x3,y3]

# 计算nP函数
def calculate_np(p_x, p_y,a,b,p):
    tem_x = p_x
    tem_y = p_y
    p_value = calculate_p_q(tem_x,tem_y, p_x, p_y,a,b,p)
    tem_x = p_value[0]
    tem_y = p_value[1]
    return p_value

# PPT例2、例3：y^2=x^3+2x+3(mod 97)
p, a, b = 97, 2, 3

# PT例2计算P+Q ,其中p=(3,10)、q=(9,7)
# print(calculate_p_q(3,10,9,7,1,1,23))

 # PT例3计算2P ,其中p=(3,10)
# print(calculate_np(3,6,a,b,p))
#
# print(calculate_p_q(3,6,80,10,a,b,p))

x1=3
y1=6
xn=3
yn=6
for i in range(1, 10):
    x = calculate_p_q(x1,y1,xn,yn,a,b,p)
    xn=x[0]
    yn=x[1]
    print(x)

