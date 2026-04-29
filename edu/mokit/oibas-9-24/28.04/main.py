n_one = 10
n_two = 0

try:
    print(n_one / n_two)
except TypeError:
    print("E101")
except ZeroDivisionError:
    print("E102")
except:
    pass


print('algo....')
